package hello

import "fmt"

func SayHello(name, gender string) string {
	var strGender string
	if gender == "male" {
		strGender = "Bro"
	} else {
		strGender = "Sis"
	}
	return fmt.Sprintf("Hello %s %s!", strGender, name)
}
