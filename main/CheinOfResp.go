package main

import "fmt"

type client struct {
	name                string
	city                string
	country             string
	registrationDone    bool
	getConsult          bool
	getPrice            bool
	connectWithOperator bool
}
type department interface {
	execute(*client)
	setNext(department)
}
type reception struct {
	next department
}

func (r *reception) execute(p *client) {
	if p.registrationDone {
		fmt.Println("Client registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering client")
	p.registrationDone = true
	r.next.execute(p)
}
func (r *reception) setNext(next department) {
	r.next = next
}

type getConsult struct {
	next department
}

func (d *getConsult) execute(p *client) {
	if p.getConsult {
		fmt.Println("GetConsult checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("1 Online ticket:It includes online-life tour and material about sightseeing in Kazakhstan,4 hours")
	fmt.Println("2 Standard ticket:Standard includes trip to Astana and Almaty, for 3 days and attempting 5 sightseeing in each city with guide, 3 days")
	fmt.Println("3 Advanced ticket:Standard includes trip to Astana and Almaty, for 5 days and attempting those place,which you want to see with guide,5 days ")
	p.getConsult = true
	d.next.execute(p)
}
func (d *getConsult) setNext(next department) {
	d.next = next
}

type price struct {
	next department
}

func (m *price) execute(p *client) {
	if p.getPrice {
		fmt.Println("Information about prices already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Prices:")
	fmt.Println("Online ticket:15$")
	fmt.Println("Standard ticket:150$ ")
	fmt.Println("Advanced ticket:350$ ")
	p.getPrice = true
	m.next.execute(p)
}
func (m *price) setNext(next department) {
	m.next = next
}

type moreInformation struct {
	next department
}

func (c *moreInformation) execute(p *client) {
	if p.connectWithOperator {
		fmt.Println("Get more Information is Done")
	}
	fmt.Println("If you have other questions? Please!contact with us:")
	fmt.Println("87001231212")
	fmt.Println("87022356512")
	fmt.Println("87473586598")
	fmt.Println("Have a nice trip!")
	var choose string
	fmt.Scan(&choose)
	if choose == "Back" || choose == "back" {
		fmt.Println()
		selectMenu("consult")
	}
}
func (c *moreInformation) setNext(next department) {
	c.next = next
}

func getInformation() {
	Client := &moreInformation{}
	GetTicket := &price{}
	GetTicket.setNext(Client)
	GetPrice := &getConsult{}
	GetPrice.setNext(GetTicket)
	reception := &reception{}
	reception.setNext(GetPrice)

	Clients := &client{name: name, city: "Nur-Sultan", country: "Kazakhstan"}
	reception.execute(Clients)

}
