package nsptopoviewer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	nspgotools "local.com/nspgo/nspGo-tools"
)

type NextUiGo struct {
	Topology NextTopology
	LogLevel uint32
	DataHtml TopologyDataHtml
}

type NextTopology struct {
	Nodes []NextNode `json:"nodes"`
	Links []NextLink `json:"links"`
}

type NextNode struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	NetworkId string `json:"networkId"`
	Group     string `json:"group"`
	Type      string `json:"type"`
}

type NextLink struct {
	Id        string `json:"id"`
	Source    string `json:"source"`
	Target    string `json:"target"`
	NetworkId string `json:"networkId"`
	Type      string `json:"type"`
}

type TopologyDataHtml struct {
	Name string
	Data string
}

func (nextGo *NextUiGo) InitLogger() {
	// init logConfig
	toolLogger := nspgotools.Tools{}
	toolLogger.InitLogger("./logs/nspGo-nextUi.log", nextGo.LogLevel)
}

// func InitLogger() {
// 	// init logConfig
// 	toolLogger := nspgotools.Tools{}
// 	toolLogger.InitLogger("./logs/nspGo-restconf.log", 5)
// }

func (nextGo *NextUiGo) ReadRawTopoJsonFile(file string) []byte {
	rawFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("unable to read file: %v", err)
	}
	log.Info("Raw Topology Loaded")
	return rawFile
	//fmt.Println(string(body))
}

func (nextGo *NextUiGo) NextUiUnmarshalIetfNetworkModel(rawTopoFile []byte, ietfModel IetfNetworkStruct) {
	json.Unmarshal(rawTopoFile, &ietfModel)
	log.Debug("Lenght of Data Network Array: " + fmt.Sprint(len(ietfModel.Response.Data.Network)))
	var node NextNode
	var link NextLink

	for i := 0; i < len(ietfModel.Response.Data.Network); i++ {

		for j := 0; j < len(ietfModel.Response.Data.Network[i].Link); j++ {
			var SourceNodeID = ietfModel.Response.Data.Network[i].Link[j].Source.SourceNode
			var DestNodeID = ietfModel.Response.Data.Network[i].Link[j].Destination.DestNode

			for k := 0; k < len(ietfModel.Response.Data.Network[i].Node); k++ {
				if ietfModel.Response.Data.Network[i].Node[k].NodeID == SourceNodeID {
					SourceNodeName := ietfModel.Response.Data.Network[i].Node[k].TeNodeAugment.Te.TeNodeID.DottedQuad.String
					node.Name = SourceNodeName
					SourceNodeUuid := ietfModel.Response.Data.Network[i].Node[k].NodeID
					node.Id = SourceNodeUuid
					node.Group = "PhysicalLayer"
					nextGo.Topology.Nodes = append(nextGo.Topology.Nodes, node)
					link.Source = SourceNodeName

				}

				if ietfModel.Response.Data.Network[i].Node[k].NodeID == DestNodeID {
					DestNodeName := ietfModel.Response.Data.Network[i].Node[k].TeNodeAugment.Te.TeNodeID.DottedQuad.String
					node.Name = DestNodeName
					DestNodeUuid := ietfModel.Response.Data.Network[i].Node[k].NodeID
					node.Id = DestNodeUuid
					node.Group = "PhysicalLayer"
					nextGo.Topology.Nodes = append(nextGo.Topology.Nodes, node)
					link.Target = DestNodeName
				}
			}
			LinkUuid := ietfModel.Response.Data.Network[i].Link[j].LinkID
			link.Id = LinkUuid
			nextGo.Topology.Links = append(nextGo.Topology.Links, link)

		}
	}
	// log.Debug(nextGo)
}

func (nextGo *NextUiGo) NextUiMarshal() []byte {
	// u, err := json.Marshal(NextNode{Id: nextGo.Topology.NextNodes[0].Name})
	jsonBytes, err := json.Marshal(NextTopology{
		Nodes: nextGo.Topology.Nodes,
		Links: nextGo.Topology.Links})
	if err != nil {
		panic(err)
	}
	log.Debug(string(jsonBytes))

	return jsonBytes
}

func (nextGo *NextUiGo) NextUiHttpServer(marshaledNextuiTopo []byte) {

	type nodeData struct {
		Name string
		// Name []byte
	}
	std1 := nodeData{string(marshaledNextuiTopo)}
	tmp1 := template.New("Template_1")
	tmp1, _ = tmp1.Parse("var data = {{.Name}}")

	filePath, _ := (os.Getwd())
	file, err := os.Create(filePath + "/nspGo-topoViewer/html-template/static/js/nextData.js")

	log.Debugf("fileOutput: ", file)

	if err == nil {
		tmp1.Execute(file, std1)
	}

	// read the file again to replace "&#34;" with single quoute "
	// aaraafat-tag: code need to be optimised
	input, err := ioutil.ReadFile(filePath + "/nspGo-topoViewer/html-template/static/js/nextData.js")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	output := bytes.Replace(input, []byte("&#34;"), []byte(`"`), -1)

	if err = ioutil.WriteFile(filePath+"/nspGo-topoViewer/html-template/static/js/nextData.js", output, 0666); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// log.Debugf("filePath: ", filePath+"/nspGo-topoViewer/html-template/static/")
	http.Handle("/", http.FileServer(http.Dir(filePath+"/nspGo-topoViewer/html-template")))
	http.ListenAndServe(":8080", nil)
}
