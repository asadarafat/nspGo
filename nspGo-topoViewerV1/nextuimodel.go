package nspgotopoviewer

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/srl-labs/containerlab/clab"
	"github.com/srl-labs/containerlab/nodes"
	"github.com/srl-labs/containerlab/types"

	log "github.com/sirupsen/logrus"

	nspgotools "local.com/nspgo/nspGo-tools"
)

type NextUiGo struct {
	Topology NextTopology
	LogLevel uint32
}

type NextTopology struct {
	Nodes        []NextNode  `json:"nodes"`
	Links        []NextLink  `json:"links"`
	ExtraData    interface{} `json:"ExtraData,omitempty"`
	ClabTopoData ClabTopo    `json:"clabTopoData"`
}

type NextNode struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	NetworkId string      `json:"networkId"`
	Group     string      `json:"group"`
	Type      string      `json:"type"`
	ExtraData interface{} `json:"ExtraData,omitempty"`
}

type NextLink struct {
	Id             string      `json:"id"`
	Source         string      `json:"source"`
	SourceEndpoint string      `json:"source_endpoint"`
	Target         string      `json:"target"`
	TargetEndpoint string      `json:"target_endpoint"`
	NetworkId      string      `json:"networkId"`
	Group          string      `json:"group"`
	Type           string      `json:"type"`
	ExtraData      interface{} `json:"ExtraData,omitempty"`
}

type ClabTopo struct {
	ClabTopoName string                `json:"clabTopoName"`
	ClabNodes    map[string]nodes.Node `json:"clabNodes"`
	ClabLinks    map[int]*types.Link   `json:"clabLinks"`
}

type ClabNodeJson struct {
	types.ContainerDetails
	Vars map[string]interface{} `json:"vars,omitempty"`
}

type ClabLinkJson struct {
	clab.Link
	Vars map[string]interface{} `json:"vars,omitempty"`
}

type ClabTopoJson struct {
	Name  string                  `json:"name"`
	Nodes map[string]ClabNodeJson `json:"nodes,omitempty"`
	Links map[int]ClabLinkJson    `json:"links,omitempty"`
}

func (nextGo *NextUiGo) InitLogger() {
	// init logConfig
	toolLogger := nspgotools.Tools{}
	toolLogger.InitLogger("./logs/nspGo-nextUi.log", nextGo.LogLevel)
}

func (nextGo *NextUiGo) ReadRawTopoJsonFile(file string) []byte {
	rawFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("unable to read file: %v", err)
	}
	log.Info("Raw Topology Loaded")
	return rawFile
	//fmt.Println(string(body))
}

func (nextGo *NextUiGo) NextUiUnmarshalContainerLabTopo(topoFile string) error {
	log.Info(topoFile)

	filePath, _ := os.Getwd()
	filePath = (filePath + "/nspGo-topoViewerV1/rawTopoFile/")

	log.Info(filePath + topoFile)

	c, err := clab.NewContainerLab(
		clab.WithTimeout(time.Second*30),
		clab.WithTopoFile(filePath+topoFile, ""),
	)

	if err != nil {
		return err
	}

	nextGo.Topology.ClabTopoData.ClabNodes = c.Nodes
	nextGo.Topology.ClabTopoData.ClabLinks = c.Links
	nextGo.Topology.ClabTopoData.ClabTopoName = c.Config.Name

	return nil
}

func (nextGo *NextUiGo) NextUiMarshalContainerLabTopo(ClabTopoJson) []byte {
	node := NextNode{
		ExtraData: make(map[string]interface{}, 0),
	}
	link := NextLink{
		ExtraData: make(map[string]interface{}, 0),
	}

	for i, n := range nextGo.Topology.ClabTopoData.ClabNodes {
		node.Id = string(i)
		node.Name = n.Config().ShortName
		node.Group = n.Config().Group

		node.ExtraData = map[string]interface{}{
			"ClabNodeName": n.Config().ShortName,
			"eggs": struct {
				source string
				price  float64
			}{"chicken", 1.75},
			"Kind":            n.Config().Kind,
			"Image":           n.Config().Image,
			"Group":           n.Config().Group,
			"MgmtIPv4Address": n.Config().MgmtIPv4Address,
			"MgmtIPv6Address": n.Config().MgmtIPv6Address,
		}

		nextGo.Topology.Nodes = append(nextGo.Topology.Nodes, node)
	}

	for i, l := range nextGo.Topology.ClabTopoData.ClabLinks {
		link.Id = strconv.Itoa(i)
		link.Source = l.A.Node.ShortName
		link.SourceEndpoint = l.A.EndpointName
		link.Target = l.B.Node.ShortName
		link.TargetEndpoint = l.B.EndpointName
		nextGo.Topology.Links = append(nextGo.Topology.Links, link)
	}

	log.Info(nextGo.Topology.Nodes)
	jsonBytesNextUi, err := json.MarshalIndent(NextTopology{
		Nodes: nextGo.Topology.Nodes,
		Links: nextGo.Topology.Links}, "", "  ")
	if err != nil {
		panic(err)
	}

	_, err = os.Stdout.Write(jsonBytesNextUi)
	log.Info("NextUiMarshal Result:", string(jsonBytesNextUi))
	return jsonBytesNextUi
}

func (nextGo *NextUiGo) NextUiUnmarshalIetfNetworkModel(rawTopoFile []byte, ietfModel IetfNetworkStruct) {
	json.Unmarshal(rawTopoFile, &ietfModel)
	log.Debugf("Lenght of Data Network Array: ", (len(ietfModel.Response.Data.Network)))
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
					node.Group = "IgpLayer"
					node.NetworkId = ietfModel.Response.Data.Network[i].NetworkID
					nextGo.Topology.Nodes = append(nextGo.Topology.Nodes, node)

					link.Source = SourceNodeName
					link.SourceEndpoint = ("From-" + SourceNodeName)
				}

				if ietfModel.Response.Data.Network[i].Node[k].NodeID == DestNodeID {
					DestNodeName := ietfModel.Response.Data.Network[i].Node[k].TeNodeAugment.Te.TeNodeID.DottedQuad.String
					node.Name = DestNodeName
					DestNodeUuid := ietfModel.Response.Data.Network[i].Node[k].NodeID
					node.Id = DestNodeUuid
					node.Group = "IgpLayer"
					node.NetworkId = ietfModel.Response.Data.Network[i].NetworkID
					nextGo.Topology.Nodes = append(nextGo.Topology.Nodes, node)

					link.Target = DestNodeName
					link.TargetEndpoint = ("From-" + DestNodeName)

				}
			}
			LinkUuid := ietfModel.Response.Data.Network[i].Link[j].LinkID
			link.Id = LinkUuid
			link.NetworkId = ietfModel.Response.Data.Network[i].NetworkID
			nextGo.Topology.Links = append(nextGo.Topology.Links, link)

		}
	}
	// log.Debug(nextGo)
}

func (nextGo *NextUiGo) NextUiUnmarshalNetSupPhysicalModel(rawTopoFile []byte, netSupModel NetSupPhysicalStruct) {
	json.Unmarshal(rawTopoFile, &netSupModel)
	var node NextNode
	var link NextLink
	for i := 0; i < len(netSupModel.Response.Data); i++ {
		// log.Debugf("sourceNode: ", netSupModel.Response.Data[i].Endpoints[0].ParentNeID)
		// log.Debugf("destNode: ", netSupModel.Response.Data[i].Endpoints[1].ParentNeID)

		// append pyhsical source Node in pyhsical Topology
		SourceNodeName := netSupModel.Response.Data[i].Endpoints[0].ParentNeID
		node.Name = ("phy-" + SourceNodeName)
		node.Group = "PhysicalLayer"
		nextGo.Topology.Nodes = append(nextGo.Topology.Nodes, node)

		// append pyhsical destination Node in pyhsical Topology
		DestNodeName := netSupModel.Response.Data[i].Endpoints[1].ParentNeID
		node.Name = ("phy-" + DestNodeName)
		node.Group = "PhysicalLayer"
		nextGo.Topology.Nodes = append(nextGo.Topology.Nodes, node)

		// append pyhsical link in pyhsical Topology
		link.Source = ("phy-" + SourceNodeName)
		link.SourceEndpoint = netSupModel.Response.Data[i].Endpoints[0].Name
		link.Target = ("phy-" + DestNodeName)
		link.TargetEndpoint = netSupModel.Response.Data[i].Endpoints[1].Name
		link.Id = strconv.Itoa(i)
		nextGo.Topology.Links = append(nextGo.Topology.Links, link)
	}
}

func (nextGo *NextUiGo) NextUiAppendInterLayerLinks(Nodes []NextNode) {
	// u, err := json.Marshal(NextNode{Id: nextGo.Topology.NextNodes[0].Name})
	var link NextLink
	var linkId int = 1
	for i := 0; i < len(Nodes); i++ {
		if !strings.Contains(Nodes[i].Name, "phy-") { // find Igp layer nodes
			link.Source = Nodes[i].Name
			link.SourceEndpoint = "igp"
			for j := 0; j < len(Nodes); j++ {
				if Nodes[j].Name == ("phy-" + Nodes[i].Name) { // find Phy Layer node with same name with Igp Layer node

					link.Target = Nodes[j].Name
					link.TargetEndpoint = "phy"

					link.Id = strconv.Itoa(linkId + j)
				}
			}
			link.Group = "interTopoLayerLink"
			nextGo.Topology.Links = append(nextGo.Topology.Links, link)
		}

	}
}

func (nextGo *NextUiGo) NextUiMarshal() []byte {
	// u, err := json.Marshal(NextNode{Id: nextGo.Topology.NextNodes[0].Name})
	jsonBytes, err := json.Marshal(NextTopology{
		Nodes: nextGo.Topology.Nodes,
		Links: nextGo.Topology.Links})
	if err != nil {
		panic(err)
	}
	log.Debugf("NextUiMarshal Result:", string(jsonBytes))

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
	file, err := os.Create(filePath + "/nspGo-topoViewerV1/html-template/static/js/nextData.js")

	log.Debugf("fileOutput: ", file)

	if err == nil {
		tmp1.Execute(file, std1)
	}

	// read the file again to replace "&#34;" with single quoute "
	// aaraafat-tag: code need to be optimised, no need to read/write again file nextData.js
	input, err := ioutil.ReadFile(filePath + "/nspGo-topoViewerV1/html-template-nextUi/static/js/nextData.js")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	output := bytes.Replace(input, []byte("&#34;"), []byte(`"`), -1)

	if err = ioutil.WriteFile(filePath+"/nspGo-topoViewerV1/html-template-nextUi/static/js/nextData.js", output, 0666); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// log.Debugf("filePath: ", filePath+"/nspGo-topoViewer/html-template/static/")
	http.Handle("/", http.FileServer(http.Dir(filePath+"/nspGo-topoViewerV1/html-template-nextUi")))
	http.ListenAndServe(":8080", nil)
	log.Infof("serving at ")
}
