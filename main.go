package main

import (
	"CoffeeMachine/structs"
	"CoffeeMachine/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	switch len(args) {
	case 1: // Program running without necessary arguments
		fmt.Println("Please supply the json file providing the Machine Json")
		break
	case 2: // Only machine Json is provided as arguments
		machine, err := utils.ParseMachineInputJson(os.Args[1])
		if err != nil {
			break
		}
		consoleModeCommandProcess(machine)
		break
	case 3: // Both machine json and beverage commands are provided
		machine, errJsonFile := utils.ParseMachineInputJson(os.Args[1])
		if errJsonFile != nil {
			break
		}
		beverageList, errCommandFile := utils.ParseCommandInputText(os.Args[2])
		if errCommandFile != nil {
			break
		}
		fileModeCommandProcess(beverageList, machine)
		break
	default: // Undefined or excess arguments provided
		fmt.Println("Undefined number of program arguments provided!!!")
		break
	}
	return
}

func fileModeCommandProcess(beverageList []string, machine *structs.Machine) {
	utils.DispenseBeverageConcurrently(beverageList, machine)
}

func consoleModeCommandProcess(machine *structs.Machine) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp := scanner.Text()
		if inp == "exit" {
			break
		}
		fmt.Println(utils.DispenseCommand(machine, inp))
	}
}
