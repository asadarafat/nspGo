package nspgotopoviewer

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	//nsptopoviewer "local.com/nspgo/nspGo-topoViewer"
)

type CyGraph struct {
	GraphNodes []CyNode `json:"nodes"`
	GraphLinks []CyLink `json:"links"`
}

type CyNode struct {
	Data struct {
		ID            string `json:"id"`
		IdInt         string `json:"idInt"`
		Name          string `json:"name"`
		SystemAddress string `json:"systemAddress"`
		RouterType    string `json:"routerType"`
	} `json:"data"`
	Group string `json:"group"`
	Pos   struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"position"`
	Removed    bool   `json:"removed"`
	Selected   bool   `json:"selected"`
	Selectable bool   `json:"selectable"`
	Locked     bool   `json:"locked"`
	Grabbed    bool   `json:"grabbed"`
	Grabbable  bool   `json:"grabbable"`
	Classes    string `json:"classes"`
}

type CyLink struct {
	Data struct {
		ID     string `json:"id"`
		Source string `json:"source,omitempty"`
		Target string `json:"target,omitempty"`
		Weight string `json:"weight"`
		Group  string `json:"group"`
		Name   string `json:"name"`
	} `json:"data"`
	Pos struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"position"`
	Removed    bool   `json:"removed"`
	Group      string `json:"group"`
	Selected   bool   `json:"selected"`
	Selectable bool   `json:"selectable"`
	Locked     bool   `json:"locked"`
	Grabbed    bool   `json:"grabbed"`
	Grabbable  bool   `json:"grabbable"`
	Classes    string `json:"classes"`
}

func (cyGraph *CyGraph) AppendCytoNode(id string, idInt string, hostName string, systemAddress string, routerType string, classes string) (result CyGraph) {
	var cyNode = CyNode{
		Data: struct {
			ID            string "json:\"id\""
			IdInt         string "json:\"idInt\""
			Name          string "json:\"name\""
			SystemAddress string "json:\"systemAddress\""
			RouterType    string "json:\"routerType\""
		}{ID: id, IdInt: idInt, Name: hostName, SystemAddress: systemAddress, RouterType: routerType},
		Group:      "nodes",
		Removed:    false,
		Selected:   false,
		Selectable: true,
		Locked:     false,
		Grabbed:    false,
		Grabbable:  true,
		Classes:    classes,
	}
	cyGraph.GraphNodes = append(cyGraph.GraphNodes, cyNode)
	return
}

func (cyGraph *CyGraph) AppendCytoLink(id string, source string, target string, weight string, group string, name string) (result CyGraph) {
	var cyLink = CyLink{
		Data: struct {
			ID     string "json:\"id\""
			Source string `json:"source,omitempty"`
			Target string `json:"target,omitempty"`
			Weight string `json:"weight"`
			Group  string `json:"group"`
			Name   string `json:"name"`
		}{ID: id, Source: source, Target: target, Weight: weight, Group: group, Name: name},
		Group:      "edges",
		Removed:    false,
		Selected:   false,
		Selectable: true,
		Locked:     false,
		Grabbed:    false,
		Grabbable:  true,
		Classes:    "asad",
	}
	cyGraph.GraphLinks = append(cyGraph.GraphLinks, cyLink)
	return
}

func (cyGraph *CyGraph) MarshalCyGraph() (result string) {

	log.Info("Marshaling to CytoScape Json Start...")

	var str strings.Builder
	//var p map[string]interface{}
	str.WriteString("[")
	for i := 0; i < len(cyGraph.GraphNodes); i++ {
		if jsn3, err := json.Marshal(cyGraph.GraphNodes[i]); err == nil {
			str.WriteString(string(jsn3))
			str.WriteString(",")
		} else {
			log.Panic(err)
		}
	}
	for i := 0; i < len(cyGraph.GraphLinks); i++ {
		//skip the empty link source or target
		if cyGraph.GraphLinks[i].Data.Source == "" || cyGraph.GraphLinks[i].Data.Target == "" {
			continue
		}
		if jsn4, err := json.Marshal(cyGraph.GraphLinks[i]); err == nil {
			str.WriteString(string(jsn4))
			if i < (len(cyGraph.GraphLinks) - 1) {
				str.WriteString(",")
			}

		} else {
			log.Panic(err)
		}
	}
	str.WriteString("]")

	log.Info("Marshaling to CytoScape Json Finished...")
	log.Debug(str.String())

	return (str.String())
}

func (cyGraph *CyGraph) DumpIetfNetworkToCyGraph(content []byte, basket IetfNetworkStruct, resultFilePath string) {

	json.Unmarshal(content, &basket)

	//fmt.Println(basket.Response.Data.Network[0].NetworkID)
	log.Debug("Lenght of Data Network Array: " + fmt.Sprint(len(basket.Response.Data.Network)))

	var SourceNodeName string
	var DestNodeName string

	for i := 0; i < len(basket.Response.Data.Network); i++ {

		for j := 0; j < len(basket.Response.Data.Network[i].Link); j++ {
			var SourceNodeID = basket.Response.Data.Network[i].Link[j].Source.SourceNode
			//fmt.Println("source: " + SourceNodeID)
			var DestNodeID = basket.Response.Data.Network[i].Link[j].Destination.DestNode
			//fmt.Println("destination: " + DestNodeID)

			for k := 0; k < len(basket.Response.Data.Network[i].Node); k++ {
				if basket.Response.Data.Network[i].Node[k].NodeID == SourceNodeID {
					//fmt.Println("sourceIP: " + basket.Response.Data.Network[i].Node[k].TeNodeAugment.Te.TeNodeID.DottedQuad.String)
					SourceNodeName = basket.Response.Data.Network[i].Node[k].TeNodeAugment.Te.TeNodeID.DottedQuad.String
					cyGraph.AppendCytoNode(SourceNodeName, strconv.Itoa(i)+"-"+strconv.Itoa(k), SourceNodeName, SourceNodeName, "7750", "classes string")
				}

				if basket.Response.Data.Network[i].Node[k].NodeID == DestNodeID {
					//fmt.Println("destinationIP: " + basket.Response.Data.Network[i].Node[k].TeNodeAugment.Te.TeNodeID.DottedQuad.String)
					DestNodeName = basket.Response.Data.Network[i].Node[k].TeNodeAugment.Te.TeNodeID.DottedQuad.String
					cyGraph.AppendCytoNode(DestNodeName, strconv.Itoa(i)+"-"+strconv.Itoa(k), DestNodeName, DestNodeName, "7750", "classes string")
				}
			}
			cyGraph.AppendCytoLink(strconv.Itoa(i)+"-"+strconv.Itoa(j), SourceNodeName, DestNodeName, "2", "ip", "link-"+SourceNodeName+"-"+DestNodeName)
		}
	}
	// Create file

	file, err := os.Create(resultFilePath)
	if err != nil {
		log.Panic("Could not create json file for graph")
	}

	// Write to file
	_, err = file.Write([]byte(cyGraph.MarshalCyGraph()))
	if err != nil {
		log.Panic("Could not write json to file")
	}

}
