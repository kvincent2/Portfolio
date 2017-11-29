package routes

import (

	"github.com/kvincent2/portfolio/controllers"

	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

var log = logrus.New()

// Handlers for route management
func GetMainRouter(workdir string) *gin.Engine {
	// Create a default gin router
	router := gin.Default()

	// Load main.gtpl HTML Template
	router.LoadHTMLFiles(fmt.Sprintf("%s/views/main.gtpl", workdir))

	// GET /status
	router.GET("/status", status)

	// GET /hello
	router.GET("/hello", hello)

	// GET /hello_ui
	router.GET("/", hello_ui)

	//static route for styles
	router.StaticFS("/styles/", http.Dir(fmt.Sprintf("%s/views/styles", workdir)))

	// return the gin.Engine
	return router
}

// Status middleware
func status(c *gin.Context) {
	code, body, _ := controllers.GetStatus()
	if code != 200 {
		log.WithFields(logrus.Fields{
			"code": code,
			"body": body,
		}).Warn("/status is not 200")
	}
	c.JSON(code, body)
}

// Hello middleware
func hello(c *gin.Context) {
	code, body, _ := controllers.GetHello()
	if code != 200 {
		log.WithFields(logrus.Fields{
			"code": code,
			"body": body,
		}).Warn("/hello is not 200")
	}
	c.JSON(code, body)
}

// Hello_UI middleware
func hello_ui(c *gin.Context) {
	code, body, _ := controllers.GetHello()
	projects := controllers.GetProjects()
	c.HTML(code, "main.gtpl", map[string]interface{}{
		"Name":     body["name"],
		"Heading":	body["heading"],
		"Projects": projects,
	})
}

