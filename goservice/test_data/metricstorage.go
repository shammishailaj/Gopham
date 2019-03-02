package test_data

var efCouplingsFile map[string][]string
var afCouplingsPackage map[string][]string
var fileLoC map[string]int
var packageLoC map[string]int
var packageEF map[string]int
var filePackage map[string]string
var projectPackages map[string]bool
var usedImportPackages map[string]bool

func InitializeStorage() {
	efCouplingsFile = make(map[string][]string)
	afCouplingsPackage = make(map[string][]string)
	fileLoC = make(map[string]int)
	packageLoC = make(map[string]int)
	packageEF = make(map[string]int)
	filePackage = make(map[string]string)
	projectPackages = make(map[string]bool)
	usedImportPackages = make(map[string]bool)
}

//Collect Packages defined in the source files
func addProjectPackage(packagename string) {
	_, found := projectPackages[packagename]
	if !found {
		projectPackages[packagename] = true
	}
}

//collect all the imported Packages in the entire project
func addUsedImportPackages(packagename string) {
	_, found := usedImportPackages[packagename]
	if !found {
		usedImportPackages[packagename] = true
	}
}

/*
shows which files use the packagename
(package names are from a list containing all
imported packagenames in the project, which means also golang packages
and other libraries)
*/
func addAfCouplingsPackage(packagename string, afCouplings []string) {
	afCouplingsPackage[packagename] = afCouplings
}

func addEfCouplingsFileItem(filename string, efCouplings []string) {
	efCouplingsFile[filename] = efCouplings
}

func addFileLoCItem(filename string, loc int) {
	fileLoC[filename] = loc
}

func addPackageLoCItem(filename string, loc int) {
	packageName := filePackage[filename]
	currentPackageLoC := packageLoC[packageName]
	packageLoC[packageName] = currentPackageLoC + loc
}

func addPackageEFItem(packagename string, loc int) {
	currentPackageEF := packageEF[packagename]
	packageEF[packagename] = loc + currentPackageEF
}

func addPackageToFile(filename string, packagename string) {
	filePackage[filename] = packagename
}
