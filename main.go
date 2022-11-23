package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"temp/student"
)

// Create a function that calls a 3rd party http PUT endpoint with unit tests.
//  The endpoint takes in order as query parameter and id as URL path parameter.
// Endpoint: /v1/:id?order=

func main() {

	// functionOne()

	studetns := student.New()

	fmt.Println(studetns)

}

type HelloWorldI interface {
	HelloWorld() (string, error)
}

func doSomething(hello HelloWorldI) {
	msg, err := hello.HelloWorld() //(string, error)
	if err != nil {
		log.Fatalf("Error: %s ", err.Error())

	}

	if msg != "hi" {
		return
	}

	fmt.Println("hello", msg)
}

func functionOne() {
	request, err := http.NewRequest("PUT", "http://localhost:81/v1/:12?order", nil)
	if err != nil {
		log.Fatalf("Error while creating client: %s ", err.Error())
	}

	request.Header.Add("Authorization", "Bearer nsndvsndvirnvsdkjvnkdsvlkdsvknslv")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("Error while getting response: %s ", err.Error())
	}

	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error while getting response: %s ", err.Error())
	}

	fmt.Println(string(msg))
}

// Create a student type with name, age and address as fields.
//  There can be multiple addresses per student with each address having street, city and PIN.
//  Initialise a slice of 3 Students with each student having at least 2 addresses.
