package main

import "fmt"

func main() {
	//Maps agrupam dados por id
	websites := map[string]string{ //Definindo tipo do id e do dado
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)
}
