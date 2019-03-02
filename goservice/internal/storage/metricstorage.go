package storage

import (
	"go/ast"

	"github.com/SerdaOzun/gopham/internal/utils"
)

type MetricBank struct {
	Date                      string
	TotalLoC                  int
	EfCouplingsFile           map[string][]string
	AfCouplingsAllPackage     map[string][]string
	AfCouplingsProjectPackage map[string][]string
	FileLoC                   map[string]int
	PackageLoC                map[string]int
	FileFunctionCount         map[string]int
	PackageFunctionCount      map[string]int
	PackageEFCount            map[string]int
	PackageEF                 map[string][]string
	FilePackage               map[string]string
	ProjectPackages           map[string]bool
	UsedImportPackages        map[string]bool
}

var ProjectMetrics map[string]*MetricBank
var CurrentMB *MetricBank

func NewMetricBank() *MetricBank {
	var mb MetricBank
	mb.EfCouplingsFile = make(map[string][]string)
	mb.AfCouplingsAllPackage = make(map[string][]string)
	mb.AfCouplingsProjectPackage = make(map[string][]string)
	mb.FileLoC = make(map[string]int)
	mb.PackageLoC = make(map[string]int)
	mb.PackageEF = make(map[string][]string)
	mb.FileFunctionCount = make(map[string]int)
	mb.PackageFunctionCount = make(map[string]int)
	mb.PackageEFCount = make(map[string]int)
	mb.FilePackage = make(map[string]string)
	mb.ProjectPackages = make(map[string]bool)
	mb.UsedImportPackages = make(map[string]bool)
	return &mb
}

//Collect Packages defined in the source files
func AddProjectPackage(packagename string) {
	_, found := CurrentMB.ProjectPackages[packagename]
	if !found {
		CurrentMB.ProjectPackages[packagename] = true
	}
}

//collect all the imported Packages in the entire project
func AddUsedImportPackages(packagename string) {
	_, found := CurrentMB.UsedImportPackages[packagename]
	if !found {
		CurrentMB.UsedImportPackages[packagename] = true
	}
}

/*
shows which files use the package
(package names are from a list containing all
imported packagenames in the project, which means also golang packages
and other libraries)
*/
func AddAfCouplingsAllPackage(packagename string, afCouplings []string) {
	CurrentMB.AfCouplingsAllPackage[packagename] = afCouplings
}

/*
shows which files use the package
(Only packages from the project itself are shown. Not thirdparty ones)
*/
func AddAfCouplingsProjectPackage(packagename string, afCouplings []string) {
	CurrentMB.AfCouplingsProjectPackage[packagename] = afCouplings
}

func AddEfCouplingsFileItem(filename string, efCouplings []string) {
	CurrentMB.EfCouplingsFile[filename] = efCouplings
}

func AddFileLoCItem(filename string, loc int) {
	CurrentMB.FileLoC[filename] = loc
}

func AddFileFunctionCount(filename string, functionCount int) {
	CurrentMB.FileFunctionCount[filename] = functionCount
}

func AddPackageFunctionCount(packagename string, functionCount int) {
	CurrentMB.PackageFunctionCount[packagename] = CurrentMB.PackageFunctionCount[packagename] + functionCount
}

func AddPackageLoCItem(filename string, loc int) {
	packageName := CurrentMB.FilePackage[filename]
	currentPackageLoC := CurrentMB.PackageLoC[packageName]
	CurrentMB.PackageLoC[packageName] = currentPackageLoC + loc
}

func AddPackageEFItem(packagename string, couplings ast.File) {
	currentPackageEF := CurrentMB.PackageEF[packagename]
	var newUniqueNodes []string
	for _, cp := range couplings.Imports {
		if !utils.ContainsStrArrayStr(currentPackageEF, cp.Path.Value) {
			newUniqueNodes = append(newUniqueNodes, cp.Path.Value)
		}
	}
	newUniqueNodes = append(newUniqueNodes, currentPackageEF...)
	CurrentMB.PackageEFCount[packagename] = len(newUniqueNodes)
	CurrentMB.PackageEF[packagename] = newUniqueNodes
}

func AddPackageToFile(filename string, packagename string) {
	CurrentMB.FilePackage[filename] = packagename
}

func SetTotalLoC(loc int) {
	CurrentMB.TotalLoC = loc
}
