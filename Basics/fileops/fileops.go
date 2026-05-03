package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(fileName string) (fileData float64, err error) {
	fileDataToBytes, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("Failed to open the file!")
	}
	fileDataToString := string(fileDataToBytes)
	fileData, err = strconv.ParseFloat(fileDataToString, 64)
	if err != nil {
		return 0, errors.New("Failed to parse the value!")
	}
	return fileData, nil
}
func WriteFloatToFile(fileName string, data float64) error {
	dataToString := fmt.Sprint(data)
	err := os.WriteFile(fileName, []byte(dataToString), 0644)
	if err != nil {
		errorString := fmt.Sprintf("Some error occurred while writing to the file! %v", err)
		return errors.New(errorString)
	}
	return nil
}
