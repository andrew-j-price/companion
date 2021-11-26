package main

import (
	"fmt"

	jlog "github.com/andrew-j-price/journey/logger"
	"github.com/andrew-j-price/journey/random"
)

func init() {
	jlog.PackageLogger()
}

func main() {
	fmt.Println("Hello, gopher")

	random.RandomGreetingMain()

	// cannot call non-function logger.Info (type *log.Logger)
	// logger.Info("Hi")
}
