package linkedlist

import (
	"testing"
)

func TestInitAddToHead(t *testing.T) {
	testLL := InitLinkedList()

	testHeadNode := InitNode(1)
	testLL.AddToHead(testHeadNode)
	// at this point list shoud be empty. so head and tail should
	// just be the newly insterted node

	checkNodes(testLL.GetHead(), testHeadNode, t)
	// checking that head was properly insterted into head of empty list
	checkNodes(testLL.GetTail(), testHeadNode, t)
	// checking that tail was properly insterted into tail of empty list

	checkLength(t, testLL, 1)
	// length of list should be one after adding head to empty list

	secondTestHeadNode := InitNode(2)
	testLL.AddToHead(secondTestHeadNode)
	checkNodes(testLL.GetHead(), secondTestHeadNode, t)
	checkNodes(testLL.GetTail(), testHeadNode, t)
	checkLength(t, testLL, 2)
}

func TestAddToTail(t *testing.T) {
	testLL := InitLinkedList()

	testHeadNode := InitNode(1)
	testLL.AddToTail(testHeadNode)
	// at this point list shoud be empty. so head and tail should
	// just be the newly insterted node

	checkNodes(testLL.GetHead(), testHeadNode, t)
	// checking that head was properly insterted into head of empty list
	checkNodes(testLL.GetTail(), testHeadNode, t)
	// checking that tail was properly insterted into tail of empty list

	checkLength(t, testLL, 1)
	// length of list should be one after adding head to empty list

	secondTestHeadNode := InitNode(2)
	testLL.AddToTail(secondTestHeadNode)
	checkNodes(testLL.GetHead(), testHeadNode, t)
	checkNodes(testLL.GetTail(), secondTestHeadNode, t)
	checkLength(t, testLL, 2)
}

func TestGetNodeAtIndex(t *testing.T) {
	testLL := InitLinkedList()
	_, isFound := testLL.GetNodeAtIndex(-1)
	if isFound {
		t.Errorf("isFound should have returned false but returned true")
	}
	// test for trying to get node at negative index, should return isFound = false

	_, isFound = testLL.GetNodeAtIndex(0)
	if isFound {
		t.Errorf("isFound should have returned false but returned true")
	}
	// test for trying to get node at first index of empty list, should return isFound = false
	nodes := []*Node{}
	for i := range 5 {
		newNode := InitNode(i)
		nodes = append(nodes, newNode)
		testLL.AddToTail(newNode)
	}

	gotNode, isFound := testLL.GetNodeAtIndex(4)
	if !isFound {
		t.Errorf("isFound should have returned true but returned false")
	}
	checkNodes(gotNode, nodes[4], t)

}

func checkLength(t *testing.T, list LinkedList, wantLength int) {
	gotLength := list.GetLength()
	if gotLength != wantLength {
		t.Errorf("got %v want %v", gotLength, wantLength)
	}
}

func checkNodes(gotNode, wantNode *Node, t *testing.T) {
	got := gotNode.String()
	want := wantNode.String()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
