package main

import (
	"fmt"
	"os"
	"regexp"
)

// validUserName validates the given username.
//
// Excerpts from man page of useradd:
//
// It is usually recommended to only use usernames that begin with a lower
// case letter or an underscore, followed by lower case letters, digits,
// underscores, or dashes. They can end with a dollar sign. In regular
// expression terms: [a-z_][a-z0-9_-]*[$]?
//
// Usernames may only be up to 32 characters long.
func validUserName(name string) bool {
	return regexp.MustCompile(`^[a-z_]([a-z0-9._-]{0,31}|[a-z0-9._-]{0,30}\$)$`).MatchString(name)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need at least one argument")
		os.Exit(1)
	}
	for _, arg := range os.Args[1:] {
		if !validUserName(arg) {
			fmt.Println(arg + "is not a valid user ID")
			continue
		}
		fmt.Println(arg + " is a valid user ID")
	}
}
