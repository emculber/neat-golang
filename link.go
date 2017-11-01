package main

type Link struct {
	Weight        float64
	In_node       Node
	Out_node      Node
	Is_recurrent  bool
	Time_delay    bool
	Adding_weight float64
}
