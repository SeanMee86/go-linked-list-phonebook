package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

type ListNode struct {
	Data Contact
	Next *ListNode
	Prev *ListNode
}

type Contact struct {
	Name  string
	Phone string
}

func main() {
	startProgram()
}

func startProgram() {
	s := bufio.NewScanner(os.Stdin)
	var o string
	var ll LinkedList

	fmt.Println("Welcome to Phonebook")
	fmt.Println("Please select one of the below options")
	fmt.Println()

	for o != "7" {
		if o == "" {
			printOptions()
		}
		s.Scan()
		o = s.Text()
		switch o {
		case "1":
			enterContact(&ll)
		case "2":
			getContact(&ll)
		case "3":
			ll.printContacts()
		case "6":
			printOptions()
		case "7":
			fmt.Println("exiting program")
		default:
			fmt.Println("unknown command")
		}
	}
}

func enterContact(ll *LinkedList) {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Name: ")
	s.Scan()
	name := s.Text()
	fmt.Println()
	fmt.Print("Phone: ")
	s.Scan()
	phone := s.Text()
	ll.insert(name, phone)
}

func (l *LinkedList) insert(n string, p string) {
	new_node := &ListNode{
		Data: Contact{
			Name: n,
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
				if curr_node.Next != nil {
					curr_node = curr_node.Next
				} else {
					curr_node.Next = new_node
					l.tail = new_node
					new_node.Prev = curr_node
					curr_node = nil
				}
			} else if new_node.Data.Name < curr_node.Data.Name {
				if curr_node.Prev == nil {
					curr_node.Prev = new_node
					new_node.Next = curr_node
					l.head = new_node
					curr_node = nil
				} else {
					new_node.Prev = curr_node.Prev
					new_node.Next = curr_node
					curr_node.Prev.Next = new_node
					curr_node.Prev = new_node
					curr_node = nil
				}
			} else {
				fmt.Println("This name already exists")
				curr_node = nil
			}
		}
	}
}

func (l *LinkedList) printContacts() {
	curr_node := l.head

	for curr_node != nil {
		fmt.Println(curr_node.Data.Name, " - ", curr_node.Data.Phone)
		curr_node = curr_node.Next
	}
}

func getContact(l *LinkedList) {

	s := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter name: ")
	s.Scan()
	name := s.Text()

	contact, err := l.getContact(name)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(contact.Name,"-",contact.Phone)
}

func (l *LinkedList) getContact(n string) (Contact, error) {
	curr_node := l.head

	for curr_node.Data.Name != n && curr_node != nil {
		curr_node = curr_node.Next
	}

	if curr_node != nil {
		return curr_node.Data, nil
	} else {
		return Contact{}, errors.New("Name not found")
	}
}

func printOptions() {
	fmt.Println("1: Enter contact information")
	fmt.Println("2: Retrieve contact information")
	fmt.Println("3: Print contacts")
	fmt.Println("4: Update contact")
	fmt.Println("5: Delete contact")
	fmt.Println("6: See options")
	fmt.Println("7: Exit program")
}
