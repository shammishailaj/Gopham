package test_data

import (
	"encoding/json"
	"fmt"
)

func ComputePackageLoC() {
	for key, value := range fileLoC {
		addPackageLoCItem(key, value)
	}
}

func GetMetric(metric string) string {
	switch metric {
	case "efCouplingsFile":
		content, _ := json.Marshal(efCouplingsFile)
		return string(content)
	case "afCouplingsPackage":
		content, _ := json.Marshal(afCouplingsPackage)
		return string(content)
	case "fileLoC":
		content, _ := json.Marshal(fileLoC)
		return string(content)
	case "packageLoC":
		content, _ := json.Marshal(packageLoC)
		return string(content)
	case "packageEF":
		content, _ := json.Marshal(packageEF)
		return string(content)
	case "filePackage":
		content, _ := json.Marshal(filePackage)
		return string(content)
	case "projectPackages":
		content, _ := json.Marshal(projectPackages)
		return string(content)
	case "usedImportPackages":
		content, _ := json.Marshal(usedImportPackages)
		return string(content)
	default:
		return string("no such metric")
	}
}

// ==================== Printing fucntions
func PrintMetricResults() {
	// printProjectPackages()
	// printPackageLoC()
	// printFileLoC()
	// printPackageEF()
	// printEFCouplings()
	printPackageAFCouplings()
	// printFilePackage()
	fmt.Println("Analysis Finished")
}

func printProjectPackages() {
	fmt.Println("======= Packages")
	for key, _ := range projectPackages {
		fmt.Println("Package: ", key)
	}

}

func printPackageLoC() {
	fmt.Println("======= Package SLOC ")
	for key, value := range packageLoC {
		fmt.Println("Package: ", key, " LoC: ", value)
	}
}

func printFileLoC() {
	fmt.Println("======= File SLOC ")
	for key, value := range fileLoC {
		fmt.Println("File: ", key, " LoC: ", value)
	}
}

func printPackageEF() {
	fmt.Println("======= Package Couplings")
	for key, value := range packageEF {
		fmt.Println("Package: ", key, " EF: ", value)
	}
}

func printEFCouplings() {
	fmt.Println("======= EFCouplings")
	for key, value := range efCouplingsFile {
		fmt.Println("File: ", key, " EF: ", len(value))
	}
}

func printPackageAFCouplings() {
	fmt.Println("======= AFCouplings")
	for key, value := range afCouplingsPackage {
		fmt.Println("Package: ", key, " ", len(value))
	}
}

func printFilePackage() {
	fmt.Println("======= Package of File")
	for key, value := range filePackage {
		fmt.Println("File: ", key, " Package: ", value)
	}
}
