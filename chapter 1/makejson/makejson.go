package main

import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
)

// Details  structure for users name and address
type Details struct {
	firstName string
	address string
}
func main(){
	// initialized the struct
	p := new(Details)

	// created a map to add name and address of user
	var userDetails map[string]string = make(map[string]string)
	
	fmt.Println("Please input your name: ")
	scanName := bufio.NewScanner(os.Stdin)
	scanName.Scan()
	p.firstName = scanName.Text()


	fmt.Println("Please input your Address: ")
	scanAddr := bufio.NewScanner(os.Stdin)
	scanAddr.Scan()
	p.address = scanAddr.Text()

	userDetails["firstName"] = p.firstName
	userDetails["address"] = p.address

	jsonEnc, err := json.Marshal(userDetails)
	if err != nil {
		fmt.Println("Encountered an error during encoding")
		return
	}

	fmt.Println("JSON READED: ", string(jsonEnc))
	

}