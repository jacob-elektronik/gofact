package parse

import "strings"

func handleIndicator(s string, indicator string, delimiter string) []string {
	var elementsStr []string
	var relIndex []int
	var indexToDelete []int
	sTmp := strings.Split(s, indicator+delimiter)
	for idx, sub := range sTmp {
		tmp := strings.Split(sub, delimiter)
		elementsStr = append(elementsStr, tmp...)
		if idx != len(sTmp)-1 {
			relIndex = append(relIndex, len(elementsStr)-1)
		}
	}

	var lastIdx int
	for _, i := range relIndex {
		if lastIdx+1 == i {
			elementsStr[lastIdx] = elementsStr[lastIdx] + delimiter + elementsStr[i+1]
		}
		elementsStr[i] = elementsStr[i] + delimiter + elementsStr[i+1]
		indexToDelete = append(indexToDelete, i+1)
		lastIdx = i
	}

	for j, i := range indexToDelete {
		elementsStr = remove(elementsStr, i-j)
	}

	return elementsStr
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}