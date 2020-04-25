package ai

import (
	"log"
	"math"
	"math/rand"

	"github.com/fr3fou/flappy-go/flappy"
	"github.com/fr3fou/gone/gone"
)

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

	for _, bird := range g.Birds {
		bird.Fitness = float64(bird.Score) / sum
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
