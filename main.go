package main

import (
	"database/sql"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

type Contact struct {
	Name  string
	Phone string
	Email string
	Ended bool
}

func main() {

	errLabel := widget.NewLabel("")
	infoLabel := widget.NewLabel("")

	// Makes db if needed vvvvv

	db, err := sql.Open("sqlite3", "contacts.db")
	if err != nil {
		errLabel.SetText("Err opening database: " + err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS contacts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			phone TEXT,
			email TEXT
		)
	`)
	if err != nil {
		errLabel.SetText("Err creating table: " + err.Error())
		return
	}
	infoLabel.SetText("Database opened and created")

	a := app.New()
	w := a.NewWindow("contact book")

	//User interface vvvvv

	label := widget.NewLabel("Contact Book")
	label2 := widget.NewLabel(" ")
	labelAdd := widget.NewLabel("Add contact")
	labelSrch := widget.NewLabel("Search contacts")
	labelExist := widget.NewLabel("Contacts")
	labelResults := widget.NewLabel("")
	labelInfo := widget.NewLabel("Searched contact")
	labelContacts := widget.NewLabel("")

	//entry vvvvv
	entryName := widget.NewEntry()
	entryName.SetPlaceHolder("Name...")

	entryPhone := widget.NewEntry()
	entryPhone.SetPlaceHolder("Phone number...")

	entryEmail := widget.NewEntry()
	entryEmail.SetPlaceHolder("Email address...")

	entrySearch := widget.NewEntry()
	entrySearch.SetPlaceHolder("Search contacts by name...")
	//button

	btnAdd := widget.NewButton("Add", func() {
		

		addContact(db, entryName.Text, entryPhone.Text, entryEmail.Text)

		entryName.SetText("")
		entryPhone.SetText("")
		entryEmail.SetText("")

		contactList := displayContacts(db)
		labelContacts.SetText(contactList)

	})
	btnSrch := widget.NewButton("Search", func() {
		rezultat := srchContact(db, entrySearch.Text)
		labelResults.SetText(rezultat)
		entrySearch.SetText("")

		contactList := displayContacts(db)
		labelContacts.SetText(contactList)
	})
	btnDel := widget.NewButton("Remove", func() {
		delContacts(db, entrySearch.Text)
		entrySearch.SetText("")

		contactList := displayContacts(db)
		labelContacts.SetText(contactList)
	})

	//listContainer := container.NewBorder(nil, nil, nil, nil, list)
	//var contactList string
	contactList := displayContacts(db)
	labelContacts.SetText(contactList)

	sadrzaj := container.NewVBox(
		label, label2, labelAdd, entryName, entryPhone,
		entryEmail, btnAdd, labelSrch, entrySearch,
		btnSrch, btnDel, labelExist, labelContacts,
		labelInfo, labelResults,
	)
	w.SetContent(sadrzaj)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func addContact(db *sql.DB, name, phone, email string) {
	_, err := db.Exec("INSERT INTO contacts (name, phone, email) VALUES (?, ?, ?)", name, phone, email)
	if err != nil {
		fmt.Println("Err adding contacts:", err.Error())
		return
	}
}

func srchContact(db *sql.DB, name string) string {
	rows, err := db.Query("SELECT id, name, phone, email FROM contacts WHERE name=? ", name)
	if err != nil {
		return "Err searching contacts: " + err.Error()
	}
	defer rows.Close()

	var results string
	found := false

	for rows.Next() {
		var id int
		var contactName, phone, email string
		err := rows.Scan(&id, &contactName, &phone, &email)
		if err != nil {
			return "Err reading contact: " + err.Error()
		}
		results += "Name: " + contactName + " | Phone: " + phone + " | Email: " + email + "\n"
		found = true
	}

	if !found {
		return "Contact not found"
	}

	return results
}

func displayContacts(db *sql.DB) string {
	rows, err := db.Query("SELECT id, name, phone, email FROM contacts")
	if err != nil {

		return "Err reading contact list: " + err.Error()
	}
	defer rows.Close()

	var results string

	found := false
	for rows.Next() {
		var id int
		var name, phone, email string

		err := rows.Scan(&id, &name, &phone, &email)
		if err != nil {

			return "Err reading contacts8: " + err.Error()
		}
		results += "Name: " + name + " | Phone: " + phone + " | Email: " + email + "\n"
		found = true
	}

	if !found {
		return "Contact not found"
	}
	return results
}
func delContacts(db *sql.DB, name string) {
	_, err := db.Exec("DELETE FROM contacts WHERE name= ?", name)

	if err != nil {
		return

	}
	infoLabel := widget.NewLabel("")

	infoLabel.SetText("Database opened and created")

}
