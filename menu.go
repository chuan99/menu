package main

import "fmt"

func main() {
	var cmd string
	fmt.Println("Hello, welcome to menu!")
	for {
		fmt.Scanln(&cmd)
		switch cmd {
		case "version":
			fmt.Println("menu v1.0")
		case "hello":
			fmt.Println("hello~")
		case "help":
			fmt.Println("help")
		case "quit":
			fmt.Println("bye~")
			return
		default:
			fmt.Println("Wrong command!")
		}
	}

}