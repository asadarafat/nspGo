// main.go

package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// Bind HTML checkboxes
// https://chenyitian.gitbooks.io/gin-web-framework/content/docs/21.html
// https://github.com/gin-gonic/gin/issues/129#issuecomment-124260092
type myForm struct {
	Colors []string `form:"colors[]"`
}

func layerHandler(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm)
	//c.JSON(200, gin.H{"color": fakeForm.Colors})
	replaceText("L2-", fakeForm.Colors[0]+"-")
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": fakeForm.Colors,
		},
	)
	log.Infoln(fakeForm.Colors)
}

func transportTunnelHandler(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm)
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": fakeForm.Colors,
		},
	)
}

func replaceText(findString string, replaceString string) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("present working directory: " + pwd)
	input, err := ioutil.ReadFile(pwd + "/nsoGo-restServer/assets/js/nextData-master.js")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if replaceString != "" {
		output := bytes.Replace(input, []byte(findString), []byte(replaceString), -1)
		if err = ioutil.WriteFile(pwd+"/nsoGo-restServer/assets/js/nextData.js", output, 0666); err != nil {
			log.Error(err)
			os.Exit(1)
		}
	}

}

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Info(pwd)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob(pwd + "/nsoGo-restServer/templates/*")

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.
	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
		log.Info("hello gin terminal")
	})

	router.POST("/layer", layerHandler)

	router.POST("/transport-tunnel", transportTunnelHandler)

	router.Static("/assets/js", pwd+"/nsoGo-restServer/assets/js")
	router.Static("/assets/css", pwd+"/nsoGo-restServer/assets/css")
	router.Static("/assets/fonts", pwd+"/nsoGo-restServer/assets/fonts")

	// Start serving the application
	router.Run()

}
