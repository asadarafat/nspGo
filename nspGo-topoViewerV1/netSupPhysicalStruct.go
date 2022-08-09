package nspgotopoviewer

type NetSupPhysicalStruct struct {
	Response struct {
		Status    int `json:"status"`
		StartRow  int `json:"startRow"`
		EndRow    int `json:"endRow"`
		TotalRows int `json:"totalRows"`
		Data      []struct {
			Fdn                string        `json:"fdn"`
			SourceType         string        `json:"sourceType"`
			SourceSystem       string        `json:"sourceSystem"`
			Sources            []string      `json:"sources"`
			Name               string        `json:"name"`
			Type               string        `json:"type"`
			Rate               interface{}   `json:"rate"`
			ActualRate         interface{}   `json:"actualRate"`
			ActualRateUnits    string        `json:"actualRateUnits"`
			Description        string        `json:"description"`
			Protection         string        `json:"protection"`
			ProtectionKind     string        `json:"protectionKind"`
			Latency            int           `json:"latency"`
			LatencyUnits       string        `json:"latencyUnits"`
			Restoration        string        `json:"restoration"`
			Direction          string        `json:"direction"`
			AdminState         string        `json:"adminState"`
			OperState          string        `json:"operState"`
			StandbyState       string        `json:"standbyState"`
			AvailabilityStates []interface{} `json:"availabilityStates"`
			ObjectDetails      struct {
				LinkDiscoveredFrom string `json:"linkDiscoveredFrom"`
				LinkScope          string `json:"linkScope"`
				IsLagMember        string `json:"isLagMember"`
				CableType          string `json:"cableType"`
				LinkType           string `json:"linkType"`
				HasLLDPAdjacency   string `json:"hasLLDPAdjacency"`
			} `json:"objectDetails"`
			Endpoints []struct {
				Name                string      `json:"name"`
				Type                string      `json:"type"`
				Port                string      `json:"port"`
				Lag                 interface{} `json:"lag"`
				ParentNe            string      `json:"parentNe"`
				ParentNeID          string      `json:"parentNeId"`
				SubnetName          interface{} `json:"subnetName"`
				ObjectDetailsSource struct {
					IsSource              string `json:"isSource"`
					HasUnderlyingEndpoint string `json:"hasUnderlyingEndpoint"`
				} `json:"objectDetailsSource,omitempty"`
				ParentNeName        interface{} `json:"parentNeName"`
				ObjectDetailsTarget struct {
					IsTarget string `json:"isTarget"`
				} `json:"objectDetailsTarget,omitempty"`
			} `json:"endpoints"`
			Links []struct {
				Rel  string `json:"rel"`
				Href string `json:"href"`
			} `json:"links"`
			ID string `json:"id"`
		} `json:"data"`
		Errors interface{} `json:"errors"`
	} `json:"response"`
}
