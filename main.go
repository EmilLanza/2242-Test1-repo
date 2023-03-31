// Filename main.go
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	//"strings"
	"time"
)



func main(){
	mux := http.NewServeMux() //creates a multiplexer
	mux.HandleFunc("/", home) //pass key "/" and value home
	mux.HandleFunc("/greeting", greeting) //pass key "/greeting" and value greeting
	mux.HandleFunc("/random", random) //pass key "/random" and value random

	log.Print("starting server on: 4000") 
	err := http.ListenAndServe(": 4000", mux) //create server
	log.Fatal(err)

}

//create handler function
func home(w http.ResponseWriter, r *http.Request){ //deals with http methods
	//html := fmt.Sprintf("<html><head><title>home</title><style>body{background-color:orange; color: black; font-size: 30px;}</style></head><body><p>Hello, My name is Emil Lanza and I'm currently doing my internship. <br/> I am a student at the University of Belize. I love coding and working with hardware. <br/> My favorite computer languages are C++ and Golang. In my free time I love to play recreational sports and video games</p></body></html>")
	html := `<html>
	<head>
	<title>home</title>
	<style>
	body{background-color:orange; color: black; font-size: 30px;}
	</style>
	</head>
	<body>
	<p>Hello, My name is Emil Lanza and I'm currently doing my internship. <br/> I am a student at the University of Belize. I love coding and working with hardware. <br/> My favorite computer languages are C++ and Golang. In my free time I love to play recreational sports and video games
	</p>
	</body>
	</html>`

	fmt.Fprintf(w, html) //writes code to the HTTP response then sends it to the browser
}

func greeting(w http.ResponseWriter, r *http.Request){
	Now := time.Now() //assing current time and date
    day := Now.Weekday().String() //return the day of the week to an interger then converts it to a string
    hour := Now.Hour() //gets current hour 

	var greeting string 
    if hour >= 5 && hour < 12 {
        greeting = "Good morning!"
    } else if hour >= 12 && hour < 18 {
        greeting = "Good afternoon!"
    } else {
        greeting = "Good night!"
    }

	html := fmt.Sprintf("<html><head><title>greeting</title><style>body{text-align:center; background-color:blue; color:yellow; font-size: 30px;}</style></head><body><h1>%s</h1><p>Today is %s.</p><p>The time is %s</p></body></html>", greeting, day, Now.Format("3:04pm"))

	//	html = strings.ReplaceAll(html, "{{time}}", Now.Format("3:04pm"))//updates html by updating string with formatted string
	//html = strings.ReplaceAll(html, "{{greeting}}", greeting)
	//html = strings.ReplaceAll(html, "{{day}}", day)
	fmt.Fprintf(w, html)

}

func random(w http.ResponseWriter, r *http.Request){
	quotes := map[int]string{//creates a map that stores a collection of qoutes with a unique integer
        1: "“Some people feel the rain. Others just get wet.”",
		2: "It is often the small steps, not the giant leaps, that bring about the most lasting change.",
		3: "Attitude is the 'little' thing that makes a big difference.",
		4: "Don't sit down and wait for the opportunities to come. Get up and make them.",
		5: "Never bend your head. Always hold it high. Look the world straight in the eye.",
		6: "The time is always right to do what is right",
		7: "The best thing to hold onto in life is each other",
		8: "In three words I can sum up everything I've learned about life: it goes on.",
		9: "Always forgive your enemies; nothing annoys them so much.",
		10: "You know you're in love when you can't fall asleep because reality is finally better than your dreams.",

    }
	//seed for the random number generator that has the current time
    rand.Seed(time.Now().UnixNano())

	//generates a random number between 1 and the number qoutes in the map
    randomIndex := rand.Intn(len(quotes)) + 1

	//retrieve the qoute with the randomly gemerated index from the map
    randomQuote := quotes[randomIndex]

    html := fmt.Sprintf("<html><head><title>random</title><style>body{background-color: yellow; font-size: 30px; color:blue}</style></head><body><h1>Random Quote:</h1><p>%s</p></body></html>", randomQuote)

    fmt.Fprintf(w, html)
}