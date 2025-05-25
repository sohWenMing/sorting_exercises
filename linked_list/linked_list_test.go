package linkedlist

import (
	"errors"
	"fmt"
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
	nodes := addNotestoList(testLL, 5)
	gotNode, isFound := testLL.GetNodeAtIndex(4)
	if !isFound {
		nodes := []*Node{}
		for i := range 5 {
			newNode := InitNode(i)
			nodes = append(nodes, newNode)
			testLL.AddToTail(newNode)
		}
		t.Errorf("isFound should have returned true but returned false")
	}
	checkNodes(gotNode, nodes[4], t)
}

func TestGetNodeWithVal(t *testing.T) {
	testLL := InitLinkedList()
	nodes := addNotestoList(testLL, 5)
	_, index, isFound := testLL.GetNodeWithVal(-1)
	if isFound {
		t.Errorf("isFound should have returned false, returned true \nnodes: %v", nodes)
		return
	}
	if index != 0 {
		t.Errorf("index should have returned 0, returned %d", index)
		return
	}

	checkVals := []int{0, 3, 4}
	for _, val := range checkVals {
		node, index, isFound := testLL.GetNodeWithVal(val)
		checkNodes(node, nodes[val], t)
		if !isFound {
			t.Errorf("search for node with val %d returned isFound == false\n nodes: %v", val, nodes)
			return
		}
		if index != val {
			t.Errorf("index was not correct for value searched %d\n nodes: %v", val, nodes)
			return
		}
	}
}

func TestAddNodeGuardChecks(t *testing.T) {
	// Index is a negative index
	testLL := InitLinkedList()
	addNotestoList(testLL, 5)
	nodeToAdd := InitNode(-1)

	err := testLL.AddNodeAtIndex(nodeToAdd, -1)
	if err == nil {
		t.Errorf("expected error, didn't get one")
	}
	err = testLL.AddNodeAtIndex(nodeToAdd, 6)
	if err == nil {
		t.Errorf("expected error, didn't get one")
	}
}

func TestAddNodeAtHeadByIndex(t *testing.T) {
	// Test 2: Adding to the first index of a list
	testLL := InitLinkedList()
	nodeToAdd1 := InitNode(-1)

	err := testLL.AddNodeAtIndex(nodeToAdd1, 0)
	if err != nil {
		t.Errorf("didn't expect error, got %v", err)
	}
	gotHead := testLL.GetHead()
	checkNodes(gotHead, nodeToAdd1, t)
	gotTail := testLL.GetTail()
	checkNodes(gotTail, nodeToAdd1, t)
	checkLength(t, testLL, 1)

	nodeToAdd2 := InitNode(-2)
	err = testLL.AddNodeAtIndex(nodeToAdd2, 0)
	if err != nil {
		t.Errorf("didn't expect error, got %v", err)
	}
	gotHead = testLL.GetHead()
	checkNodes(gotHead, nodeToAdd2, t)
	gotTail = testLL.GetTail()
	checkNodes(gotTail, nodeToAdd1, t)
	checkLength(t, testLL, 2)
}

func TestAddNodeAtTailByIndex(t *testing.T) {

	testLL := InitLinkedList()
	addNotestoList(testLL, 5)
	nodeToAdd := InitNode(-1)

	err := testLL.AddNodeAtIndex(nodeToAdd, 5)
	if err != nil {
		t.Errorf("didn't expect error, got %v", err)
	}
	gotTail := testLL.GetTail()
	checkNodes(gotTail, nodeToAdd, t)
	checkLength(t, testLL, 6)
}

func TestAddNodeAtMiddleByIndex(t *testing.T) {
	testLL := InitLinkedList()
	addNotestoList(testLL, 5)
	// curr state of list {0,1,2,3,4}
	nodeToAdd := InitNode(-1)
	err := testLL.AddNodeAtIndex(nodeToAdd, 2)
	// expected state of list would be {0, 1, -1, 2, 3, 4}
	if err != nil {
		t.Errorf("didn't expect error, got %v", err)
	}
	gotNode, isFound := testLL.GetNodeAtIndex(2)
	if !isFound {
		t.Errorf("is found returned false should return true")
		stringReps := testLL.GetStringReprs()
		fmt.Println("Nodes: ")
		for _, stringRep := range stringReps {
			fmt.Println(stringRep)
		}
		return
	}
	err = checkNodes(gotNode, nodeToAdd, t)
	if err != nil {
		stringReps := testLL.GetStringReprs()
		fmt.Println("Nodes: ")
		for _, stringRep := range stringReps {
			fmt.Println(stringRep)
		}
	}
}

func TestGetStringReprs(t *testing.T) {
	testLL := InitLinkedList()
	addNotestoList(testLL, 5)
	stringReps := testLL.GetStringReprs()
	for _, stringRep := range stringReps {
		fmt.Println(stringRep)
	}

}

func addNotestoList(l *LinkedList, numNodes int) []*Node {
	nodes := []*Node{}
	for i := range numNodes {
		newNode := InitNode(i)
		nodes = append(nodes, newNode)
		l.AddToTail(newNode)
	}
	return nodes
}

func checkLength(t *testing.T, list *LinkedList, wantLength int) {
	gotLength := list.GetLength()
	if gotLength != wantLength {
		t.Errorf("got %v want %v", gotLength, wantLength)
	}
}

func checkNodes(gotNode, wantNode *Node, t *testing.T) error {
	got := gotNode.String()
	want := wantNode.String()
	if got != want {
		t.Errorf("\ngot \n%v \nwant \n%v", got, want)
		return errors.New("error returned")
	}
	return nil
}
