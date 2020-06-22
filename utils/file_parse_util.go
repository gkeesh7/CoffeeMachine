package utils

import (
	"CoffeeMachine/structs"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ParseMachineInputJson(fileName string) (*structs.Machine, error) {

	var fileInputStruct structs.Input
	jsonFile, jFileErr := os.Open(fileName)

	if jFileErr != nil {
		fmt.Println("File Error!!!")
		fmt.Println(jFileErr)
		return nil, jFileErr
	}
	defer jsonFile.Close()

	byteValue, readErr := ioutil.ReadAll(jsonFile)

	if readErr != nil {
		fmt.Println("Reader Error!!!")
		fmt.Println(readErr)
		return nil, readErr
	}

	unMarshalErr := json.Unmarshal(byteValue, &fileInputStruct)

	if unMarshalErr != nil {
		fmt.Println("Unmarshal Error!!!")
		fmt.Println(unMarshalErr)
		return nil, unMarshalErr
	}

	return &fileInputStruct.Machine, nil
}

func ParseCommandInputText(fileName string) ([]string, error) {

	txtFile, txtFileErr := os.Open(fileName)

	if txtFileErr != nil {
		fmt.Println("File Error!!!")
		fmt.Println(txtFileErr)
		return nil, txtFileErr
	}

	defer txtFile.Close()

	scanner := bufio.NewScanner(txtFile)
	scanner.Split(bufio.ScanLines)
	var cmdLines []string

	for scanner.Scan() {
		cmdLines = append(cmdLines, scanner.Text())
	}

	return cmdLines, nil
}
