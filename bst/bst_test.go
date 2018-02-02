
package bst

import (
    "testing"
    "fmt"
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


// func TestSearch(t *testing.T)