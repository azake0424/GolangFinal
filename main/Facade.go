package main

import "fmt"

type Facade struct {
	client       *Client
	securityCode *securityCode
}

func newClientFacade(name string, surname string, visa bool, done bool, code int) *Facade {
	fmt.Println("Starting create account")
	wallet := &Facade{
		client:       newClient(name, surname, visa, done),
		securityCode: newSecurityCode(code),
	}
	fmt.Println("Account created")
	fmt.Println("=======================================")
	menu()
	return wallet
}

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

//Global variable
var code int
var name string
var surname string
var typeofticket string
var selectmenu string
var t string

func (s *securityCode) checkCode(incomingCode int) {
	if code == incomingCode {
		t = typeofticket
		fmt.Println("SecurityCode Verified")
		fmt.Println("Thank you for your purchase")
		fmt.Println()
		menu()
	} else {
		fmt.Println("Incorrect code")
		fmt.Println()
		typesOfTickets()
	}
}
