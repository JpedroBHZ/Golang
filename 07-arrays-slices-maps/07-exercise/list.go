package main

import "fmt"

func main() {
	//1
	var hobbies = [3]string{"hobbie1", "hobbie2", "hobbie3"}
	fmt.Println(hobbies)

	//2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	//3
	exec3 := hobbies[0:2]
	fmt.Println(exec3)
	exec3 = hobbies[:2]
	fmt.Println(exec3)

	//4
	exec4 := hobbies[1:3]
	fmt.Println(exec4)
	exec4 = hobbies[1:]
	fmt.Println(exec4)

	//5
	goals := []string{"goal1", "goal2"}

	//6
	goals[1] = "goal22"
	goals = append(goals, "goal3")
	fmt.Println(goals)

	//7
	type product struct {
		id    int
		title string
		price float64
	}

	products := []product{
		{
			1,
			"primeiro",
			12.99,
		},
		{
			2,
			"segundo",
			129.99,
		},
	}

	fmt.Println(products)

	newProduct := product{
		3,
		"terceiro",
		15.99,
	}

	products = append(products, newProduct)
	fmt.Println(products)

}
