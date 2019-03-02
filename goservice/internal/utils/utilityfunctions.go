package utils

import (
	"strings"
)

func ContainsStrArrayStr(s []string, str string) bool {
	for _, a := range s {
		if a == str {
			return true
		}
	}
	return false
}

//Converts string like "./internal/somepackage" to "package"
func ConvertRelativeDependencyPaths(couplings []string) []string {
	var newCouplings []string
	for _, ef := range couplings {
		if strings.HasPrefix(ef, "\".") {
			str := ef[strings.LastIndex(ef, "/")+1:]
			newCouplings = append(newCouplings, "\""+str)
		} else {
			newCouplings = append(newCouplings, ef)
		}
	}
	return newCouplings
}

//return all keys from map in array
func GetAllKeysFromBoolMap(myMap map[string]bool) []string {
	keys := make([]string, len(myMap))
	if myMap != nil {
		i := 0
		for k := range myMap {
			keys[i] = "\"" + k + "\""
			i++
		}
	}
	return keys
}
