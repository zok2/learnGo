package main

import "errors"

var (
	ErrEmpty = errors.New("stack is empty")
)

type Node struct {
	data interface{}
	next *Node
}

type LinkStack struct {
	top    *Node
	length int
}

func InitStack() *LinkStack  {
	return new(LinkStack)
}

func (s *LinkStack) IsEmpty() bool  {
	if s.length == 0 {
		return true
	}else {
		return false
	}
}

func (s *LinkStack)Push(data interface{})  {
	item := Node{data: data,next: s.top}
	s.top= &item
	s.length++
}

func (s *LinkStack) Pop() (data interface{},err error)  {
	if s.IsEmpty() {
		return false ,ErrEmpty
	}else {
		s.length --
		data = s.top.data
		s.top = s.top.next
		return data,errors.New("")
	}
}

func (s *LinkStack) Length () int  {
	return s.length
}
