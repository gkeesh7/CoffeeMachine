package utils

import (
	"CoffeeMachine/structs"
	"fmt"
	"sync"
)

func DispenseBeverageConcurrently(beverageList []string, machine *structs.Machine) {
	ch := make(chan string, len(beverageList))
	for len(beverageList) != 0 {
		parallelCount := Min(machine.Outlets.CountN, len(beverageList))
		beverageListBatch := beverageList[:parallelCount]
		var wg sync.WaitGroup
		wg.Add(parallelCount)
		for _, beverageFromList := range beverageListBatch {
			go func(wg *sync.WaitGroup, beverageFromList string) {
				defer wg.Done()
				ch <- DispenseCommand(machine, beverageFromList)
			}(&wg, beverageFromList)
		}
		wg.Wait()
		beverageList = beverageList[parallelCount:]
	}
	close(ch)
	for element := range ch {
		fmt.Println(element)
	}
}

func DispenseCommand(machine *structs.Machine, command string) string {
	return machine.DispenseBeverage(command)
}
