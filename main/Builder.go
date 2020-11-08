package main

import "fmt"

type iBuilder interface {
	setmaterial()
	setPrice()
	setDuration()
	getTicket() ticket
}

func getBuilder(builderType string) iBuilder {
	if builderType == "Online" {

		return &onlineTicket{}
	}
	if builderType == "Standard" {
		return &standardTicket{}
	}
	if builderType == "Advanced" {
		return &advancedTicket{}
	}
	return nil
}

type advancedTicket struct {
	material string
	price    int
	duration string
}

func newadvancedTicket() *advancedTicket {
	return &advancedTicket{}
}

func (a *advancedTicket) setmaterial() {
	a.material = "Standard includes trip to Astana and Almaty, for 5 days and attempting those place,which you want to see with guide"
}

func (a *advancedTicket) setPrice() {
	a.price = 350
}

func (a *advancedTicket) setDuration() {
	a.duration = " 5 days"
}

func (a *advancedTicket) getTicket() ticket {
	return ticket{
		material: a.material,
		price:    a.price,
		duration: a.duration,
	}
}

type standardTicket struct {
	material string
	price    int
	duration string
}

func newStandardTicket() *standardTicket {
	return &standardTicket{}
}

func (s *standardTicket) setmaterial() {
	s.material = "Standard includes trip to Astana and Almaty, for 3 days and attempting 5 sightseeing in each city with guide "
}

func (s *standardTicket) setPrice() {
	s.price = 150
}

func (s *standardTicket) setDuration() {
	s.duration = " 3 days"
}

func (s *standardTicket) getTicket() ticket {
	return ticket{
		material: s.material,
		price:    s.price,
		duration: s.duration,
	}
}

type onlineTicket struct {
	material string
	price    int
	duration string
}

func newOnlineTicket() *onlineTicket {
	return &onlineTicket{}
}

func (o *onlineTicket) setmaterial() {
	o.material = "It includes online-life tour and material about sightseeing in Kazakhstan"
}

func (o *onlineTicket) setPrice() {
	o.price = 15
}

func (o *onlineTicket) setDuration() {
	o.duration = "4 hours"
}

func (o *onlineTicket) getTicket() ticket {
	return ticket{
		material: o.material,
		price:    o.price,
		duration: o.duration,
	}
}

type ticket struct {
	material string
	price    int
	duration string
}

type Instructor struct {
	buiilder iBuilder
}

func newInstructor(o iBuilder) *Instructor {
	return &Instructor{
		buiilder: o,
	}
}

func (i *Instructor) setBuilder(o iBuilder) {
	i.buiilder = o
}

func (i *Instructor) bookticket() ticket {
	i.buiilder.setmaterial()
	i.buiilder.setDuration()
	i.buiilder.setPrice()
	return i.buiilder.getTicket()
}

func getOnlineTickets() {
	fmt.Println("In online ticket we have")
	onlineBuilder := getBuilder("Online")
	Instructor := newInstructor(onlineBuilder)
	onlineTicket := Instructor.bookticket()

	fmt.Printf("Material included: %s\n", onlineTicket.material)
	fmt.Printf("You should pay %d $ for this tour\n", onlineTicket.price)
	fmt.Printf("Duration is %s \n", onlineTicket.duration)
	fmt.Println("Have a wonderful trip!")
	fmt.Println()
	fmt.Println("Buy")
	fmt.Println("Back")
	var text string
	fmt.Scan(&text)
	if text == "Back" || text == "back" {
		typesOfTickets()
	} else if text == "Buy" || text == "buy" {
		fmt.Println()
		fmt.Printf("Price: %d$\n", onlineTicket.price)
		var code int
		fmt.Println("Enter your code vor visa")
		fmt.Scan(&code)
		securitycode := securityCode{
			code: code,
		}
		securitycode.checkCode(code)
	}
}

func getAdvancedTickets() {
	fmt.Println("In advanced ticket we have")
	onlineBuilder := getBuilder("Advanced")
	Instructor := newInstructor(onlineBuilder)
	advancedTicket := Instructor.bookticket()

	fmt.Printf("Material included: %s\n", advancedTicket.material)
	fmt.Printf("You should pay %d $ for this tour\n", advancedTicket.price)
	fmt.Printf("Duration is %s \n", advancedTicket.duration)
	fmt.Println("Have a wonderful trip!")
	fmt.Println()
	fmt.Println("Buy")
	fmt.Println("Back")
	var text string
	fmt.Scan(&text)
	if text == "Back" || text == "back" {
		typesOfTickets()
	} else if text == "Buy" || text == "buy" {
		fmt.Println()
		fmt.Printf("Price: %d$\n", advancedTicket.price)
		var code int
		fmt.Println("Enter your code vor visa")
		fmt.Scan(&code)
		securitycode := securityCode{
			code: code,
		}
		securitycode.checkCode(code)
	}
}

func getStandardTickets() {
	fmt.Println("In standard ticket we have")
	onlineBuilder := getBuilder("Standard")
	Instructor := newInstructor(onlineBuilder)
	standardTicket := Instructor.bookticket()

	fmt.Printf("Material included: %s\n", standardTicket.material)
	fmt.Printf("You should pay %d $ for this tour\n", standardTicket.price)
	fmt.Printf("Duration is %s \n", standardTicket.duration)
	fmt.Println("Have a wonderful trip!")
	fmt.Println()
	fmt.Println("Buy")
	fmt.Println("Back")
	var text string
	fmt.Scan(&text)
	if text == "Back" || text == "back" {
		typesOfTickets()
	} else if text == "Buy" || text == "buy" {
		fmt.Println()
		fmt.Printf("Price: %d$\n", standardTicket.price)
		var code int
		fmt.Println("Enter your code vor visa")
		fmt.Scan(&code)
		securitycode := securityCode{
			code: code,
		}
		securitycode.checkCode(code)
	}
}
