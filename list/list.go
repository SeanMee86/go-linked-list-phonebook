package list

import (
	"errors"
)

type LinkedList struct {
	head *listNode
	tail *listNode
}

type listNode struct {
	Data contact
	next *listNode
	prev *listNode
}

type contact struct {
	Name  string
	Phone string
}

func (l *LinkedList) DeleteContact(c *listNode) {
	if c.next != nil {
		c.next.prev = c.prev
	} else {
		l.tail = c.prev
	}
	if c.prev != nil {
		c.prev.next = c.next
	} else {
		l.head = c.next
	}
}

func (l *LinkedList) GetNode(n string) (*listNode, error) {
	curr_node := l.head
	for curr_node != nil && curr_node.Data.Name != n {
		curr_node = curr_node.next
	}
	if curr_node != nil {
		return curr_node, nil
	} else {
		return &listNode{}, errors.New("name not found")
	}
}

func (l *LinkedList) Insert(n string, p string) error {
	var err error
	new_node := &listNode{
		Data: contact{
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
			if new_node.Data.Name > curr_node.Data.Name {
				if curr_node.next != nil {
					curr_node = curr_node.next
				} else {
					curr_node.next = new_node
					l.tail = new_node
					new_node.prev = curr_node
					curr_node = nil
				}
			} else if new_node.Data.Name < curr_node.Data.Name {
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
				err = errors.New("name already exists")
				curr_node = nil
			}
		}
	}
	return err
}

func (l *LinkedList) PrintContacts() []contact {
	curr_node := l.head
	var contactSlice []contact
	for curr_node != nil {
		contactSlice = append(contactSlice, curr_node.Data)
		curr_node = curr_node.next
	}
	return contactSlice
}

func (l *LinkedList) UpdateContact(c *contact, newPhone string) {
	c.Phone = newPhone
}
