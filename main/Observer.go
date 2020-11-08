package main

import "fmt"

type observerer interface {
	follow(observer observer)
	unfollow(observer observer)
	extranews()
}

type tickets struct {
	observerlist []observer
	types        string
	available    bool
}

func NewType(types string) *tickets {
	return &tickets{
		types: types,
	}
}

func (c *tickets) checkAvailability() {
	onlineBuilder := getBuilder(typeofticket)
	if typeofticket == "Online" {
		Instructor := newInstructor(onlineBuilder)
		onlineTicket := Instructor.bookticket()

		fmt.Println("Material included:")
		fmt.Println(onlineTicket.material)
		fmt.Println("Duration of the trip:")
		fmt.Println(onlineTicket.duration)
		fmt.Println("Prince in $:")
		fmt.Println(onlineTicket.price)

	} else if typeofticket == "Advanced" {
		Instructor := newInstructor(onlineBuilder)
		advancedTicket := Instructor.bookticket()

		fmt.Println("Material included:")
		fmt.Println(advancedTicket.material)
		fmt.Println("Duration of the trip:")
		fmt.Println(advancedTicket.duration)
		fmt.Println("Prince in $:")
		fmt.Println(advancedTicket.price)
	} else if typeofticket == "Standard" {
		Instructor := newInstructor(onlineBuilder)
		standardTicket := Instructor.bookticket()

		fmt.Println("Material included:")
		fmt.Println(standardTicket.material)
		fmt.Println("Duration of the trip:")
		fmt.Println(standardTicket.duration)
		fmt.Println("Prince in $:")
		fmt.Println(standardTicket.price)
	}
	c.available = true
	c.extranews()
}

func (c *tickets) follow(o observer) {
	c.observerlist = append(c.observerlist, o)
}

func (c *tickets) unfollow(o observer) {
	c.observerlist = remove(c.observerlist, o)
}

func (c *tickets) extranews() {
	for _, observer := range c.observerlist {
		observer.update(c.types)
	}
}

func remove(observerlist []observer, remove observer) []observer {
	len := len(observerlist)
	for i, observer := range observerlist {
		if remove.getid() == observer.getid() {
			observerlist[len-1], observerlist[i] = observerlist[i], observerlist[len-1]
			return observerlist[:len-1]
		}
	}
	return observerlist
}

type observer interface {
	update(string)
	getid() string
}

type clients struct {
	id string
}

func (k *clients) update(types string) {
	fmt.Printf("Dear %s we are glad to say that %s tour is available now", k.id, types)
	var text string
	fmt.Println()
	fmt.Println("Back")
	fmt.Scan(&text)
	if text == "Back" || text == "back" {
		menu()
	}
}

func (k *clients) getid() string {
	return k.id
}

func showcheck() {
	if selectmenu == "" || t == "" {
		fmt.Println("You haven't bought your ticket yet")
		fmt.Println()
		fmt.Println()
		menu()
	} else {
		fmt.Println("CHECK WITH INFORMATION")
		fmt.Println("Dear " + name + " " + surname)
		course := NewType(typeofticket)
		firemail := &clients{id: name}
		course.follow(firemail)
		course.checkAvailability()
	}
}
