package main

import (
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"
	"os"
	"strings"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		fmt.Println("\nMenú")
		fmt.Println("1. List contacts")
		fmt.Println("2. Get contact by ID")
		fmt.Println("3. Create contact")
		fmt.Println("4. Update contact")
		fmt.Println("5. Delete contact")
		fmt.Println("6. Exit")
		fmt.Print("Select an option: ")

		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			handlers.ListContacts(db)
		case 2:
			fmt.Println("Enter the contact ID: ")
			var contactID int
			fmt.Scanln(&contactID)
			handlers.GetContactByID(db, contactID)
		case 3:
			contact := inputContactDetails(3)
			handlers.CreateContact(db, contact)
			handlers.ListContacts(db)
		case 4:
			contact := inputContactDetails(4)
			handlers.UpdateContact(db, contact)
			handlers.ListContacts(db)
		case 5:
			fmt.Println("Enter the contact ID: ")
			var contactID int
			fmt.Scanln(&contactID)
			handlers.DeleteContact(db, contactID)
			handlers.ListContacts(db)
		case 6:
			os.Exit(0)
		}

	}
}

func inputContactDetails(option int) models.Contact {
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact
	if option == 4 {
		fmt.Println("Enter the contact ID: ")
		var contactID int
		fmt.Scanln(&contactID)
		contact.ID = contactID
	}
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)
	fmt.Println("Enter email: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)
	fmt.Println("Enter phone: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
