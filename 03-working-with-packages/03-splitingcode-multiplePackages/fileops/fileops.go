package fileops

//Não pode ter mais de um pacote na mesma pasta, então cada um tem a sua

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// O nome da função importada tem que começar com maiusculo, diferente de js que tem que escrever export antes
func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("falha ao encontrar o arquivo")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("falha ao gravar valor")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
