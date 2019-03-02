package storage

import (
	"encoding/json"
	"fmt"
	"time"
)

func ComputePackageLoC() {
	for key, value := range CurrentMB.FileLoC {
		AddPackageLoCItem(key, value)
	}
}

func AddDate() {
	current_time := time.Now().Local()
	CurrentMB.Date = current_time.Format("2006-01-02")
}

func GetMetric(metric string, name string) string {
	mb := GetAnalysis(name)

	switch metric {
	case "totalLoC":
		content, _ := json.Marshal(mb.TotalLoC)
		return string(content)
	case "efCouplingsFile":
		content, _ := json.Marshal(transformMap(mb.EfCouplingsFile))
		return string(content)
	case "afCouplingsAllPackage":
		content, _ := json.Marshal(transformMap(mb.AfCouplingsAllPackage))
		return string(content)
	case "afCouplingsProjectPackage":
		content, _ := json.Marshal(transformMap(mb.AfCouplingsProjectPackage))
		return string(content)
	case "fileLoC":
		content, _ := json.Marshal(mb.FileLoC)
		return string(content)
	case "packageLoC":
		content, _ := json.Marshal(mb.PackageLoC)
		return string(content)
	case "fileFunctionCount":
		content, _ := json.Marshal(mb.FileFunctionCount)
		return string(content)
	case "packageFunctionCount":
		content, _ := json.Marshal(mb.PackageFunctionCount)
		return string(content)
	case "packageEFCount":
		content, _ := json.Marshal(mb.PackageEFCount)
		return string(content)
	case "packageEF":
		content, _ := json.Marshal(mb.PackageEF)
		return string(content)
	case "filePackage":
		content, _ := json.Marshal(mb.FilePackage)
		return string(content)
	case "projectPackages":
		content, _ := json.Marshal(mb.ProjectPackages)
		return string(content)
	case "usedImportPackages":
		content, _ := json.Marshal(mb.UsedImportPackages)
		return string(content)
	default:
		return string("no such metric")
	}
}

// converts array value to its length
func transformMap(tmp map[string][]string) map[string]int {
	var newMap = make(map[string]int)
	for key, value := range tmp {
		newMap[key] = len(value)
	}
	return newMap
}

// ==================== Printing fucntions
func PrintMetricResults(projectName string) {
	var mb = ProjectMetrics[projectName]
	printProjectPackages(mb)
	printAllImportedPackages(mb)
	// printPackageLoC(mb)
	// printFileLoC(mb)
	printPackageEF(mb)
	printEFCouplings(mb)
	printPackageAFCouplings(mb)
	// printFilePackage(mb)
	fmt.Println("Analysis Finished")
}

func printProjectPackages(mb *MetricBank) {
	fmt.Println("======= Packages")
	for key, _ := range mb.ProjectPackages {
		fmt.Println("Package: ", key)
	}
}

func printAllImportedPackages(mb *MetricBank) {
	fmt.Println("======= All Imported Packages")
	for key, _ := range mb.UsedImportPackages {
		fmt.Println("Package: ", key)
	}
}

func printPackageLoC(mb *MetricBank) {
	fmt.Println("======= Package SLOC ")
	for key, value := range mb.PackageLoC {
		fmt.Println("Package: ", key, " LoC: ", value)
	}
}

func printFileLoC(mb *MetricBank) {
	fmt.Println("======= File SLOC ")
	for key, value := range mb.FileLoC {
		fmt.Println("File: ", key, " LoC: ", value)
	}
}

func printPackageEF(mb *MetricBank) {
	fmt.Println("======= Package Couplings")
	for key, value := range mb.PackageEF {
		fmt.Println("Package: ", key, " EF: ", value)
	}
}

func printEFCouplings(mb *MetricBank) {
	fmt.Println("======= EFCouplings")
	for key, value := range mb.EfCouplingsFile {
		fmt.Println("File: ", key, " EF: ", len(value))
		for _, coupling := range value {
			fmt.Println("               ", coupling)
		}
	}
}

func printPackageAFCouplings(mb *MetricBank) {
	fmt.Println("======= AFCouplings All Packages")
	for key, value := range mb.AfCouplingsAllPackage {
		fmt.Println("Package: ", key, " ", len(value))
		for _, coupling := range value {
			fmt.Println("               ", coupling)
		}
	}
	fmt.Println("======= AFCouplings Project Packages")
	for key, value := range mb.AfCouplingsProjectPackage {
		fmt.Println("Package: ", key, " ", len(value))
		for _, coupling := range value {
			fmt.Println("               ", coupling)
		}
	}
}

func printFilePackage(mb *MetricBank) {
	fmt.Println("======= Package of File")
	for key, value := range mb.FilePackage {
		fmt.Println("File: ", key, " Package: ", value)
	}
}
