package utils

import (
	"math/big"
)

func OneOfStrings(mainString string, stringList ...string) bool {
	for _, str := range stringList {
		if mainString == str {
			return true
		}
	}
	return false
}

// Must be a nice formated string as numbers
func CompareStringNumber(a, b string) int {
	bigA, _ := new(big.Int).SetString(a, 10)
	bigB, _ := new(big.Int).SetString(b, 10)
	return bigA.Cmp(bigB)
}
