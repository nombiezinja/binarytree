package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	x := NewNode()

	assert.Nil(t, x.Left)
	assert.Nil(t, x.Right)
}

func TestNewNodeKeyValue(t *testing.T) {
	x := NewNodeKeyValue(IntKey(2), "Help!")

	assert.Equal(t, x.Key, IntKey(2))
	assert.Equal(t, x.Value, "Help!")
}

func TestFind(t *testing.T) {
	Root := getTestTreeBalanced(1)

	assert.Equal(t, Root.Find(IntKey(4)), Root)
	assert.Equal(t, Root.Find(IntKey(6)), Root.Right)
	assert.Equal(t, Root.Find(IntKey(2)), Root.Left)
	assert.Equal(t, Root.Find(IntKey(5)), Root.Right.Left)
	assert.Equal(t, Root.Find(IntKey(7)), Root.Right.Right)
	assert.Equal(t, Root.Find(IntKey(1)), Root.Left.Left)
	assert.Equal(t, Root.Find(IntKey(3)), Root.Left.Right)

	assert.Nil(t, Root.Find(IntKey(9)))
	assert.Nil(t, Root.Find(IntKey(-1)))
}

func TestCopy(t *testing.T) {
	Root1 := getTestTreeBalanced(1)

	Root2 := Root1.Copy()

	Root2.Remove(IntKey(3))
	Root2.Remove(IntKey(4))
	Root2.Remove(IntKey(5))

	assert.NotNil(t, Root1.Find(IntKey(3)))
	assert.NotNil(t, Root1.Find(IntKey(4)))
	assert.NotNil(t, Root1.Find(IntKey(5)))
}

func TestPrevious(t *testing.T) {
	var Root, node *Node

	// Left Unbalanced
	Root = getTestTreeLeftUnbalanced(1)
	assert.Equal(t, IntKey(7), Root.Previous(IntKey(10)).Key)
	for i := 6; i > 1; i-- {
		node = Root.Previous(IntKey(i))
		assert.NotNil(t, node)
		if node != nil {
			assert.Equal(t, IntKey(i-1), node.Key)
		}
	}
	assert.Nil(t, Root.Previous(IntKey(1)))

	// Right Unbalanced
	Root = getTestTreeRightUnbalanced(1)
	assert.Equal(t, IntKey(7), Root.Previous(IntKey(10)).Key)
	for i := 6; i > 1; i-- {
		node = Root.Previous(IntKey(i))
		assert.NotNil(t, node)
		if node != nil {
			assert.Equal(t, IntKey(i-1), node.Key)
		}
	}
	assert.Nil(t, Root.Previous(IntKey(1)))

	// Balanced
	Root = getTestTreeBalanced(1)
	assert.Equal(t, IntKey(7), Root.Previous(IntKey(10)).Key)
	for i := 6; i > 1; i-- {
		node = Root.Previous(IntKey(i))
		assert.NotNil(t, node)
		if node != nil {
			assert.Equal(t, IntKey(i-1), node.Key)
		}
	}
	assert.Nil(t, Root.Previous(IntKey(1)))

	// Left Unbalanced Scaled (3)
	Root = getTestTreeLeftUnbalanced(3)
	assert.Equal(t, IntKey(21), Root.Previous(IntKey(30)).Key)
	for i := 18; i > 3; i-- {
		node = Root.Previous(IntKey(i))
		assert.NotNil(t, node)
		if node != nil {
			assert.Equal(t, IntKey(((i-1)/3)*3), node.Key)
		}
	}
	assert.Nil(t, Root.Previous(IntKey(1)))

	// Right Unbalanced Scaled (3)
	Root = getTestTreeRightUnbalanced(3)
	assert.Equal(t, IntKey(21), Root.Previous(IntKey(30)).Key)
	for i := 18; i > 3; i-- {
		node = Root.Previous(IntKey(i))
		assert.NotNil(t, node)
		if node != nil {
			assert.Equal(t, IntKey(((i-1)/3)*3), node.Key)
		}
	}
	assert.Nil(t, Root.Previous(IntKey(1)))

	// Balanced Scaled (3)
	Root = getTestTreeBalanced(3)
	assert.Equal(t, IntKey(21), Root.Previous(IntKey(30)).Key)
	for i := 18; i > 3; i-- {
		node = Root.Previous(IntKey(i))
		assert.NotNil(t, node)
		if node != nil {
			assert.Equal(t, IntKey(((i-1)/3)*3), node.Key)
		}
	}
	assert.Nil(t, Root.Previous(IntKey(1)))

	// // Single node should search higher should return nil
	// Root = NewNodeKeyValue(IntKey(7),"seven")
	// node = Root.Previous(IntKey(8))
	// assert.Nil(t, node)
}

func TestNext(t *testing.T) {
	var Root, node *Node

	// Left Unbalanced
	Root = getTestTreeLeftUnbalanced(1)
	assert.Equal(t, IntKey(1), Root.Next(IntKey(0)).Key)
	for i := 0; i < 6; i++ {
		node = Root.Next(IntKey(i))
		assert.NotNil(t, node)
		if node != nil {
			assert.Equal(t, IntKey(i+1), node.Key)
		}
	}
	assert.Nil(t, Root.Next(IntKey(7)))

	// Right Unbalanced
	Root = getTestTreeRightUnbalanced(1)
	assert.Equal(t, IntKey(1), Root.Next(IntKey(0)).Key)
	for i := 0; i < 6; i++ {
		node = Root.Next(IntKey(i))
		assert.NotNil(t, node)
		if node != nil {
			assert.Equal(t, IntKey(i+1), node.Key)
		}
	}
	assert.Nil(t, Root.Next(IntKey(7)))

	// Balanced
	Root = getTestTreeBalanced(1)
	assert.Equal(t, IntKey(1), Root.Next(IntKey(0)).Key)
	for i := 0; i < 6; i++ {
		node = Root.Next(IntKey(i))
		assert.NotNil(t, node)
		if node != nil {
			assert.Equal(t, IntKey(i+1), node.Key)
		}
	}
	assert.Nil(t, Root.Next(IntKey(7)))

	// Left Unbalanced Scaled (3)
	Root = getTestTreeLeftUnbalanced(3)
	assert.Equal(t, IntKey(3), Root.Next(IntKey(0)).Key)
	for i := 0; i < 21; i++ {
		node = Root.Next(IntKey(i))
		assert.NotNil(t, node)
		if node != nil {
			assert.Equal(t, IntKey(((i/3)+1)*3), node.Key)
		}
	}
	assert.Nil(t, Root.Next(IntKey(21)))

	// Right Unbalanced Scaled (3)
	Root = getTestTreeRightUnbalanced(3)
	assert.Equal(t, IntKey(3), Root.Next(IntKey(0)).Key)
	for i := 0; i < 21; i++ {
		node = Root.Next(IntKey(i))
		assert.NotNil(t, node)
		if node != nil {
			assert.Equal(t, IntKey(((i/3)+1)*3), node.Key)
		}
	}
	assert.Nil(t, Root.Next(IntKey(21)))

	// Balanced Scaled (3)
	Root = getTestTreeBalanced(3)
	assert.Equal(t, IntKey(3), Root.Next(IntKey(0)).Key)
	for i := 0; i < 21; i++ {
		node = Root.Next(IntKey(i))
		assert.NotNil(t, node)
		if node != nil {
			assert.Equal(t, IntKey(((i/3)+1)*3), node.Key)
		}
	}
	assert.Nil(t, Root.Next(IntKey(21)))

	// Single node should search higher should return nil
	Root = NewNodeKeyValue(IntKey(7), "seven")
	node = Root.Next(IntKey(8))
	assert.Nil(t, node)
}

func TestNodeAdd(t *testing.T) {
	x := NewNodeKeyValue(IntKey(5), "five")
	y := NewNodeKeyValue(IntKey(2), "two")
	z := NewNodeKeyValue(IntKey(7), "seven")
	q := NewNodeKeyValue(IntKey(9), "nine")

	x.Add(y)
	assert.Equal(t, x.Left, y)
	assert.Nil(t, x.Right)
	x.Add(z)
	assert.Equal(t, x.Left, y)
	assert.Equal(t, x.Right, z)
	x.Add(q)
	assert.Equal(t, x.Left, y)
	assert.Equal(t, x.Right, z)
	assert.Equal(t, x.Right.Right, q)
}

func TestRemove(t *testing.T) {
	// Remove only node
	Root := NewNodeKeyValue(IntKey(2), "two")
	Root = Root.Remove(IntKey(2))
	assert.Nil(t, Root)

	// Remove this node with right child
	Root = NewNodeKeyValue(IntKey(2), "two")
	other := NewNodeKeyValue(IntKey(4), "four")
	Root.Add(other)
	Root = Root.Remove(IntKey(2))
	assert.Equal(t, Root, other)

	// Remove this node with left child
	Root = NewNodeKeyValue(IntKey(2), "two")
	other = NewNodeKeyValue(IntKey(1), "one")
	Root.Add(other)
	Root = Root.Remove(IntKey(2))
	assert.Equal(t, Root, other)

	// Remove this node with both children child
	Root = NewNodeKeyValue(IntKey(2), "two")
	other1 := NewNodeKeyValue(IntKey(1), "one")
	other4 := NewNodeKeyValue(IntKey(4), "four")
	Root.Add(other1)
	Root.Add(other4)

	Root = Root.Remove(IntKey(2))
	assert.Equal(t, Root, other1)
	assert.Equal(t, Root.Right, other4)
	assert.Nil(t, Root.Left)

	// Brute Force Test
	Root = getTestTreeBalanced(1)

	Root = Root.Remove(IntKey(2))
	Root = Root.Remove(IntKey(3))
	Root = Root.Remove(IntKey(6))

	assert.NotNil(t, Root.Find(IntKey(1)))
	assert.Nil(t, Root.Find(IntKey(2)))
	assert.Nil(t, Root.Find(IntKey(3)))
	assert.NotNil(t, Root.Find(IntKey(4)))
	assert.NotNil(t, Root.Find(IntKey(5)))
	assert.Nil(t, Root.Find(IntKey(6)))
	assert.NotNil(t, Root.Find(IntKey(7)))
}

func TestBalance(t *testing.T) {
	Root := getTestTreeRightUnbalanced(1)
	Root = Root.Balance()

	assert.Equal(t, IntKey(4), Root.Key)
	assert.Equal(t, IntKey(2), Root.Left.Key)
	assert.Equal(t, IntKey(6), Root.Right.Key)
	assert.Equal(t, IntKey(1), Root.Left.Left.Key)
	assert.Equal(t, IntKey(3), Root.Left.Right.Key)
	assert.Equal(t, IntKey(5), Root.Right.Left.Key)
	assert.Equal(t, IntKey(7), Root.Right.Right.Key)

	Root = getTestTreeLeftUnbalanced(1)
	Root = Root.Balance()

	assert.Equal(t, IntKey(4), Root.Key)
	assert.Equal(t, IntKey(2), Root.Left.Key)
	assert.Equal(t, IntKey(6), Root.Right.Key)
	assert.Equal(t, IntKey(1), Root.Left.Left.Key)
	assert.Equal(t, IntKey(3), Root.Left.Right.Key)
	assert.Equal(t, IntKey(5), Root.Right.Left.Key)
	assert.Equal(t, IntKey(7), Root.Right.Right.Key)

	Root = getTestTreeBalanced(1)
	Root = Root.Balance()

	assert.Equal(t, IntKey(4), Root.Key)
	assert.Equal(t, IntKey(2), Root.Left.Key)
	assert.Equal(t, IntKey(6), Root.Right.Key)
	assert.Equal(t, IntKey(1), Root.Left.Left.Key)
	assert.Equal(t, IntKey(3), Root.Left.Right.Key)
	assert.Equal(t, IntKey(5), Root.Right.Left.Key)
	assert.Equal(t, IntKey(7), Root.Right.Right.Key)
}

// Internals

func TestMinimum(t *testing.T) {
	Root := getTestTreeBalanced(1)
	assert.Equal(t, Root.Minimum().Key, IntKey(1))

	Root = getTestTreeLeftUnbalanced(1)
	assert.Equal(t, Root.Minimum().Key, IntKey(1))

	Root = getTestTreeRightUnbalanced(1)
	assert.Equal(t, Root.Minimum().Key, IntKey(1))
}

func TestMaximum(t *testing.T) {
	Root := getTestTreeBalanced(1)
	assert.Equal(t, Root.Maximum().Key, IntKey(7))

	Root = getTestTreeLeftUnbalanced(1)
	assert.Equal(t, Root.Maximum().Key, IntKey(7))

	Root = getTestTreeRightUnbalanced(1)
	assert.Equal(t, Root.Maximum().Key, IntKey(7))
}

func TestDepthLeft(t *testing.T) {
	Root := getTestTreeRightUnbalanced(1)
	assert.Equal(t, 0, Root.DepthLeft())

	Root = getTestTreeLeftUnbalanced(1)
	assert.Equal(t, 6, Root.DepthLeft())

	Root = getTestTreeBalanced(1)
	assert.Equal(t, 2, Root.DepthLeft())
}

func TestDepthRight(t *testing.T) {
	Root := getTestTreeRightUnbalanced(1)
	assert.Equal(t, 6, Root.DepthRight())

	Root = getTestTreeLeftUnbalanced(1)
	assert.Equal(t, 0, Root.DepthRight())

	Root = getTestTreeBalanced(1)
	assert.Equal(t, 2, Root.DepthRight())
}

func TestWalkForward(t *testing.T) {
	Root := getTestTreeBalanced(1)

	out := []string{}

	Root.WalkForward(func(me *Node) {
		out = append(out, me.Value.(string))
	})

	assert.Equal(t, out, []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
	})
}

func TestWalkBackward(t *testing.T) {
	Root := getTestTreeBalanced(1)

	out := []string{}

	Root.WalkBackward(func(me *Node) {
		out = append(out, me.Value.(string))
	})

	assert.Equal(t, out, []string{
		"seven",
		"six",
		"five",
		"four",
		"three",
		"two",
		"one",
	})
}

// Helpers

func getTestTreeRightUnbalanced(factor int) *Node {
	Root := NewNodeKeyValue(IntKey(1*factor), "one")
	Root.Add(NewNodeKeyValue(IntKey(2*factor), "two"))
	Root.Add(NewNodeKeyValue(IntKey(3*factor), "three"))
	Root.Add(NewNodeKeyValue(IntKey(4*factor), "four"))
	Root.Add(NewNodeKeyValue(IntKey(5*factor), "five"))
	Root.Add(NewNodeKeyValue(IntKey(6*factor), "six"))
	Root.Add(NewNodeKeyValue(IntKey(7*factor), "seven"))
	return Root
}

func getTestTreeLeftUnbalanced(factor int) *Node {
	Root := NewNodeKeyValue(IntKey(7*factor), "seven")
	Root.Add(NewNodeKeyValue(IntKey(6*factor), "six"))
	Root.Add(NewNodeKeyValue(IntKey(5*factor), "five"))
	Root.Add(NewNodeKeyValue(IntKey(4*factor), "four"))
	Root.Add(NewNodeKeyValue(IntKey(3*factor), "three"))
	Root.Add(NewNodeKeyValue(IntKey(2*factor), "two"))
	Root.Add(NewNodeKeyValue(IntKey(1*factor), "one"))
	return Root
}

func getTestTreeBalanced(factor int) *Node {
	Root := getTestTreeLeftUnbalanced(factor)
	return Root.Balance()
}
