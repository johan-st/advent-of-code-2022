package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_que(t *testing.T) {
	q := que{}
	ps := []pos{{1, 2}, {2, 3}, {3, 4}}
	if !q.empty() {
		fmt.Println("new que was not empty")
		t.Fail()
	}
	q.enque(ps[0])
	if q.empty() {
		fmt.Println("que with element was empty")
		t.Fail()
	}
	if !reflect.DeepEqual(q.deque(), ps[0]) {
		fmt.Println("dequed element was not the same as the qued")
		t.Fail()
	}
	q.enque(ps[0])
	q.enque(ps[1])
	q.enque(ps[2])
	if !reflect.DeepEqual(q.peek(), ps[0]) {
		fmt.Printf("peeked at element was not the same as the qued, first, \nexpected:%v\ngot:%v\n", ps[0], q.peek())
		t.Fail()
	}
	if !reflect.DeepEqual(q.deque(), ps[0]) {
		fmt.Println("dequed element was not the same as the qued, first")
		t.Fail()
	}
	if !reflect.DeepEqual(q.deque(), ps[1]) {
		fmt.Println("dequed element was not the same as the qued, second")
		t.Fail()
	}
	if !reflect.DeepEqual(q.deque(), ps[2]) {
		fmt.Println("dequed element was not the same as the qued, third")
		t.Fail()
	}
}
