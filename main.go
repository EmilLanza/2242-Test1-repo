// Filename main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"math/rand"
)



func main(){
	mux := http.NewServeMux() //create multiplexer
	mux.HandleFunc("/", home) //pass key "/" and value home
	mux.HandleFunc("/greeting", greeting) //pass key "/greeting" and value greeting
	mux.HandleFunc("/random", random) //pass key "/random" and value random

	log.Print("starting server on: 4000") //print line
	err := http.ListenAndServe(": 4000", mux) //create server
	log.Fatal(err)

}

//create handler function
func home(w http.ResponseWriter, r *http.Request){ //deals with http methods
	w.Write([]byte("Hello, My name is Emil Lanza and I'm currently doing my internship. I am a student at the University of Belize. I love coding and working with hardware. My favorite computer languages are C++ and Golang. In my free time I love to play recreational sports and video games"))
}

func greeting(w http.ResponseWriter, r *http.Request){
	/*Now := time.Now()
	hr, min := Now.Clock()
	day := Now.Day()
	fmt.Println(day, hr,":",min)
	*/
	Now := time.Now()
    day := Now.Weekday().String()
    hour := Now.Hour()

	var greeting string
    if hour >= 5 && hour < 12 {
        greeting = "Good morning!"
    } else if hour >= 12 && hour < 18 {
        greeting = "Good afternoon!"
    } else {
        greeting = "Good night!"
    }

	html := fmt.Sprintf("<html><body><h1>%s</h1><p>Today is %s.</p><p>The time is %s.</p></body></html>", greeting, day, Now.Format("3:04"))
    fmt.Fprintf(w, html)

}

func random(w http.ResponseWriter, r *http.Request){
	qoute := []string{ //create array
		"To be or not to be",
		"Hello from mars",
		"Not random enough",
		"Just Hello",
	}
	w.Write([]byte(qoute[rand.Intn(len(qoute))]))
}