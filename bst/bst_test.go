package bst

import (
	"fmt"
	"testing"
)

func TestArrayPrint(t *testing.T) {
	tree1 := ExampleTree()
	PrintTreeArray(&tree1)
}

func TestEqualTrees(t *testing.T) {

	// Create and define the example tree for this problem.
	tree1 := ExampleTree()
	tree2 := ExampleTree2()

	answer1 := AreTreesEqual(&tree1, &tree1)
	answer2 := AreTreesEqual(&tree1, &tree2)

	fmt.Println(answer1, answer2)
	if !answer1 {
		t.Error("the trees are identical, but program said they weren't.")
	}
	if answer2 {
		t.Error(`tree1 and tree2 are different, but program said equal`)
	}

}

func TestSearch(t *testing.T) {
	tree1 := ExampleTree()
	branch := TreeSearch(&tree1, 8)
	PrintTreeWalk(&branch)
}

func TestMinMax(t *testing.T) {

	tree1 := ExampleTree()
	tree2 := ExampleTree2()

	a, b := TreeMinimum(&tree1).Value, TreeMinimum(&tree2).Value
	if a != 1 || b != 1 {
		t.Error("Minimum of examples should be equal to 1.")
	}

	c, d := TreeMaximum(&tree1).Value, TreeMaximum(&tree2).Value
	if c != 9 || d != 9 {
		t.Error("Maximum of examples should be equal to 9.")
	}

	fmt.Println("Mins:", a, b, "Maxes:", c, d)

}

// func TestSearch(t *testing.T)
