package testingo

import "fmt"

func private() { fmt.Println("private") }

func Public() { fmt.Println("Public"); private() }
