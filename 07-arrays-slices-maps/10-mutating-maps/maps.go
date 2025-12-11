package main

import "fmt"

func main() {
	//Maps agrupam dados por id
	websites := map[string]string{ //Definindo tipo do id e do dado
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["LinkedIn"] = "https://linkedin.com" //Maps sempre são dinamicos
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
