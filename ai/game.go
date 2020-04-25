package ai

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/fr3fou/flappy-go/flappy"
	"github.com/fr3fou/gone/gone"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	*flappy.Game
	Birds            []*Bird
	PopulationAmount int
	MutationRate     float64
	CrossoverOdds    float64
	Generation       int
}

var BrainLayers = []gone.Layer{
	{
		Nodes:     5,
		Activator: gone.Sigmoid(),
	},
	{
		Nodes:     8,
		Activator: gone.Sigmoid(),
	},
	{
		Nodes:     2,
		Activator: gone.Sigmoid(), // ideally should be softmax
	},
}

func NewGame(populationAmount int, mutationRate, crossoverOdds float64) *Game {
	birds := make([]*Bird, populationAmount)

	// Initialize population
	for i := range birds {
		birds[i] = NewBird(flappy.BirdSize*2, flappy.Height/2, gone.New(1, gone.MSE(), BrainLayers...))
	}

	g := &Game{
		Game:             flappy.NewGame(),
		Birds:            birds,
		Generation:       1,
		CrossoverOdds:    crossoverOdds,
		PopulationAmount: populationAmount,
		MutationRate:     mutationRate,
	}
	g.Init()

	return g
}

func (g *Game) Init() {
	g.Ground = flappy.NewGround()

	g.Pipes = make([]*flappy.Pipe, flappy.MaxPipes)
	initialOffset := flappy.HorizontalGap + flappy.BirdSize*2*2
	offset := initialOffset
	for i := range g.Pipes {
		g.Pipes[i] = flappy.NewPipe((i)*flappy.PipeWidth+offset, 0)
		offset = flappy.HorizontalGap*(i+1) + initialOffset
	}
}

func (g *Game) NextGeneration() error {
	g.Generation++
	log.Printf("Generation Number %d...", g.Generation)

	// penalize worse scores
	// reward higher scores
	for _, bird := range g.Birds {
		bird.Score = int(math.Pow(float64(bird.Score), 2))
	}

	// Normalize fitness
	sum := 0.0
	for _, bird := range g.Birds {
		sum += float64(bird.Score)
	}

	for i, bird := range g.Birds {
		bird.Fitness = float64(bird.Score) / sum
		fmt.Println(i, bird.Fitness)
	}

	newBirds := make([]*Bird, g.PopulationAmount)
	// Make the new population
	for i := 0; i < g.PopulationAmount; i++ {
		// Should we do crossover or a bare copy of the best one
		if rand.Float64() > g.CrossoverOdds {
			child := g.NaturalSelection()
			child.Brain.Mutate(gone.GaussianMutation(g.MutationRate, 1, 0))
			newBirds[i] = child
		} else {
			firstParent := g.NaturalSelection()
			secondParent := g.NaturalSelection()
			childBrain, err := firstParent.Brain.Crossover(secondParent.Brain)
			if err != nil {
				return err
			}
			childBrain.Mutate(gone.GaussianMutation(g.MutationRate, 1, 0))
			child := NewBird(flappy.BirdSize*2, flappy.Height/2, childBrain)
			newBirds[i] = child
		}
	}
	g.Birds = newBirds

	return nil
}

func (g *Game) NaturalSelection() *Bird {
	// https://github.com/CodingTrain/Toy-Neural-Network-JS/blob/master/examples/neuroevolution-flappybird/ga.js#L63-L87
	i := 0
	probability := rand.Float64()
	for probability > 0 {
		probability -= g.Birds[i].Fitness
		i++
	}
	i--
	bird := g.Birds[i]
	return NewBird(flappy.BirdSize*2, flappy.Height/2, bird.Brain.Copy())
}

func (g *Game) Update() {
	for _, pipe := range g.Pipes {
		pipe.X -= flappy.Speed
	}

	for _, bird := range g.Birds {
		if !bird.Alive {
			continue
		}

		for _, pipe := range g.Pipes {
			if pipe.IsAround(bird.Rectangle) {
				bird.Score++
			}

			if pipe.CollidesWith(bird.Rectangle) {
				bird.Alive = false
				break
			}
		}

		if g.Ground.CollidesWith(bird.Rectangle) || bird.AboveSky() {
			bird.Alive = false
		}
	}

	for _, bird := range g.Birds {
		if !bird.Alive {
			continue
		}

		// the next pipe is the closest one if the closestPipe is behind the bird
		closestPipe := g.Pipes[0]
		if g.Pipes[1].X < bird.X {
			closestPipe = g.Pipes[1]
		}

		if bird.ShouldJump(closestPipe) {
			bird.Jump()
		}

		bird.Update()
	}

	// Remove first pipe if it's offscreen
	if g.Pipes[0].IsOffscreen() {
		g.Pipes = g.Pipes[1:]
	}

	allDead := true
	for _, bird := range g.Birds {
		if bird.Alive {
			allDead = false
			break
		}
	}

	if allDead {
		g.Init()
		err := g.NextGeneration()
		if err != nil {
			log.Println(err)
		}
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.SkyBlue)

	for _, bird := range g.Birds {
		if bird.Alive {
			bird.Draw(100)
		}
	}

	for _, pipe := range g.Pipes {
		pipe.Draw()
	}

	g.Ground.Draw()
	rl.EndDrawing()
}
