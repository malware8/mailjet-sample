package main

import (
	"fmt"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"github.com/mailjet/mailjet-apiv3-go/resources"
)

var client *mailjet.Client

func main() {
	client = mailjet.NewMailjetClient(PUBLIC_KEY, PRIVATE_KEY)
	// createList("My List")
	// lists := getLists()
	// if len(lists) > 0 {
	// createContact()
	addProps()
	// }
}

func createList(listName string) {
	var data []resources.Contactslist
	mr := &mailjet.Request{
		Resource: "contactslist",
	}
	fmr := &mailjet.FullRequest{
		Info: mr,
		Payload: &resources.Contactslist{
			Name: listName,
		},
	}
	err := client.Post(fmr, &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data array: %+v\n", data)
}

func getLists() []resources.Contactslist {
	var data []resources.Contactslist
	_, _, err := client.List("contactslist", &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data array: %+v\n", data)

	return data
}

func createContact() {
	var data []resources.Contact
	mr := &mailjet.Request{
		Resource: "contact",
	}
	fmr := &mailjet.FullRequest{
		Info: mr,
		Payload: &resources.Contact{
			Name:  "",
			Email: YOUR_EMAIL,
		},
	}
	err := client.Post(fmr, &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data array: %+v\n", data)
}

func getContacts() []resources.Contact {
	var data []resources.Contact
	_, _, err := client.List("contact", &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data array: %+v\n", data)
	return data
}

func addProps() {
	var data []resources.ContactslistManageManyContacts
	mr := &mailjet.Request{
		Resource: "contactslist",
		ID:       YOUR_LIST_ID,
		Action:   "managemanycontacts",
	}
	fmr := &mailjet.FullRequest{
		Info: mr,
		Payload: &resources.ContactslistManageManyContacts{
			Action: "addnoforce",
			Contacts: []resources.AddContactAction{
				{
					Email:      YOUR_EMAIL,
					Properties: map[string]string{"name": "Jhon Smith"},
				},
			},
		},
	}
	err := client.Post(fmr, &data)
	if err != nil {
		panic(err)
	}
}
