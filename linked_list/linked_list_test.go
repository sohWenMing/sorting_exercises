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
