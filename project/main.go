package main

import (
	"log"
	"os"
	"strings"
	"time"
)

type User struct {
	Email string
	Name  string
}

var Database = []User{
	{Email: "surajv.elec@gmail.com", Name: "Suraj Verma"},
	{Email: "alexander.davis@gmail.com", Name: "Alexander Davis"},
	{Email: "alexander.jackson@gmail.com", Name: "Alexander Jackson"},
	{Email: "avery.williamson@gmail.com", Name: "Avery Williamson"},
	{Email: "tom.hardy@gmail.com", Name: "Tom Hardy"},
	{Email: "ema.jackson@gmail.com", Name: "Ema Jackson"},
	{Email: "paul.watson@gmail.com", Name: "Paul Watson"},
	{Email: "tom.cruse@gmail.com", Name: "Tom Cruse"},
	{Email: "william.shakespear@gmail.com", Name: "William Shakespear"},
	{Email: "albert.einstein@gmail.com", Name: "Albert Einstein"},
	{Email: "albert.paulson@gmail.com", Name: "Albert Paulson"},
	{Email: "issac.newton@gmail.com", Name: "Issac Newton"},
	{Email: "grahambell.fox@gmail.com", Name: "Grahambell Fox"},
	{Email: "hello.world@gmail.com", Name: "Hello World"},
	{Email: "madam.curie@gmail.com", Name: "Madam Curie"},
	{Email: "kevin.peterson@gmail.com", Name: "Kevin Peterson"},
	{Email: "stuard.broad@gmail.com", Name: "Stuard Broad"},
	{Email: "ema.watson@gmail.com", Name: "Ema Watson"},
}

type Worker struct {
	users []User
	name  string
	ch    chan *User
}

func NewWorker(usrs []User, chnl chan *User, name string) *Worker {
	return &Worker{users: usrs, ch: chnl, name: name}
}

func (w *Worker) Find(email string) {
	for i := range w.users {
		user := &w.users[i]
		if strings.Contains(user.Email, email) {
			log.Printf("Channel Name: %s\n", w.name)
			w.ch <- user
		}
	}
}

func main() {
	email := os.Args[1]
	ch := make(chan *User)
	log.Printf("Looking for %s email\n", email)
	go NewWorker(Database[:6], ch, "#1").Find(email)
	go NewWorker(Database[6:12], ch, "#2").Find(email)
	go NewWorker(Database[12:], ch, "#3").Find(email)

	for {
		select {
		case user := <-ch:
			log.Printf("Found User: %s having Email: %s\n", user.Email, user.Name)
		case <-time.After(1 * time.Second):
			return
		}
	}

}
