package structs

import (
	"fmt"
	"sync"
)

type Machine struct {
	Outlets struct {
		CountN int `json:"count_n"`
	} `json:"outlets"`
	TotalItemsQuantity map[string]int            `json:"total_items_quantity"`
	Beverages          map[string]map[string]int `json:"beverages"`
	mux                sync.Mutex
}

type Input struct {
	Machine Machine `json:"machine"`
}

func (m *Machine) DispenseBeverage(beverageName string) string {
	if m == nil {
		fmt.Println("The machine given is undefined!!!")
		return "Machine Undefined!!!"
	}
	m.mux.Lock()
	defer m.mux.Unlock()
	if beverageMap, ok := m.Beverages[beverageName]; ok {
		for ingredientName, quantity := range beverageMap {
			if totalQty, ok := m.TotalItemsQuantity[ingredientName]; ok {
				if totalQty < quantity {
					return beverageName + " cannot be prepared because item " + ingredientName + " is not sufficient"
				}
			} else {
				return beverageName + " cannot be prepared because " + ingredientName + " is not available"
			}
		}
		for ingredientName, quantity := range beverageMap {
			m.TotalItemsQuantity[ingredientName] -= quantity
		}
		return beverageName + " is prepared"
	}
	fmt.Println("The beverage doesnt exist!!!")
	return "Beverage Undefined!!!"
}
