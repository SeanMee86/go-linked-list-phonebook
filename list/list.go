package list

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type LinkedList struct {
	head *listNode
	tail *listNode
}

type listNode struct {
	data contact
	next *listNode
	prev *listNode
}

type contact struct {
	Name  string
	Phone string
}

func (l *LinkedList) DeleteContact(n string) error {
	curr_node := l.head
	for curr_node != nil && curr_node.data.Name != n {
		curr_node = curr_node.next
	}
	if curr_node != nil {
		if curr_node.next != nil {
			curr_node.next.prev = curr_node.prev
		} else {
			l.tail = curr_node.prev
		}
		if curr_node.prev != nil {
			curr_node.prev.next = curr_node.next
		} else {
			l.head = curr_node.next
		}
		return nil
	} else {
		return errors.New("name not found")
	}
}

func (l *LinkedList) GetContact(n string) (contact, error) {
	curr_node := l.head
	for curr_node.data.Name != n && curr_node != nil {
		curr_node = curr_node.next
	}
	if curr_node != nil {
		return curr_node.data, nil
	} else {
		return contact{}, errors.New("name not found")
	}
}

func (l *LinkedList) Insert(n string, p string) {
	new_node := &listNode{
		data: contact{
			Name:  n,
			Phone: p,
		},
	}
	if l.head == nil {
		l.head = new_node
		l.tail = new_node
	} else {
		curr_node := l.head
		for curr_node != nil {
			if new_node.data.Name > curr_node.data.Name {
				if curr_node.next != nil {
					curr_node = curr_node.next
				} else {
					curr_node.next = new_node
					l.tail = new_node
					new_node.prev = curr_node
					curr_node = nil
				}
			} else if new_node.data.Name < curr_node.data.Name {
				if curr_node.prev == nil {
					curr_node.prev = new_node
					new_node.next = curr_node
					l.head = new_node
					curr_node = nil
				} else {
					new_node.prev = curr_node.prev
					new_node.next = curr_node
					curr_node.prev.next = new_node
					curr_node.prev = new_node
					curr_node = nil
				}
			} else {
				fmt.Println("This name already exists")
				curr_node = nil
			}
		}
	}
}

func (l *LinkedList) PrintContacts() {
	curr_node := l.head
	for curr_node != nil {
		fmt.Println(curr_node.data.Name, " - ", curr_node.data.Phone)
		curr_node = curr_node.next
	}
}

func (l *LinkedList) UpdateContact(n string) error {
	curr_node := l.head
	for curr_node != nil && curr_node.data.Name != n {
		curr_node = curr_node.next
	}
	if curr_node == nil {
		return errors.New("name not found")
	}
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter new phone number: ")
	s.Scan()
	curr_node.data.Phone = s.Text()
	fmt.Println(curr_node.data.Name+"'s number has been updated.")
	return nil
}
