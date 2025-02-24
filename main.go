package main

import (
	"fmt"
)


type node struct{
	data int
	next *node
}

type linked_list struct{
	head *node
}

func (target_linked *linked_list) print(){
	linked_list_node := target_linked.head

	for linked_list_node != nil {
		fmt.Println(linked_list_node.data)
		linked_list_node = linked_list_node.next
	}
}

func (target_linked *linked_list) add_First(data int){
	if target_linked == nil {
		new_node := &node{data:data ,next:nil}
		target_linked.head = new_node
	}
	new_node := &node{data: data , next:target_linked.head}
	target_linked.head = new_node
}

func (target_linked *linked_list) add_data(data int){
	linked_list_node := target_linked.head
	new_node := &node{data: data, next: nil}
	for linked_list_node.next != nil {
		linked_list_node = linked_list_node.next
	}
	linked_list_node.next = new_node
}

func (target_linked *linked_list) remove_last(){
	linked_list_node := target_linked.head
	for linked_list_node.next.next != nil {
		linked_list_node = linked_list_node.next
	}
	linked_list_node.next = nil
}

func (target_linked *linked_list) remove_first (){
	linked_list_node := target_linked.head
	if linked_list_node!=nil{
		target_linked.head = linked_list_node.next
	}
}

func (target_linked *linked_list) reverse () {
	// reversed_linked_list := &linked_list{head: nil}
	// fmt.Println(reversed_linked_list)
	// fmt.Println(reversed_linked_list.head)
	target_node := target_linked.head.next
	if target_node == nil {
		return
	}
	sample_node := &node{data: target_linked.head.data , next: nil }
	reversed_linked_list := &linked_list{head:sample_node}
	for target_node != nil {
		reversed_linked_list.add_First(target_node.data)
		target_node = target_node.next
	}

	// reversed_linked_list.print()
	target_linked.head = reversed_linked_list.head
}

func main(){
	sample_node := &node{data: 12,next:nil}
	sample_linked_list:= &linked_list{head: sample_node}
	sample_linked_list.add_data(34)
	sample_linked_list.add_data(64)
	sample_linked_list.add_data(5656)
	sample_linked_list.add_First(23)
	// sample_linked_list.remove_last()
	// sample_linked_list.remove_first()
	sample_linked_list.reverse()
	sample_linked_list.print()
}
