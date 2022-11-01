package console

import (
	"bufio"
	"fmt"
	"os"

	"github.com/SeanMee86/phonebook/list"
)

var colors = map[string]string{
	"blue": "\033[34m",
	"green": "\033[32m",
	"red": "\033[31m",
	"reset": "\033[0m",
}

func deleteContact(l *list.LinkedList, s *bufio.Scanner) {
	name := getName(s)
	node, err := l.GetNode(name)
	if err != nil {
		printError(err)
		return
	}
	l.DeleteContact(node)
	coloredName := colors["red"]+name+colors["reset"]
	fmt.Println("\nContact", coloredName, "successfully deleted")
	optionsMessage()
}

func enterContact(ll *list.LinkedList, s *bufio.Scanner) {
	name := getName(s)
	phone := getPhone(s)
	err := ll.InsertContact(name, phone)
	if err != nil {
		printError(err)
		return
	}
	coloredName := colors["green"]+name+colors["reset"]
	fmt.Println("\n" + coloredName, "has been added to contacts!")
	optionsMessage()
}

func getName(s *bufio.Scanner) string {
	fmt.Print("\nEnter name: ")
	s.Scan()
	return s.Text()
}

func getPhone(s *bufio.Scanner) string {
	fmt.Print("Enter phone number: ")
	s.Scan()
	return s.Text()
}

func optionsMessage() {
	fmt.Println("\n6 to see options.")
	fmt.Println()
}

func printContact(l *list.LinkedList, s *bufio.Scanner) {
	name := getName(s)
	node, err := l.GetNode(name)
	if err != nil {
		printError(err)
		return
	}
	fmt.Println("\n"+node.Data.Name, "-", node.Data.Phone)
	optionsMessage()
}

func printContacts(l *list.LinkedList) {
	fmt.Println()
	contacts := l.PrintContacts()
	for _, contact := range contacts {
		fmt.Println(contact.Name, "-", contact.Phone)
	}
	optionsMessage()
}

func printError(msg error) {
	fmt.Println()
	fmt.Println(colors["red"]+msg.Error()+colors["reset"])
	optionsMessage()
}

func printOptions() {
	fmt.Println()
	fmt.Println("1: Enter contact information")
	fmt.Println("2: Retrieve contact information")
	fmt.Println("3: Print contacts")
	fmt.Println("4: Update contact")
	fmt.Println("5: Delete contact")
	fmt.Println("6: See options")
	fmt.Println("7: Exit program")
	fmt.Println()
}

func updateContact(l *list.LinkedList, s *bufio.Scanner) {
	name := getName(s)
	node, err := l.GetNode(name)
	if err != nil {
		printError(err)
		return
	}
	fmt.Print("Enter new phone: ")
	s.Scan()
	l.UpdateContact(&node.Data, s.Text())
	coloredName := colors["blue"]+name+colors["reset"]
	fmt.Println("\nContact", coloredName, "has been updated.")
	optionsMessage()
}

func StartProgram() {
	s := bufio.NewScanner(os.Stdin)
	var o string
	var ll list.LinkedList
	fmt.Println()
	fmt.Println()
	fmt.Println("*************** Welcome to Phonebook ***************")
	fmt.Println()
	fmt.Println("****** Please select one of the below options ******")
	printOptions()

	for o != "7" {
		s.Scan()
		o = s.Text()
		switch o {
		case "1":
			enterContact(&ll, s)
		case "2":
			printContact(&ll, s)
		case "3":
			printContacts(&ll)
		case "4":
			updateContact(&ll, s)
		case "5":
			deleteContact(&ll, s)
		case "6":
			printOptions()
		case "7":
			fmt.Println("\nExiting program...")
		default:
			fmt.Println()
			fmt.Println("Unknown Command")
			optionsMessage()
		}
	}
}
