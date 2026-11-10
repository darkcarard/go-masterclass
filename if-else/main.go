package main

import "fmt"

func main() {

	userName := "john"
	if hasAcces(userName) {
		fmt.Println("user can access the system")
	} else {
		fmt.Println("user can not access the system!!")
	}
}

func hasAcces(userName string) bool {
	userAccess := map[string]bool{
		"jane": true,
		"john": false,
	}

	return userAccess[userName]
}
