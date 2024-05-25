package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	// serialize the data
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts) // Write to file
	if err != nil {
		return err
	}
	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// deserialize the data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts) // Read from file
	if err != nil {
		return err
	}
	return nil

}

func main() {
	var contacts []Contact

	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error loading contacts from file:", err)
	}

	// create instance of fubio
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("====Contacts Manager====\n",
			"1. Add Contact\n",
			"2. List Contacts\n",
			"3. Save and Exit\n",
			"Enter your choice: ")
		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error reading option:", err)
			return
		}
		switch option {
		case 1:
			var c Contact
			fmt.Print("Enter name: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Enter email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Enter phone: ")
			c.Phone, _ = reader.ReadString('\n')
			contacts = append(contacts, c)

			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error saving contacts to file:", err)
			}
		case 2:
			fmt.Println("====Contacts====")
			for i, c := range contacts {
				fmt.Printf("%d. %s %s %s\n", i+1, c.Name, c.Email, c.Phone)
			}
		case 3:
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error saving contacts to file:", err)
			}
		default:
			fmt.Println("Invalid option")
		}
	}
}
