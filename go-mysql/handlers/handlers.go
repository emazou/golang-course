package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

func ListContacts(db *sql.DB) {
	query := "SELECT * FROM contact"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("ID | Name | Email | Phone")
	for rows.Next() {
		contact := models.Contact{}
		var valueEmail sql.NullString
		err := rows.Scan(&contact.ID, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}
		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "No email provided"
		}
		fmt.Printf("%v | %v | %v | %v\n", contact.ID, contact.Name, contact.Email, contact.Phone)
	}
}

func GetContactByID(db *sql.DB, contactID int) models.Contact {
	query := "SELECT * FROM contact WHERE id = ?"
	row := db.QueryRow(query, contactID)

	contact := models.Contact{}
	var valueEmail sql.NullString
	err := row.Scan(&contact.ID, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("Contact with ID %v not found", contactID)
		}
	}
	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "No email provided"
	}
	fmt.Println("ID | Name | Email | Phone")
	fmt.Printf("%v | %v | %v | %v\n", contact.ID, contact.Name, contact.Email, contact.Phone)
	return contact
}

func CreateContact(db *sql.DB, contact models.Contact) {
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"
	result, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Contact with ID %v created successfully\n", id)
}

func UpdateContact(db *sql.DB, contact models.Contact) {
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"
	result, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.ID)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Contact with ID %v updated successfully\n", contact.ID)
	fmt.Printf("Rows affected: %v\n", rowsAffected)
}

func DeleteContact(db *sql.DB, contactID int) {
	query := "DELETE FROM contact WHERE id = ?"
	result, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Contact with ID %v deleted successfully\n", contactID)
	fmt.Printf("Rows affected: %v\n", rowsAffected)
}
