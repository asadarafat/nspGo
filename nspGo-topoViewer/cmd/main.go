package main

import (
	nspgotools "local.com/nspgo/nspGo-tools"
	nsptopoviewer "local.com/nspgo/nspGo-topoViewer"
)

func main() {
	// graph1 := nsptopoviewer.CyGraph{}

	// content, err := ioutil.ReadFile("../../ietfNetwork.json")
	// if err != nil {
	// 	log.Fatal("Error when opening file: ", err)
	// }
	// filePath, _ := os.Getwd()
	// filePath = (filePath + "../../../vis-library/colajs-asad-graph/data-cytoMarshall.json")
	// graph1.DumpIetfNetworkToCyGraph(content, nsptopoviewer.IetfNetworkStruct{}, filePath)

	nextUiGo := nsptopoviewer.NextUiGo{}
	nextUiGo.LogLevel = 5
	// nextUiGo.ReadRawTopoJsonFile("./nspGo-topoViewer/L3networks.json")
	nextUiGo.InitLogger()
	ietfRawFile := nextUiGo.ReadRawTopoJsonFile("./nspGo-topoViewer/ietfNetwork.json")

	// value := gjson.Get(string(nextUiGo.RawTopo), "response.data.network")
	// log.Info(value)

	nextUiGo.NextUiUnmarshalIetfNetworkModel(ietfRawFile, nsptopoviewer.IetfNetworkStruct{})
	marshaledNextuiTopo := nextUiGo.NextUiMarshal()

	tools := nspgotools.Tools{}
	tools.WriteDataToFile(marshaledNextuiTopo, "next.json")

}
