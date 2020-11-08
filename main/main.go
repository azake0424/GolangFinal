package main

import "fmt"

type Client struct {
	name             string
	surname          string
	visa             bool
	registrationDone bool
}

func newClient(name string, surname string, visa bool, done bool) *Client {
	return &Client{
		name:             name,
		surname:          surname,
		visa:             visa,
		registrationDone: done,
	}
}

//=====================================menu=======================================
func menu() {
	if name != "" || surname != "" {
		fmt.Println("Get consult(consult)")
		fmt.Println("Go to tickets(tickets)")
		fmt.Println("Show your check(check)")
		fmt.Println("Exit")
		fmt.Println("=======================================")
	} else {
		fmt.Println("Get consult(consult)")
		fmt.Println("Go to tickets(tickets)")
		fmt.Println("Create account(account)")
		fmt.Println("Exit")
		fmt.Println("=======================================")
	}
	fmt.Println("Select: ")
	fmt.Scan(&selectmenu)
	selectMenu(selectmenu)
}

func selectMenu(choose string) {
	if choose == "account" {
		accountcreate()
	} else if choose == "consult" {
		takeconsult()
	} else if choose == "tickets" {
		typesOfTickets()
	} else if choose == "check" {
		showcheck()
	} else if choose == "exit" || choose == "Exit" {

	} else {
		fmt.Println("You entered incorrectly")
		fmt.Println()
		fmt.Println()
		menu()
	}

}

func typesOfTickets() {
	fmt.Println("You can choose from the list: ")
	fmt.Println("Online")
	fmt.Println("Advanced")
	fmt.Println("Standard")
	fmt.Println("Back")
	fmt.Println("=======================================")
	fmt.Println("Select:")
	fmt.Scan(&typeofticket)
	if typeofticket == "Online" || typeofticket == "online" {
		fmt.Println()
		fmt.Println()
		getOnlineTickets()
	} else if typeofticket == "Advanced" || typeofticket == "advanced" {
		fmt.Println()
		fmt.Println()
		getAdvancedTickets()
	} else if typeofticket == "Standard" || typeofticket == "standard" {
		fmt.Println()
		fmt.Println()
		getStandardTickets()
	} else if typeofticket == "Back" || typeofticket == "back" {
		fmt.Println()
		fmt.Println()
		menu()
	} else {
		fmt.Println("You entered incorrectly")
		fmt.Println()
		fmt.Println()
		selectMenu("tickets")
	}
}

func accountcreate() {
	var visa string
	fmt.Println("Enter Name: ")
	fmt.Scan(&name)
	fmt.Println("Enter Surname: ")
	fmt.Scan(&surname)
	fmt.Println("Do you have a visa?yes/no")
	fmt.Scan(&visa)
	if visa == "yes" || visa == "Yes" {
		fmt.Println("Your visa code")
		fmt.Scan(&code)
		newaccount := newClientFacade(name, surname, true, true, code)
		fmt.Println(newaccount)
		fmt.Println()
		fmt.Println()
		menu()
	} else if visa == "no" || visa == "No" {
		fmt.Println("Sorry but we work only with a visa")
		fmt.Println()
		fmt.Println()
		menu()
	} else {
		fmt.Println("You entered incorrectly")
		fmt.Println()
		fmt.Println()
		selectMenu("account")
	}
}

func takeconsult() {
	var choose string
	fmt.Println("Welcome to our travel agency. What do you want to know?")
	fmt.Println("Information about tickets(information)")
	fmt.Println("Ticket prices(prices)")
	fmt.Println("Contact numbers(contact)")
	fmt.Println("Back")
	fmt.Println("===============================")
	fmt.Println("Select: ")
	fmt.Scan(&choose)
	if choose == "information" || choose == "Information" {

	} else if choose == "prices" || choose == "Prices" {

	} else if choose == "contact" || choose == "Contact" {

	} else if choose == "Back" || choose == "back" {
		menu()
	} else {
		fmt.Println("You entered incorrectly")
		takeconsult()
	}
}

//=====================================end menu=====================================

func main() {
	fmt.Println("Welcome to our travel agency. Marat, Azhar and Zangar worked here.")
	fmt.Println()
	menu()
}
