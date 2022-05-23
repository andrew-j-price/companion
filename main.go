package main

import (
	"fmt"

	jlog "github.com/andrew-j-price/journey/logger"
	"github.com/andrew-j-price/journey/random"
	"syreclabs.com/go/faker"
)

func init() {
	jlog.PackageLogger()
}

func main() {
	// Running various functions just to test integrations
	jlog.Info.Println("Companion logging")
	fmt.Println("Hello, gopher")
	random.RandomGreetingMain()
	something()
}

func something() {
	name := faker.Name().FirstName()
	fmt.Printf("Random name: %v\n\n", name)
}
