package main

import "fmt"

func main() {
	fmt.Print("Hello World")
}

//Para compilar precisamos criar um modulo: (go mod init example.com/first-app)
//Depois basta fazer o build: (go build)
//Para executar no terminal sem o go instalado é: (./first-app)
//depois da fazer o build pode ser executado com: go run .
