package main

const (
	NEURON = 0
	SENSOR = 1
)

const (
	HIDDEN = 0
	INPUT  = 1
	OUTPUT = 2
	BIAS   = 3
)

const (
	SIGMOID = 0
)

type Node struct {
	Activation_count int
	Last_activation  float64
	Last_activation2 float64
	//dup              Node
	//analogue         Node
	Override_value float64
	Frozen         bool
	Function_type  int
	Node_type      int
	Activesum      float64
	Activation     float64
	Active_flag    bool
	Params         []float64
	Income         []Link
	Outgoing       []Link
	Rowlevels      []float64
	Row            int
	Ypos           int
	Xpos           int
	Node_id        int
	Gen_node_label int
}
