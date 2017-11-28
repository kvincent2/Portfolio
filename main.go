// ** This sample program was auto-generated using go-api-skeleton! **

package main

import (
	"os"

	"github.com/kvincent2/portfolio/routes"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.Info("Starting hello-world...")
	workDir, _ := os.Getwd()
	// Start server on port 8080
	routes.GetMainRouter(workDir).Run(":8080")
}
