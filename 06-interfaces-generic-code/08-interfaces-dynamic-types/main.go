package main

func main() {}

func add(a, b interface{}) interface{} { //Entrada e retorno de qualquer tipo
	aInt, aIsInt := a.(int) //Se a for int, armazena em o valor em aInt e verdadeiro em aIsInt
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt { //Se a e b forem int, soma e retorna ambos
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	if aIsString && bIsString {
		return aString + bString
	}

	return nil
}
