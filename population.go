package main

import "fmt"

type Population struct {
	Organisms            []Organism
	Species              []Species
	Innovations          []Innovation
	last_species         int
	mean_fitness         float64
	variance             float64
	standard_deviation   float64
	winnergen            int
	highest_fitness      float64
	highest_last_changed float64
}

func PopulationInit(genome Genome, size int) {
	population := Population{
		winnergen:            0,
		highest_fitness:      0.0,
		highest_last_changed: 0,
	}
	population.spawn(genome, size)
}

func (population *Population) spawn(genome Genome, size int) {
	for count := 1; count <= size; count++ {
		fmt.Println("Creating Organism", count, "/", size)

		new_genome := genome
		new_genome.Genome_id = count
		new_genome.MutateLinkWeights(1.0, 1.0, COLDGAUSSIAN)
		new_organism := Organism{}
		population.Organisms = append(population.Organism, new_organism)
		fmt.Println(genome.Genome_id)
		fmt.Println(new_genome.Genome_id)
	}
}
