/*package bst has implementations of Binary Search Trees.

Important Notes:

- Tree{0, nil, nil} is the default empty Tree structure, and is therefore
equivalent to a nil Tree.  So... reserve Tree.value = 0  for only nil trees.

*/
package bst

import (
    "fmt"
)



/*The basic tree structure has an integer for a value.  Like most trees,
it has "children" that can also be defined as trees. */
type Tree struct {
    Value int
    Left *Tree
    Right *Tree
}


/*TreeSearch returns a tree with a root node of the given value.
If the desired value cannot be found, returns a Tree{0, nil, nil}, 
which is a tree with value = 0, and.

Example

    output := TreeSearch(input, 8)

          input             output

            3                  8
           / \                / \
          2   7              5   9
         /   / \
        1   4   8
               / \
              5   9

*/
func TreeSearch(t *Tree, value int) Tree {
    for (value != t.Value) {
        if (t.Left == nil && t.Right == nil) {
            return Tree{}
        }
        if value < t.Value {
            t = t.Left
            continue
        }
        t = t.Right
    }
    return *t
}

/*InOrderTreeWalk performs an exhaustive walk through every node in the tree.
Assuming the tree is a binary search tree, each value sent
along the channel will be greater than the previous value.  */
func InOrderTreeWalk(t *Tree, channel chan int) {
    if (t != nil) {
        InOrderTreeWalk(t.Left, channel)
        channel <- t.Value
        InOrderTreeWalk(t.Right, channel)
    }
}



// NOTE:  HIDDEN FUNCTION
// Supports AreTreesEqual.
func tree_walker (t *Tree) <-chan int {
    channel := make(chan int)
    go func() {
        InOrderTreeWalk(t, channel)
        close(channel)
    }()
    return channel
}

/*AreTreesEqual compares two trees for equality by taking 1 step at a time
through each tree, and comparing the values.  At any given step, if the
values aren't equal, then the walkers stop and the function returns false.

Credits: slightly modified from: https://golang.org/doc/play/tree.go*/
func AreTreesEqual(t1, t2 *Tree) bool {

    // Walker function returns a channel, receiving values from an
    // In-Order Tree Walk. Channel closes when there are no values left.
    Walker := func (t *Tree) <-chan int {
        channel := make(chan int)
        go func() {
            InOrderTreeWalk(t, channel)
            close(channel)
        }()
        return channel
    }

    c1, c2 := Walker(t1), Walker(t2)

    for {

        // blocks; waits for both channels for either values or closes.
        value1, isChanOpen1 := <- c1
        value2, isChanOpen2 := <- c2

        /*Check: is either channel closed?

        At this point, We assume: all previous values in tree1 & tree2 equal.
        Therefore, if both channels are closed, then there are no more values.
        If all values are equal, Then the trees are equal.  

        If one channel is closed, but the other sends another value,
        then the two trees have a different quantity of nodes.
        Therefore the two trees are not equal.
        [QED] */
        if !isChanOpen1 || !isChanOpen2 {
            return isChanOpen1 == isChanOpen2
        }

        if value1 != value2 {
            break
        }
    }
    return false 
}




















/*printTreeWalk does a simple In-Order Tree Walk on a Binary Search Tree, and
prints out all of the values to standard fmt output.*/
func PrintTreeWalk(t *Tree) {
    if (t != nil) {
        PrintTreeWalk(t.Left)
        fmt.Println(t.Value)
        PrintTreeWalk(t.Right)
    }
}

//printTreeArray prints out all the values of a tree in-order as an array.
func PrintTreeArray(t *Tree) {

    var output []int 
    c := tree_walker(t)
    for {
        value, isChanOpen := <- c
        if isChanOpen {
            output = append(output, value)
            continue
        }
        break
    }
    fmt.Println(output)
}


/*isBST checks if a Tree is follows the Binary Search Tree Properties.
returns True if all requirment conditions are met, False if not.

In order for a Tree to be considered a Binary Search Tree, the values of
all nodes to the left of a parent node must be smaller, and all values of the 
nodes to right must be larger.*/
func isBST(t *Tree) {

}


// AreIntSlicesEqual compares two integer slices for equality.
func AreIntSlicesEqual(a, b []int) bool {
    if a == nil && b == nil {
        return true
    }
    if a == nil || b == nil {
        return false
    }
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}




/*exampleTree returns a specific example of a Binary Search Tree structure.
This is the tree that exampleTree creates:

            3
           / \
          2   7
         /   / \
        1   4   8
               / \
              5   9

*/
func ExampleTree() Tree {
    n1 := Tree{1, nil, nil}
    n2 := Tree{2, &n1, nil}
    n4 := Tree{4, nil, nil}
    n5 := Tree{5, nil, nil}
    n9 := Tree{9, nil, nil}
    n8 := Tree{8, &n5, &n9}
    n7 := Tree{7, &n4, &n8}
    pn := Tree{3, &n2, &n7}
    return pn
}

/*
ExampleTree2

slightly similar, but different.  Tries to get a walker caught up.

            3
           / \
          2   7
         /     \
        1       8
                 \
                  9

*/
func ExampleTree2() Tree {
    n1 := Tree{1, nil, nil}
    n2 := Tree{2, &n1, nil}
    n9 := Tree{9, nil, nil}
    n8 := Tree{8, nil, &n9}
    n7 := Tree{7, nil, &n8}
    pn := Tree{3, &n2, &n7}
    return pn
}




/*Under advice of [legal council], I have been advised not to comment on this
function.*/
func NoComment() string {
    return "No comment"
}



