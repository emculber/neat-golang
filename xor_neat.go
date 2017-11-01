package main

import "fmt"

var number_runs int = 100
var population_size int = 150

func main() {
	start_genome := GenomeInit()
	for expcount := 0; expcount < number_runs; expcount++ {
		fmt.Println("Experiment #", expcount)
		fmt.Println("Start Genome:", start_genome)
		fmt.Println("Spawning Population off Genome")
		//population := PopulationInit(start_genome, population_size)
		PopulationInit(start_genome, population_size)
		break
	}
}
