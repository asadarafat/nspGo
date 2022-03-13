package nspgorestconf

import (
	"crypto/tls"
	"io/ioutil"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	nspgoconstants "local.com/nspgo/nspGo-constants"
	nspgotools "local.com/nspgo/nspGo-tools"
)

type RestConf struct {
	Payload      []byte
	ResponseData []byte
}

func init() {
	// init logConfig
	toolLogger := nspgotools.Tools{}
	toolLogger.InitLogger("./logs/nspGo-restconf.log", 5)
}

func (rConf *RestConf) ReadRestConfPayload(file string) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	rConf.Payload = body
	//fmt.Println(string(body))
}

func (rConf *RestConf) PatchRestConfNetworkDevice(urlHost string, token string, proxyEnable string, proxyAddress string, neId string, xPath string, payload []byte, asycn bool) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	if asycn { //asycn == true
		url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_RESTCONF_BASE_URL + "data/network-device-mgr:network-devices/network-device=" + neId + xPath + "?async=true")
		log.Info("Sending Request: " + neId + " via" + nspgoconstants.GLBL_NSP_RESTCONF_BASE_URL)
		resp, err := client.R().
			SetHeader("Accept", "application/yang-data+json").
			SetHeader("Content-Type", "application/yang-patch+json").
			SetHeader("authorization", "Bearer "+token).
			SetBody(payload).
			Patch(url)

		// log.Info("PatchRestConfNetworkDevice is success: ", resp.IsSuccess())
		// log.Debug("PatchRestConfNetworkDevice debug ", resp)
		log.Info("PatchRestConfNetworkDevice "+neId+" Response: ", resp.String())
		//fmt.Println(result)
		if err != nil {
			log.Error("PatchRestConfNetworkDevice is unsuccesful: ", err)
			return
		}
		return resp.String()
	} else { //asycn == false
		url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_RESTCONF_BASE_URL + "data/network-device-mgr:network-devices/network-device=" + neId + xPath)
		log.Info("Sending Request: " + neId + " via" + nspgoconstants.GLBL_NSP_RESTCONF_BASE_URL)
		resp, err := client.R().
			SetHeader("Accept", "application/yang-data+json").
			SetHeader("Content-Type", "application/yang-patch+json").
			SetHeader("authorization", "Bearer "+token).
			SetBody(payload).
			Patch(url)

		// log.Info("PatchRestConfNetworkDevice is success: ", resp.IsSuccess())
		// log.Debug("PatchRestConfNetworkDevice debug ", resp)
		log.Info("PatchRestConfNetworkDevice "+neId+" Response: ", resp.String())
		//fmt.Println(result)
		if err != nil {
			log.Error("PatchRestConfNetworkDevice is unsuccesful: ", err)
			return
		}
	}
	return
}

func (rConf *RestConf) PatchRestConfMdc(urlHost string, token string, proxyEnable string, proxyAddress string, neId string, xPath string, payload []byte, asycn bool) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	if asycn { //asycn == true
		url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_MDC_BASE_URL + "data/network-device-mgr:network-devices/network-device=" + neId + xPath + "?async=true")
		log.Info("Sending Request: " + neId + " via " + nspgoconstants.GLBL_NSP_MDC_BASE_URL)
		resp, err := client.R().
			SetHeader("Accept", "application/yang-data+json").
			SetHeader("Content-Type", "application/yang-patch+json").
			SetHeader("authorization", "Bearer "+token).
			SetBody(payload).
			Patch(url)

		// log.Info("PatchRestConfNetworkDevice is success: ", resp.IsSuccess())
		// log.Debug("PatchRestConfNetworkDevice debug ", resp)
		log.Info("PatchRestConfMdc "+neId+" Response: ", resp.String())
		//fmt.Println(result)
		if err != nil {
			log.Error("PatchRestConfMdc is unsuccesful: ", err)
			return
		}
		return resp.String()

	} else { //asycn == false
		url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_MDC_BASE_URL + "data/network-device-mgr:network-devices/network-device=" + neId + xPath)
		log.Info("Sending Request: " + neId + " via " + nspgoconstants.GLBL_NSP_MDC_BASE_URL)
		resp, err := client.R().
			SetHeader("Accept", "application/yang-data+json").
			SetHeader("Content-Type", "application/yang-patch+json").
			SetHeader("authorization", "Bearer "+token).
			SetBody(payload).
			Patch(url)

		// log.Debug("PatchRestConfNetworkDevice debug ", resp)
		// log.Info("Payload: " + string(payload))
		// log.Info("URL: " + url)
		log.Info("Received Response "+neId+" Response: ", resp.String())
		//fmt.Println(result)
		if err != nil {
			log.Error("PatchRestConfMdc is unsuccesful: ", err)
			return
		}
	}
	return

}

func (rConf *RestConf) NspRestconfInventory(urlHost string, token string, proxyEnable string, proxyAddress string, urlPath string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	//asycn == false
	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_RESTCONF_BASE_URL + urlPath)
	resp, err := client.R().
		SetHeader("Accept", "application/yang-data+json").
		SetHeader("Content-Type", "application/yang-patch+json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	// log.Info("NspRestconfInventory Response: ", resp.String())
	// fmt.Println(result)
	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	return resp.String()

}

// RestConf xPath
// xPath := "/root/nokia-conf:configure/router=Base"
// xPath := "/root/nokia-conf:configure"
