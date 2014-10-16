// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Huf8jEDUJO

// Create a file with an array of JSON documents that contain a user name and email address. Declare a struct
// type that maps to the JSON document. Using the json package, read the file and create a slice of this struct
// type. Display the slice.
//
// From example 1, Marshal the slice into pretty print strings and display each element.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// user is the data we need to unmarshal and marshal.
type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// main is the entry point for the application.
func main() {
	// Open the file.
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Open File", err)
		return
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	// Decode the file into this slice.
	var users []user
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		fmt.Println("Decode File", err)
		return
	}

	// Iterate over the slice and display
	// each user.
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}

	uData, err := json.MarshalIndent(&users, "", "    ")
	if err != nil {
		fmt.Println("MarshalIndent", err)
		return
	}

	// Convert the byte slice to a string and display.
	fmt.Println(string(uData))
}
