// CookieJar - A contestant's algorithm toolbox
// Copyright (c) 2013 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: use of this source code is governed by a BSD
// license that can be found in the LICENSE file. Alternatively, the CookieJar
// toolbox may be used in accordance with the terms and conditions contained
// in a signed written agreement between you and the author(s).

// Package stack implements a LIFO (last in first out) data structure supporting
// arbitrary types (even a mixture).
//
// Internally it uses a dynamically growing slice of blocks, resulting in faster
// resizes than a simple dynamic array/slice would allow.
package main

// The size of a block of data
const blockSize = 4096

// Stack is a last in, first out data structure.
type Stack struct {
	size     int
	capacity int
	offset   int

	blocks [][]interface{}
	active []interface{}
}

// NewStack creates a new, empty stack.
func NewStack() *Stack {
	result := new(Stack)
	result.active = make([]interface{}, blockSize)
	result.blocks = [][]interface{}{result.active}
	result.capacity = blockSize
	return result
}

// Push es a value onto the stack, expanding it if necessary.
func (s *Stack) Push(data interface{}) {
	if s.size == s.capacity {
		s.active = make([]interface{}, blockSize)
		s.blocks = append(s.blocks, s.active)
		s.capacity += blockSize
		s.offset = 0
	} else if s.offset == blockSize {
		s.active = s.blocks[s.size/blockSize]
		s.offset = 0
	}
	s.active[s.offset] = data
	s.offset++
	s.size++
}

// Pop s a value off the stack and returns it. Currently no shrinking is done.
func (s *Stack) Pop() (res interface{}) {
	s.size--
	s.offset--
	if s.offset < 0 {
		s.offset = blockSize - 1
		s.active = s.blocks[s.size/blockSize]
	}
	res, s.active[s.offset] = s.active[s.offset], nil
	return
}

// Top returns the value currently on the top of the stack. No bounds are checked.
func (s *Stack) Top() interface{} {
	if s.offset > 0 {
		return s.active[s.offset-1]
	}

	return s.blocks[(s.size-1)/blockSize][blockSize-1]
}

// Empty checks whether the stack is empty or not.
func (s *Stack) Empty() bool {
	return s.size == 0
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	return s.size
}

// Reset the stack, effectively clearing its contents.
func (s *Stack) Reset() {
	*s = *NewStack()
}
