package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	nspgoipoptim "local.com/nspgo/nspGo-ipOptim"
	nspgosession "local.com/nspgo/nspGo-session"
	nspgotopoviewernew "local.com/nspgo/nspGo-topoViewerV2"
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
	input, err := ioutil.ReadFile(pwd + "/nspGo-topoViewerV2/assets/js/nextData-master.js")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if replaceString != "" {
		output := bytes.Replace(input, []byte(findString), []byte(replaceString), -1)
		if err = ioutil.WriteFile(pwd+"/nspGo-topoViewerV2/assets/js/nextData.js", output, 0666); err != nil {
			log.Error(err)
			os.Exit(1)
		}
	}

}
func main() {
	// graph1 := nsptopoviewer.CyGraph{}

	// content, err := ioutil.ReadFile("../../ietfNetwork.json")
	// if err != nil {
	// 	log.Fatal("Error when opening file: ", err)
	// }
	// filePath, _ := os.Getwd()
	// filePath = (filePath + "../../../vis-library/colajs-asad-graph/data-cytoMarshall.json")
	// graph1.DumpIetfNetworkToCyGraph(content, nsptopoviewer.IetfNetworkStruct{}, filePath)

	s := nspgosession.Session{}

	s.LogLevel = 5
	s.InitLogger()
	s.LoadConfig()

	s.EncodeUserName()
	log.Info(s.EncodeUserName())

	s.GetRestToken()

	// // Get IETF from NSP Topology
	nspgoipoptim := nspgoipoptim.IpOptim{}
	nspgoipoptim.LogLevel = 5
	nspgoipoptim.InitLogger()
	// var dummyPayload []byte
	// ietfRawFile := nspgoipoptim.IpoV4IetfTeNetworksGet(s.IpAdressNspOs, s.Token, s.Proxy.Enable, s.Proxy.ProxyAddress, dummyPayload)

	// // Load  ietfRawFile from file
	filePath, _ := os.Getwd()
	filePath = (filePath + "/nspGo-topoViewerV2/rawTopoFile/ietfNetwork.json")
	ietfRawFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Draw MultiLayer Topology
	nextUiGo := nspgotopoviewernew.NextUiGo{}
	nextUiGo.LogLevel = 5

	nextUiGo.NextUiUnmarshalIetfNetworkModel([]byte(ietfRawFile), nspgotopoviewernew.IetfNetworkStruct{})

	s.RevokeRestToken()

	// Load NetSupFile and Make Physical Topology
	filePathNetSup, _ := os.Getwd()
	filePathNetSup = (filePathNetSup + "/nspGo-topoViewerV2/rawTopoFile/netSupPhysical.json")
	netSupFile, err := ioutil.ReadFile(filePathNetSup)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	nextUiGo.NextUiUnmarshalNetSupL2Model(netSupFile, nspgotopoviewernew.NetSupL2Struct{})

	// connect phy node + igp
	// not implmenting clever logic, whether the node in each layer valid to be inteconnecting.
	// assumming all nodes in each layer MUST be connected
	nextUiGo.NextUiAppendInterLayerLinks(nextUiGo.Topology.Nodes, "L2", "L3")

	// remove duplicated node from Nodelist
	// nextUiGo.RemovedDuplicatedNode(nextUiGo.Topology.Nodes)

	// NextUiMarshal() function is to marshal NextTopology struct into json bytes, then writen into js.
	nextUiGo.NextUiGenerateNextDataJs(nextUiGo.NextUiMarshal())

	// Serve using Gin
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Info(pwd)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob(pwd + "/nspGo-topoViewerV2/templates/*")

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

	router.Static("/assets/js", pwd+"/nspGo-topoViewerV2/assets/js")
	router.Static("/assets/css", pwd+"/nspGo-topoViewerV2/assets/css")
	router.Static("/assets/fonts", pwd+"/nspGo-topoViewerV2/assets/fonts")

	// Start serving the application
	router.Run()

}
