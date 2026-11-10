package main

import "fmt"

type contact struct {
	id    int
	name  string
	email string
	phone string
}

var contactList []contact
var contactIndexByName map[string]int
var nextID = 1

func init() {
	contactList = make([]contact, 0)
	contactIndexByName = make(map[string]int)
}

func main() {

	printContacts()
	addContact("Carlos", "carlos@mail.com", "1234567890")
	addContact("Javier", "javier@mail.com", "0987654321")
	printContacts()
	contactFound := findContactByName("javier")
	if contactFound != nil {
		printContact(*contactFound)
	}
	contactFound = findContactByName("Javier")
	if contactFound != nil {
		printContact(*contactFound)
	}

}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact already exists: %v\n", name)

		return
	}

	newContact := contact{
		id:    nextID,
		name:  name,
		email: email,
		phone: phone,
	}
	nextID++

	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1

	fmt.Println()
	fmt.Println("Contact added: ", name)
}

func findContactByName(name string) *contact {
	index, exists := contactIndexByName[name]
	if exists {
		return &contactList[index]
	}

	fmt.Printf("Contact '%s' not found!\n", name)
	return nil
}

func printContacts() {
	fmt.Println()
	fmt.Println("--------Contacts List----------")
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	for i, contact := range contactList {
		fmt.Printf("%d. ID:  %d, Name: %s, Email: %s, Phone: %s\n", i+1,
			contact.id, contact.name, contact.email, contact.phone)
	}

	fmt.Println()
}

func printContact(contact contact) {
	fmt.Printf("ID:  %d, Name: %s, Email: %s, Phone: %s\n", contact.id, contact.name,
		contact.email, contact.phone)
}
