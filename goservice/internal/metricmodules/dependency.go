package metricmodules

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/SerdaOzun/gopham/internal/storage"
	"github.com/SerdaOzun/gopham/internal/utils"
)

func ComputeDependencies(root string) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	parseFiles(files)
}

func parseFiles(files []string) {
	fset := token.NewFileSet()

	for _, file := range files {
		node, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
		runMetricAnalysis(file, node)
		if err != nil {
			// go src file probably is not compiling
			// log.Fatal(err)
		}
	}

	runFinalMetrics()
}

func runMetricAnalysis(file string, node *ast.File) {
	doEFCouplings(file, node)
	doFilePackageRatio(file, node)
	doFunctions(file, node)
}

func runFinalMetrics() {
	doAFCouplings()
}

func doEFCouplings(filename string, node *ast.File) {
	storage.AddProjectPackage(node.Name.Name)
	storage.AddPackageEFItem(node.Name.Name, *node)
	var efCouplings []string
	for _, name := range node.Imports {
		efCouplings = append(efCouplings, name.Path.Value)
		storage.AddUsedImportPackages(name.Path.Value)
	}
	storage.AddEfCouplingsFileItem(filename, efCouplings)
}

//compute how many times packages are being imported by files
func doAFCouplings() {
	projectOnlyPackages := utils.GetAllKeysFromBoolMap(storage.CurrentMB.ProjectPackages)
	for packagename, _ := range storage.CurrentMB.UsedImportPackages {
		var afCouplings []string
		for filename, efcouplings := range storage.CurrentMB.EfCouplingsFile {
			formattedCouplings := utils.ConvertRelativeDependencyPaths(efcouplings)
			if utils.ContainsStrArrayStr(formattedCouplings, packagename) {
				afCouplings = append(afCouplings, filename)
			}
		}
		storage.AddAfCouplingsAllPackage(packagename, afCouplings)
		if utils.ContainsStrArrayStr(projectOnlyPackages, packagename) {
			storage.AddAfCouplingsProjectPackage(packagename, afCouplings)
		}
	}
}

//save the packages of files
func doFilePackageRatio(filename string, node *ast.File) (err error) {
	storage.AddPackageToFile(filename, node.Name.Name)
	return err
}

func doFunctions(filename string, node *ast.File) {
	functionCount := 0
	for _, d := range node.Decls {
		if _, isFn := d.(*ast.FuncDecl); isFn {
			functionCount++
			// fmt.Println(*fn.Body)
		}
	}
	storage.AddFileFunctionCount(filename, functionCount)
	storage.AddPackageFunctionCount(node.Name.Name, functionCount)
}
