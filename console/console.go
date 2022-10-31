package console

import (
	"bufio"
	"fmt"
	"os"

	"github.com/SeanMee86/phonebook/list"
)

func deleteContact(l *list.LinkedList) {
	name := getName()
	node, err := l.GetNode(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	l.DeleteContact(node)
	fmt.Println(name, "successfully deleted")
}

func enterContact(ll *list.LinkedList) {
	name := getName()
	fmt.Println()
	phone := getPhone()
	err := ll.Insert(name, phone)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name, "entered to contacts")
}

func printContact(l *list.LinkedList) {
	name := getName()
	node, err := l.GetNode(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(node.Data.Name, "-", node.Data.Phone)
}

func getName() string {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter name: ")
	s.Scan()
	return s.Text()
}

func getPhone() string {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Phone: ")
	s.Scan()
	return s.Text()
}

func printContacts(l *list.LinkedList) {
	contacts := l.PrintContacts()
	for _, contact := range contacts {
		fmt.Println(contact.Name, "-", contact.Phone)
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

func updateContact(l *list.LinkedList) {
	name := getName()
	node, err := l.GetNode(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter new phone: ")
	s.Scan()
	l.UpdateContact(&node.Data, s.Text())
	fmt.Println(name, "updated")
}

func StartProgram() {
	s := bufio.NewScanner(os.Stdin)
	var o string
	var ll list.LinkedList

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
			printContact(&ll)
		case "3":
			printContacts(&ll)
		case "4":
			updateContact(&ll)
		case "5":
			deleteContact(&ll)
		case "6":
			printOptions()
		case "7":
			fmt.Println("exiting program")
		default:
			fmt.Println("unknown command")
		}
	}
}
