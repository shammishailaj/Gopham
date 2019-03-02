package metricmodules

import (
	"os"
	"testing"

	"github.com/SerdaOzun/gopham/internal/storage"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	root := "../../test_data/"
	name := "test"
	storage.ProjectMetrics = make(map[string]*storage.MetricBank)
	storage.CurrentMB = storage.NewMetricBank()
	go func() { ComputeSourceLoC(root) }()
	ComputeDependencies(root)
	storage.ComputePackageLoC()
	storage.ProjectMetrics[name] = storage.CurrentMB
	// storage.PrintMetricResults(name)

}

func TestLoC(t *testing.T) {
	metricstorageLoC := storage.CurrentMB.FileLoC["../../test_data/metricstorage.go"]
	if metricstorageLoC != 52 {
		t.Errorf("LoC was incorrect. Got: %d, want: %d", metricstorageLoC, 52)
	}
	metricstorageLogicLoC := storage.CurrentMB.FileLoC["../../test_data/metricstoragelogic.go"]
	if metricstorageLogicLoC != 86 {
		t.Errorf("LoC was incorrect. Got: %d, want: %d", metricstorageLoC, 86)
	}
	totalPackageLoC := storage.CurrentMB.PackageLoC["test_data"]
	if totalPackageLoC != 138 {
		t.Errorf("PackageLoC was incorrect. Got: %d, want: %d", totalPackageLoC, 138)
	}
}

func TestPackageNames(t *testing.T) {
	packagename := storage.CurrentMB.ProjectPackages["test_data"]
	if packagename != true {
		t.Errorf("Packagename \"test_data\" not found.")
	}
}

func TestEfCouplings(t *testing.T) {
	tmp1 := len(storage.CurrentMB.EfCouplingsFile["../../test_data/metricstorage.go"])
	if tmp1 != 0 {
		t.Errorf("EF Couplings incorrect. Got: %d, want: %d", tmp1, 0)
	}
	tmp2 := len(storage.CurrentMB.EfCouplingsFile["../../test_data/metricstoragelogic.go"])
	if tmp2 != 2 {
		t.Errorf("EF Couplings incorrect. Got: %d, want: %d", tmp2, 2)
	}
	tmp3 := storage.CurrentMB.PackageEF["test_data"]
	if len(tmp3) != 2 {
		t.Errorf("EF Package Couplings incorrect. Got: %d, want: %d", len(tmp3), 2)
	}
}

func TestAfCouplings(t *testing.T) {
	//manually add a non existing dependency to test functionality
	var tmp []string
	tmp = append(tmp, "test_data")
	storage.AddAfCouplingsAllPackage("test_data", tmp)

	tmp1 := len(storage.CurrentMB.AfCouplingsAllPackage["test_data"])
	if tmp1 != 1 {
		t.Errorf("AF Couplings were incorrect. Got: %d, want: %d", tmp1, 1)
	}
}

func TestPackageOfFile(t *testing.T) {
	tmp1 := storage.CurrentMB.FilePackage["../../test_data/metricstorage.go"]
	if tmp1 != "test_data" {
		t.Errorf("Package name incorrect. Got: %v, want: %v", tmp1, "test_data")
	}
}
