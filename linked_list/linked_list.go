package linkedlist

import (
	"fmt"
	"strings"
)

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

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func InitLinkedList() LinkedList {
	return LinkedList{
		nil,
		nil,
		0,
	}
}

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
func (l *LinkedList) GetLength() int {
	return l.length
}

func (l *LinkedList) incSize() {
	l.length++
}
