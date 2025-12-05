package main

import (
	
	"fmt"
	"sort"
)

var capacity int
var freeSlots []int
var regToSlot map[string]int
var slotToReg map[int]string
var isInitialized bool  

func NewParkingLot(size int) {
	
	capacity = size
	freeSlots = []int{}
	for i := 1; i <= size; i++ {
		freeSlots = append(freeSlots, i)
	}
	regToSlot = make(map[string]int)
	slotToReg = make(map[int]string)
	isInitialized = true
}

func Park(reg string) (string, error) {
	
	if !isInitialized {
    return "Please create parking lot first", nil
	}
	if len(regToSlot) == capacity {
		return "Sorry, parking lot is full", nil
	}

	sort.Ints(freeSlots)
	slot := freeSlots[0]
	freeSlots = freeSlots[1:]

	regToSlot[reg] = slot
	slotToReg[slot] = reg

	return fmt.Sprintf("Allocated slot number: %d", slot), nil
}

func Leave(reg string, hours int) (string, error) {

	if !isInitialized {
    return "Please create parking lot first", nil
	}

	slot, exists := regToSlot[reg]
	if !exists {
		return fmt.Sprintf("Registration number %s not found", reg), nil
	}

	delete(regToSlot, reg)
	delete(slotToReg, slot)
	freeSlots = append(freeSlots, slot)

	charge := CalcCharge(hours)
	return fmt.Sprintf("Registration number %s with Slot Number %d is free with Charge $%d",
		reg, slot, charge), nil
}

func Status() string {
	if !isInitialized {
    return "Please create parking lot first\n"
	}
	result := "Slot No. Registration No.\n"
	for i := 1; i <= capacity; i++ {
		if reg, ok := slotToReg[i]; ok {
			result += fmt.Sprintf("%d %s\n", i, reg)
		}
	}
	return result
}

func CalcCharge(hours int) int {
	if hours <= 2 {
		return 10
	}
	return 10 + (hours-2)*10
}
