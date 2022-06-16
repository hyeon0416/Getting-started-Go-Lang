package main

import (
	"fmt"

	"HowToModuleImport/package"
)

func main() {
	fmt.Println("Hello, Modules!")

	mypackage.PrintHello()
}