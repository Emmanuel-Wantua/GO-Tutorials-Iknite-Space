package main

import (
	"fmt"
	"regexp"
	"strings"
)

func StringSanitizer(s string) (string, error) {
	str := strings.TrimSpace(s)
	str = strings.ToLower(str)

	if strings.Contains(str, " ") {
		return "", fmt.Errorf("Your email cannot contain spaces")
	}

	r := regexp.MustCompile("^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,}$")
	if !r.MatchString(str) {
		return "", fmt.Errorf("Please enter a valid Email")
	}

	return str, nil
}

func main() {
	email := "EMMANUEL WANTUA@GMAIL .COM"
	email2 := " EMMANUELWANTUA@GMAIL.COM "
	email3 := "EMMANUELWANTUAGMAIL.COM"
	s, err := StringSanitizer(email)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Your correct email is: %s\n", s)
	}

	s2, err := StringSanitizer(email2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Your correct email is: %s\n", s2)
	}

	s3, err := StringSanitizer(email3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Your correct email is: %s\n", s3)
	}
}
