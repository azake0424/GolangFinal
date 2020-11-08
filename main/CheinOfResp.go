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
	fmt.Println("1 Online ticket:sightseeing with guide,information book")
	fmt.Println("2 Economy ticket:sightseeing with guide,trips to Astana,Almaty each cities for three days")
	fmt.Println("3 Advanced ticket:sightseeing with guide,trip to Astana,Almaty and you will choose one city ,includes:3 times meals,flight's ticket,accommodation each for three days ")
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
	fmt.Println("Online ticket:5000 tenge")
	fmt.Println("Economy ticket:8000 tenge")
	fmt.Println("Advanced ticket:12000 tenge")
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
}
func (c *moreInformation) setNext(next department) {
	c.next = next
}
