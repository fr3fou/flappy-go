package ai

import (
	"log"
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
		Nodes:     4,
		Activator: gone.Sigmoid(),
	},
	{
		Nodes:     6,
		Activator: gone.Sigmoid(),
	},
	{
		Nodes:     2,
		Activator: gone.Sigmoid(),
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

	// Normalize fitness
	sum := 0.0
	for _, bird := range g.Birds {
		sum += float64(bird.Score)
	}

	for _, bird := range g.Birds {
		bird.Fitness = float64(bird.Score) / sum
	}

	newBirds := make([]*Bird, g.PopulationAmount)

	// Make the new population
	for i := 0; i < g.PopulationAmount; i++ {
		// Should we do crossover or a bare copy of the best one
		if rand.Float64() > g.CrossoverOdds {
			child := g.NaturalSelection()

			child.Brain.Mutate(gone.GaussianMutation(g.MutationRate))

			newBirds[i] = child
		} else {
			firstParent := g.NaturalSelection()
			secondParent := g.NaturalSelection()

			childBrain, err := firstParent.Brain.Crossover(secondParent.Brain)
			if err != nil {
				return err
			}

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
	brainCopy := bird.Brain.Copy()
	return NewBird(flappy.BirdSize*2, flappy.Height/2, brainCopy)
}

func (g *Game) Update() {
	firstPipe := g.Pipes[0]

	for _, pipe := range g.Pipes {
		pipe.X -= flappy.Speed
	}

	for _, bird := range g.Birds {
		for _, pipe := range g.Pipes {
			if pipe.IsAround(pipe.Rectangle) {
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

		if bird.ShouldJump(firstPipe) {
			bird.Jump()
		}

		bird.Update()
	}

	// Remove first pipe if it's offscreen
	if firstPipe.IsOffscreen() {
		g.Pipes = g.Pipes[1:]
	}

	hasAlive := false
	for _, bird := range g.Birds {
		if bird.Alive {
			hasAlive = true
			break
		}
	}

	if !hasAlive {
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
