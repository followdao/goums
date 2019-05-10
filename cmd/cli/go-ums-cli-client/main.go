// usage:
// ./go-ums-cli-client register test@gmail.com passwordddd
package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	var cmd, email, password, url string
	url = "http://localhost:3001/register"

	cmd = os.Args[1]
	if cmd == "register" {

		email = os.Args[2]
		password = os.Args[3]

		err = postRegister(url, email, password)
		if err != nil {
			fmt.Println(err.Error())
		}

	}
}
