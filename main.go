// Filename main.go
package main

import (
	"log"
	"net/http"
)



func main(){
	mux := http.NewServeMux() //create multiplexer
	mux.HandleFunc("/", home) //pass key "/" and value home
	mux.HandleFunc("/greeting", greeting) //pass key "/greeting" and value greeting
	//mux.HandleFunc("/random", random) //pass key "/random" and value random

	log.Print("starting server on: 4000") //print line
	err := http.ListenAndServe(": 4000", mux) //create server
	log.Fatal(err)

}

//create handler function
func home(w http.ResponseWriter, r *http.Request){ //deals with http methods
	w.Write([]byte("Hello, My name is Emil Lanza and I'm currently doing my internship. I am a student at the University of Belize. I love coding and working with hardware. My favorite computer languages are C++ and Golang. In my free time I love to play recreational sports and video games"))
}

func greeting(w http.ResponseWriter, r *http.Request){
	Now := time.Now()
}

