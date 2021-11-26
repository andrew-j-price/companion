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
	fmt.Println("Hello, gopher")

	random.RandomGreetingMain()

	// cannot call non-function logger.Info (type *log.Logger)
	// logger.Info("Hi")

	something()
}

func something() {
	name := faker.Name().FirstName()
	fmt.Printf("Random name: %v\n\n", name)
}
