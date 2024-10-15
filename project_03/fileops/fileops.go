package fileops

import (
	"fmt"
	"os"
	"errors"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)

	os.WriteFile(fileName, []byte(valueText), 0644)
}

func ReadFloatFromFile(fileName string) (float64, error) {
	fileData, err := os.ReadFile(fileName)

	if err != nil {
		return 0.0, errors.New("Failed to read file!")
	}

	valueText := string(fileData)

	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0.0, errors.New("Failed to parse stored value!")
	}

	return value, nil

}

