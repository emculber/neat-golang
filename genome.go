package main

import "fmt"

const (
	GAUSSIAN     = 0
	COLDGAUSSIAN = 1
)

type Genome struct {
	Genome_id int
	Nodes     []Node
	Gene      []Gene
	Phenotype Network
}

func GenomeInit() Genome {
	genome := Genome{Genome_id: 1}
	Nodes := []Node{
		Node{
			Activesum:      0,
			Node_id:        1,
			Node_type:      SENSOR,
			Gen_node_label: BIAS,
		},
		Node{
			Activesum:      0,
			Node_id:        2,
			Node_type:      SENSOR,
			Gen_node_label: INPUT,
		},
		Node{
			Activesum:      0,
			Node_id:        3,
			Node_type:      SENSOR,
			Gen_node_label: INPUT,
		},
		Node{
			Activesum:      0,
			Node_id:        4,
			Node_type:      SENSOR,
			Gen_node_label: INPUT,
		},
		Node{
			Activesum:      0,
			Node_id:        5,
			Node_type:      SENSOR,
			Gen_node_label: INPUT,
		},
		Node{
			Activesum:      0,
			Node_id:        6,
			Node_type:      NEURON,
			Gen_node_label: OUTPUT,
		},
		Node{
			Activesum:      0,
			Node_id:        7,
			Node_type:      NEURON,
			Gen_node_label: OUTPUT,
		},
	}
	fmt.Println(Nodes)
	Gene := []Gene{
		Gene{
			Link: Link{
				Weight:        0.0,
				In_node:       Nodes[0],
				Out_node:      Nodes[5],
				Is_recurrent:  false,
				Time_delay:    false,
				Adding_weight: 0,
			},
			Innovation_number: 1,
			Mutation_number:   0,
			Enable:            true,
		},
		Gene{
			Link: Link{
				Weight:        0.0,
				In_node:       Nodes[1],
				Out_node:      Nodes[5],
				Is_recurrent:  false,
				Time_delay:    false,
				Adding_weight: 0,
			},
			Innovation_number: 2,
			Mutation_number:   0,
			Enable:            true,
		},
		Gene{
			Link: Link{
				Weight:        0.0,
				In_node:       Nodes[2],
				Out_node:      Nodes[5],
				Is_recurrent:  false,
				Time_delay:    false,
				Adding_weight: 0,
			},
			Innovation_number: 3,
			Mutation_number:   0,
			Enable:            true,
		},
		Gene{
			Link: Link{
				Weight:        0.0,
				In_node:       Nodes[3],
				Out_node:      Nodes[5],
				Is_recurrent:  false,
				Time_delay:    false,
				Adding_weight: 0,
			},
			Innovation_number: 4,
			Mutation_number:   0,
			Enable:            true,
		},
		Gene{
			Link: Link{
				Weight:        0.0,
				In_node:       Nodes[4],
				Out_node:      Nodes[5],
				Is_recurrent:  false,
				Time_delay:    false,
				Adding_weight: 0,
			},
			Innovation_number: 5,
			Mutation_number:   0,
			Enable:            true,
		},
		Gene{
			Link: Link{
				Weight:        0.0,
				In_node:       Nodes[0],
				Out_node:      Nodes[6],
				Is_recurrent:  false,
				Time_delay:    false,
				Adding_weight: 0,
			},
			Innovation_number: 6,
			Mutation_number:   0,
			Enable:            true,
		},
		Gene{
			Link: Link{
				Weight:        0.0,
				In_node:       Nodes[1],
				Out_node:      Nodes[6],
				Is_recurrent:  false,
				Time_delay:    false,
				Adding_weight: 0,
			},
			Innovation_number: 7,
			Mutation_number:   0,
			Enable:            true,
		},
		Gene{
			Link: Link{
				Weight:        0.0,
				In_node:       Nodes[2],
				Out_node:      Nodes[6],
				Is_recurrent:  false,
				Time_delay:    false,
				Adding_weight: 0,
			},
			Innovation_number: 8,
			Mutation_number:   0,
			Enable:            true,
		},
		Gene{
			Link: Link{
				Weight:        0.0,
				In_node:       Nodes[3],
				Out_node:      Nodes[6],
				Is_recurrent:  false,
				Time_delay:    false,
				Adding_weight: 0,
			},
			Innovation_number: 9,
			Mutation_number:   0,
			Enable:            true,
		},
		Gene{
			Link: Link{
				Weight:        0.0,
				In_node:       Nodes[4],
				Out_node:      Nodes[6],
				Is_recurrent:  false,
				Time_delay:    false,
				Adding_weight: 0,
			},
			Innovation_number: 10,
			Mutation_number:   0,
			Enable:            true,
		},
	}
	genome.Nodes = Nodes
	genome.Gene = Gene
	fmt.Println(Gene)
	fmt.Println(genome)
	return genome
}
