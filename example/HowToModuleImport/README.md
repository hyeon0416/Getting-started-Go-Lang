# How To Make Module and How To Import Module?

go mod init HowToModuleImport

mkdir package

create main.go file

---HowToModuleImport
        ---package
            ---go.mod
            ---main.go


open your main.go file and push this one

```go
package main

import (
	"fmt"

	"HowToModuleImport/package"
)

func main() {
	fmt.Println("Hello, Modules!")

	mypackage.PrintHello()
}
```
cd package

create mypackage.go

open your mypackage.go and push this code

```go
package mypackage

import "fmt"

func PrintHello() {
	fmt.Println("Hello, Modules! This is mypackage speaking!")
}
```

cd ..

go run main.go

