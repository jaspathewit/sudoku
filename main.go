package main

import (
	"log"
	"sudoku/grid"
)

func main() {
	// g, err := puzzelLibelle()
	// g, err := puzzelQuasidoku()
	// g, err := puzzelLovedoku()
	g, err := puzzelTriadoku()
	if err != nil {
		log.Fatalf("failed to create grid for puzzel %s", err)
	}
	//fmt.Printf("Grid \n%s", g)

	//for _, ref := range g.Cells {
	//	c, _ := g.Get(ref)
	//	fmt.Printf("Cell: %s\n", c)
	//}

	cg, err := g.Clone()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	//for _, ref := range cg.Cells {
	//	c, _ := cg.Get(ref)
	//	fmt.Printf("Cell: %s\n", c)
	//}

	_, err = Solve(cg, 0)
	if err != nil {
		log.Fatalf("solve: %s", err)
	}

	//fmt.Printf("Grid \n%s", g)

}

// puzzelLibelle
func puzzelLibelle() (*grid.Grid, error) {

	topology := grid.Normal{}

	g, err := grid.NewGrid(topology)
	if err != nil {
		return nil, err
	}

	// set up the starting values
	g.Set("1_2", "5")
	g.Set("1_3", "2")
	g.Set("1_5", "6")
	g.Set("1_6", "8")
	g.Set("1_8", "3")
	g.Set("2_2", "7")
	g.Set("2_5", "5")
	g.Set("2_7", "9")
	g.Set("2_8", "2")
	g.Set("3_2", "3")
	g.Set("3_5", "1")
	g.Set("3_9", "6")
	g.Set("5_3", "4")
	g.Set("5_6", "5")
	g.Set("5_7", "6")
	g.Set("6_3", "8")
	g.Set("6_5", "4")
	g.Set("6_7", "2")
	g.Set("7_1", "1")
	g.Set("7_2", "9")
	g.Set("7_6", "2")
	g.Set("7_8", "7")
	g.Set("8_6", "6")
	g.Set("8_9", "2")
	g.Set("9_5", "8")
	g.Set("9_7", "1")

	return g, nil
}

// puzzelQuasidoku
func puzzelQuasidoku() (*grid.Grid, error) {

	topology := grid.Quasidoku{}

	g, err := grid.NewGrid(topology)
	if err != nil {
		return nil, err
	}

	// set up the starting values
	g.Set("1_1", "2")
	g.Set("1_4", "8")
	g.Set("1_5", "1")
	g.Set("2_3", "3")
	g.Set("2_5", "6")
	g.Set("3_4", "2")
	g.Set("3_9", "6")
	g.Set("4_1", "7")
	g.Set("5_4", "1")
	g.Set("5_9", "9")
	g.Set("6_1", "8")
	g.Set("6_6", "2")
	g.Set("6_7", "3")
	g.Set("6_8", "5")
	g.Set("7_6", "4")
	g.Set("9_4", "9")
	g.Set("9_5", "3")

	g.SetLabel("1_3", "C")
	g.SetLabel("6_4", "B")
	g.SetLabel("7_5", "A")

	return g, nil
}

// puzzelLovedoku
func puzzelLovedoku() (*grid.Grid, error) {

	topology := grid.Lovedoku{}

	g, err := grid.NewGrid(topology)
	if err != nil {
		return nil, err
	}

	// set up the starting values
	g.Set("1_3", "1")
	g.Set("1_7", "7")
	g.Set("1_16", "1")
	g.Set("1_19", "5")
	g.Set("1_21", "3")
	g.Set("2_6", "4")
	g.Set("2_7", "2")
	g.Set("2_16", "3")
	g.Set("2_20", "4")

	g.Set("3_2", "7")

	g.Set("3_8", "6")
	g.Set("3_15", "2")
	g.Set("3_17", "8")
	g.Set("3_21", "9")

	g.Set("4_5", "4")
	g.Set("4_6", "9")

	g.Set("4_16", "8")
	g.Set("5_16", "5")
	g.Set("6_4", "6")
	g.Set("6_5", "2")
	g.Set("6_17", "1")

	g.Set("8_11", "3")

	g.SetLabel("1_18", "B")
	g.SetLabel("2_8", "A")
	g.SetLabel("3_4", "D")
	g.SetLabel("3_19", "E")

	g.SetLabel("4_4", "F")
	g.SetLabel("8_12", "C")

	return g, nil
}

// puzzelTriadoku
func puzzelTriadoku() (*grid.Grid, error) {

	topology := grid.Triadoku{}

	g, err := grid.NewGrid(topology)
	if err != nil {
		return nil, err
	}

	// set up the starting values
	g.Set("1_3", "6")
	g.Set("1_11", "8")
	g.Set("2_20", "6")

	g.Set("3_11", "3")
	g.Set("3_21", "7")
	g.Set("4_5", "3")
	g.Set("4_24", "2")
	g.Set("5_5", "5")
	g.Set("5_22", "5")
	g.Set("6_14", "4")
	g.Set("6_22", "4")

	g.Set("7_7", "3")
	g.Set("7_8", "4")
	g.Set("7_16", "4")
	g.Set("7_25", "1")
	g.Set("7_27", "9")
	g.Set("8_7", "8")
	g.Set("8_9", "7")
	g.Set("8_18", "1")
	g.Set("9_18", "5")

	g.SetLabel("3_2", "D")
	g.SetLabel("4_14", "G")
	g.SetLabel("5_24", "I")
	g.SetLabel("3_19", "E")

	g.SetLabel("9_16", "H")
	g.SetLabel("9_25", "F")

	return g, nil
}
