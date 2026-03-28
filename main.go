package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Contact struct {
	Name  string
	Phone string
	Email string
	Ended bool
}

func main() {
	a := app.New()
	w := a.NewWindow("contact book")

	contact := []Contact{}
	//	contactBook := ContactBook{}

	//adding contacts info into list
	var list *widget.List
	list = widget.NewList(
		func() int {
			return len(contact)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel("Name"),         // Objects[0]
				widget.NewLabel("Phone"),        // Objects[1]
				widget.NewLabel("Email"),        // Objects[2]
				widget.NewButton("Remove", nil), // Objects[3].
			)

		},
		func(i widget.ListItemID, item fyne.CanvasObject) {
			red := item.(*fyne.Container)
			Name := red.Objects[0].(*widget.Label)
			Phone := red.Objects[1].(*widget.Label)
			Email := red.Objects[2].(*widget.Label)
			btn := red.Objects[3].(*widget.Button)

			Name.SetText(contact[i].Name)
			Phone.SetText(contact[i].Phone)
			Email.SetText(contact[i].Email)

			btn.OnTapped = func() {
				contact = append(contact[:i], contact[i+1:]...)
				list.Refresh()
			}
		},
	)

	//User interface vvvvv

	label := widget.NewLabel("Contact Book")
	label2 := widget.NewLabel(" ")
	labelAdd := widget.NewLabel("Add contact")
	labelSrch := widget.NewLabel("Search contacts")
	labelExist := widget.NewLabel("Contacts")
	labelResults := widget.NewLabel("")
	labelInfo := widget.NewLabel("Searched contact")

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
		contact = append(contact, Contact{Name: entryName.Text, Phone: entryPhone.Text, Email: entryEmail.Text})
		list.Refresh()
		entryName.SetText("")
		entryPhone.SetText("")
		entryEmail.SetText("")

	})
	btnSrch := widget.NewButton("Search", func() {
		name := entrySearch.Text
		found := false
		for _, c := range contact {
			if c.Name == name {
				labelResults.SetText("Name: " + c.Name + ", Phone: " + c.Phone + ", Email: " + c.Email)

				found = true
				break
			}
		}
		if !found {
			labelResults.SetText("Contact not found")
		}
	})

	listContainer := container.NewBorder(nil, nil, nil, nil, list)

	sadrzaj := container.NewVBox(
		label, label2, labelAdd, entryName, entryPhone,
		entryEmail, btnAdd, labelSrch, entrySearch,
		btnSrch, labelExist, listContainer,
		labelInfo, labelResults,
	)
	w.SetContent(sadrzaj)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
