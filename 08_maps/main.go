package main

import "fmt"

func main() {
		// define map
		// emails := make(map[string]string)

		// Assing key values
		// emails["Marcello"] = "marcellovcs@gmail.com"
		// emails["Teste"] = "teste@gmail.com"
		// emails["Mike"] = "mike@gmail.com"

		// declare map and add key values

		emails := map[string]string{
			"Marcello":"marcellovcs@gmail.com", 
			"Teste":"teste@gmail.com", 
			"Mike": "mike@gmail.com",
		}

		fmt.Println(emails)
		fmt.Println(len(emails))
		fmt.Println(emails["Marcello"])
		
		// delete
		delete(emails, "Mike")
		fmt.Println(len(emails))
		fmt.Println(emails)
}
