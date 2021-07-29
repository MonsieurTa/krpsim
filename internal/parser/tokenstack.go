package parser

import (
	"fmt"
	"sync"

	"github.com/MonsieurTa/go-lexer"
)

type Stack struct {
	mux        sync.Mutex
	head, tail *StackNode
	size       int
}

type StackNode struct {
	Val  lexer.Token
	Next *StackNode
}

func (sn StackNode) IsType(typ lexer.TokenType) bool {
	return sn.Val.Type() == typ
}

func (s *Stack) PushFront(n *StackNode) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.head == nil && s.tail == nil {
		s.head = n
		s.tail = n
	} else {
		n.Next = s.head
		s.head = n
	}
	s.size++
}

func (s *Stack) BatchPushFront(n []*StackNode) {
	for i := len(n) - 1; i >= 0; i-- {
		s.PushFront(n[i])
	}
}

func (s *Stack) PushBack(v lexer.Token) {
	s.mux.Lock()
	defer s.mux.Unlock()

	n := &StackNode{v, nil}
	if s.head == nil {
		s.head = n
		s.tail = n
	} else {
		s.tail.Next = n
		s.tail = n
	}
	s.size++
}

func (s *Stack) PopFront() *StackNode {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.head == nil {
		return nil
	}
	rv := s.head
	s.head = s.head.Next
	rv.Next = nil

	s.size--
	if s.size == 0 {
		s.head = nil
		s.tail = nil
	}
	return rv
}

func (s *Stack) Peek() *StackNode {
	rv := s.PopFront()
	if rv == nil {
		return nil
	}
	s.PushFront(rv)
	return rv
}

func (s *Stack) Accept(valid []lexer.TokenType) bool {
	v := s.PopFront()
	for _, t := range valid {
		if v.Val.Type() == t {
			return true
		}
	}
	s.PushFront(v)
	return false
}

func (s *Stack) IgnoreIf(valid []lexer.TokenType) {
	if len(valid) == 0 {
		return
	}
	for s.Accept(valid) {
	}
}

func (s *Stack) PrintNext(n int) {
	head := s.head
	for head != nil && n > 0 {
		fmt.Println(head.Val.Value())
		head = head.Next
		n--
	}
}
