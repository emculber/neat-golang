package main

type Organism struct {
	Fitness               float64
	Original_fitness      float64
	Error_value           float64
	Winner                bool
	Network               Network
	Genome                Genome
	Species               Species
	Expected_offspring    float64
	Generation            int
	Eliminate             bool
	Champion              bool
	Super_champ_offspring int
	Pop_champ             bool
	Pop_champ_child       bool
	High_fit              float64
	Time_alive            int
	Mutation_struct_baby  bool
	Mate_baby             bool
	Metadata              string
	Modified              bool
}

func OrganismInit(fit float64, genome Genome, gen int) {
	organism := Organism{
		Fitness:               fit,
		Original_fitness:      fit,
		Genome:                genome,
		Network:               genome.Genesis(),
		Species:               0,
		Generation:            gen,
		Eliminate:             false,
		Error_value:           0,
		Winner:                false,
		Champion:              false,
		Super_champ_offspring: 0,
		Time_alive:            0,
		Pop_champ:             false,
		Pop_champ_child:       false,
		High_fit:              0,
		Mutation_struct_baby:  0,
		Mate_baby:             0,
		Modified:              true,
	}
}
