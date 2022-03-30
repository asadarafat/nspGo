package nsptopoviewer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	nspgotools "local.com/nspgo/nspGo-tools"
)

type NextUiGo struct {
	Topology NextTopology
	LogLevel uint32
}

type NextTopology struct {
	Nodes []NextNode `json:"nodes"`
	Links []NextLink `json:"links"`
}

type NextNode struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Group string `json:"group"`
	Kind  string `json:"kind"`
}

type NextLink struct {
	Id     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
	Kind   string `json:"kind"`
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
					// SourceNodeId := ietfModel.Response.Data.Network[i].Node[k].NodeID
					// node.id = SourceNodeId
					nextGo.Topology.Nodes = append(nextGo.Topology.Nodes, node)
					link.Source = SourceNodeName
				}

				if ietfModel.Response.Data.Network[i].Node[k].NodeID == DestNodeID {
					DestNodeName := ietfModel.Response.Data.Network[i].Node[k].TeNodeAugment.Te.TeNodeID.DottedQuad.String
					node.Name = DestNodeName
					// DestNodeId := ietfModel.Response.Data.Network[i].Node[k].NodeID
					// node.id = DestNodeId
					nextGo.Topology.Nodes = append(nextGo.Topology.Nodes, node)
					link.Target = DestNodeName

				}
			}
			nextGo.Topology.Links = append(nextGo.Topology.Links, link)
		}
	}
	// log.Debug(nextGo)
}

func (nextGo *NextUiGo) NextUiMarshal() string {
	// u, err := json.Marshal(NextNode{Id: nextGo.Topology.NextNodes[0].Name})
	jsonBytes, err := json.Marshal(NextTopology{
		Nodes: nextGo.Topology.Nodes,
		Links: nextGo.Topology.Links})
	if err != nil {
		panic(err)
	}
	log.Debug(string(jsonBytes))
	return string(jsonBytes)
}
