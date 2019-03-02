package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/SerdaOzun/gopham/internal/metricmodules"
	"github.com/SerdaOzun/gopham/internal/storage"
)

func main() {
	storage.StartDB()
	runEcho()
}

func runAnalysis(root string, name string) {
	root = "/projects/" + root
	storage.ProjectMetrics = make(map[string]*storage.MetricBank)
	storage.CurrentMB = storage.NewMetricBank()
	metricmodules.ComputeSourceLoC(root)
	metricmodules.ComputeDependencies(root)
	storage.ComputePackageLoC()
	storage.AddDate()
	storage.WriteAnalysis(name, *storage.CurrentMB)
	storage.ProjectMetrics[name] = storage.CurrentMB
}

func runEcho() {
	e := echo.New()
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", hello)
	e.GET("/getanalysis/:name/:metric", getMetricRequest)
	e.GET("/projectlist", getProjectList)
	e.POST("/analysis", postAnalysisRoot)
	e.POST("/deleteanalysis", deleteAnalysisFromDB)

	// Start server
	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}

func getMetricRequest(c echo.Context) error {
	name := c.Param("name")
	metric := c.Param("metric")
	fmt.Println("From Project: ", name, " || Get Metric: ", metric)
	return c.String(http.StatusOK, storage.GetMetric(metric, name))
}

func getProjectList(c echo.Context) error {
	projectList := storage.GetProjectList()
	return c.String(http.StatusOK, projectList)
}

func postAnalysisRoot(c echo.Context) error {
	projectroot := new(root)
	if err := c.Bind(projectroot); err != nil {
		return err
	}
	runAnalysis(projectroot.Root, projectroot.Name)
	return c.String(http.StatusOK, "root:"+projectroot.Root+" = Analysis Finished!")
}

func deleteAnalysisFromDB(c echo.Context) error {
	name := new(projectID)
	if err := c.Bind(name); err != nil {
		return err
	}
	storage.DeleteAnalysis(name.Name)
	return c.String(http.StatusOK, name.Name+" succesfully deleted from database")
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type root struct {
	Root string `json:"root"`
	Name string `json:"name"`
}

type projectID struct {
	Name string `json:"name"`
}
