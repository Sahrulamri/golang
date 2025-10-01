package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "username")
	var password *string = flag.String("password", "root", "password")
	var host *string = flag.String("host", "127.0.0.1", "host")
	var port *int = flag.Int("port", 3306, "port")

	flag.Parse()

	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
	fmt.Println("host:", *host)
	fmt.Println("port:", *port)
}