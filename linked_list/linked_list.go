package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

// ------------Node Definitions and Methods ------------------
type Node struct {
	val  int
	next *Node
}

func InitNode(val int) *Node {
	return &Node{
		val,
		nil,
	}
}

func (n *Node) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(
		"node val: %d\n", n.val,
	))
	switch n.next != nil {
	case true:
		b.WriteString(fmt.Sprintf(
			"node next val: %d\n", n.next.val,
		))
	case false:
		b.WriteString(fmt.Sprint(
			"node next val: nil\n",
		))
	}
	return b.String()
}

// ------------Linked List Definitions and Methods ------------------

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func InitLinkedList() *LinkedList {
	return &LinkedList{
		nil,
		nil,
		0,
	}
}

// define a method to add a new node to the head of the list
func (l *LinkedList) AddToHead(node *Node) {
	switch l.head == nil {
	case true:
		l.SetHead(node)
		l.SetTail(node)
	case false:
		cur := l.head
		node.next = cur
		l.SetHead(node)

	}
	l.incSize()
	return
}

// define a method to add a new node to the tail of the list
func (l *LinkedList) AddToTail(node *Node) {
	switch l.head == nil {
	case true:
		l.SetHead(node)
		l.SetTail(node)
	case false:
		curTail := l.GetTail()
		curTail.next = node
		l.SetTail(node)
	}
	l.incSize()
	return
}

func (l *LinkedList) GetNodeAtIndex(searchIndex int) (node *Node, isFound bool) {
	node = nil
	isFound = false

	if searchIndex < 0 || searchIndex > l.GetLength()-1 {
		return
	}
	idx := 0
	cur := l.GetHead()

	for idx < searchIndex {
		cur = cur.next
		idx++
	}
	return cur, true
}

func (l *LinkedList) GetNodeWithVal(searchVal int) (node *Node, index int, isFound bool) {
	node = nil
	index = 0
	isFound = false

	// init negative values first
	if l.GetLength() == 0 {
		return
	}

	curNode := l.GetHead()
	curIndex := 0
	for curNode != nil {

		if curNode.val == searchVal {
			node = curNode
			index = curIndex
			isFound = true
			return
		}

		curNode = curNode.next
		curIndex += 1
	}
	return
}

func (l *LinkedList) AddNodeAtIndex(node *Node, index int) (err error) {
	// handle cases for negative index, and out of bounds error
	if index < 0 {
		return errors.New("index cannot be negative")
	}

	curLength := l.GetLength()
	if index > curLength {
		return errors.New("index would be be out of bounds")
	}

	// if inserting at index 0, then we can just use the AddToHead method which will already handle head and tail
	if index == 0 {
		l.AddToHead(node)
		return nil
	}
	// if inserting after last index of list, then can just use the AddToTail method which will already handle head and tail
	if index == curLength {
		l.AddToTail(node)
		return nil
	}

	// if inserting at middle, will need to find the previous node to attach the new node to, and the next node to attach the new node to
	currNode := l.GetHead()
	var prevNode *Node
	currIdx := 0

	for currIdx < index {
		prevNode = currNode
		currNode = currNode.next
		currIdx++
	}

	prevNode.next = node
	node.next = currNode

	return nil
}

func (l *LinkedList) GetStringReprs() []string {
	returnSlice := []string{}
	currNode := l.GetHead()
	for currNode != nil {
		returnSlice = append(returnSlice, currNode.String())
		currNode = currNode.next
	}
	return returnSlice
}

func (l *LinkedList) GetLength() int {
	return l.length
}

func (l *LinkedList) SetHead(node *Node) {
	l.head = node
	return
}
func (l *LinkedList) SetTail(node *Node) {
	l.tail = node
	return
}

func (l *LinkedList) GetHead() *Node {
	return l.head
}
func (l *LinkedList) GetTail() *Node {
	return l.tail
}

func (l *LinkedList) incSize() {
	l.length++
}
