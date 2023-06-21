package main

import "fmt"

// Run - is going to be responsible for
// the initialization and startup of our app
func Run() error {
	fmt.Println("Starting up our app")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
