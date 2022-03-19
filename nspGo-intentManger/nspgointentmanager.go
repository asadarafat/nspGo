package nspgointentmanger

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	nspgoconstants "local.com/nspgo/nspGo-constants"
	nspgotools "local.com/nspgo/nspGo-tools"
)

type IntentManager struct {
	Payload      []byte
	ResponseData []byte
	LogLevel     uint32
}

func formatJSON(data []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "    ")
	if err == nil {
		return out.Bytes(), err
	}
	return data, nil
}

func (rConf *IntentManager) InitLogger() {
	// init logConfig
	toolLogger := nspgotools.Tools{}
	toolLogger.InitLogger("./logs/nspGo-IntentManager.log", rConf.LogLevel)
}

func (rConf *IntentManager) ReadIntentManagerPayload(file string) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	rConf.Payload = body
	//fmt.Println(string(body))
}

func (rConf *IntentManager) NspIntentManagerPost(urlHost string, token string, proxyEnable string, proxyAddress string, urlPath string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}
	//asycn == false
	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IM_BASE_URL + urlPath)
	resp, err := client.R().
		SetHeader("Accept", "application/yang-data+json").
		SetHeader("Content-Type", "application/yang-patch+json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)
	log.Debug("url: ", url)

	// re, errRe := regexp.Compile(`[\r?\n?\t]`)
	// if errRe != nil {
	// 	log.Fatal(err)
	// }
	// payloadString := re.ReplaceAllString(string(payload), " ")
	// log.Debug(strings.Replace(payloadString, "\\", "/", -1))
	log.Debug("Payload: ")
	if log.GetLevel() == log.DebugLevel {
		strSlice := strings.Split(string(payload), "\n")
		for _, s := range strSlice {
			log.Debug(s)
		}
	}

	log.Debug("Response: ")

	prettyJSON, errPrettyJSON := formatJSON([]byte(resp.String()))
	if errPrettyJSON != nil {
		log.Fatal(err)
	}
	log.Debug(string(prettyJSON))

	// log.Debug("NspIntentManagerPost Response: ", resp.String())
	if err != nil {
		log.Error("NspIntentManagerPost is unsuccesful: ", err)
		return
	}
	return resp.String()

}

func (rConf *IntentManager) NspIntentManagerGet(urlHost string, token string, proxyEnable string, proxyAddress string, urlPath string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}
	//asycn == false
	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IM_BASE_URL + urlPath)
	resp, err := client.R().
		SetHeader("Accept", "application/yang-data+json").
		SetHeader("Content-Type", "application/yang-patch+json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)
	log.Debug("url: ", url)
	log.Debug("NspIntentManagerGet Response: ", resp.String())
	log.Debug("Response: ", result)
	if err != nil {
		log.Error("NspIntentManagerGet is unsuccesful: ", err)
		return
	}
	return resp.String()
}

func (rConf *IntentManager) NspIntentManagerDel(urlHost string, token string, proxyEnable string, proxyAddress string, urlPath string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}
	//asycn == false
	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IM_BASE_URL + urlPath)
	resp, err := client.R().
		SetHeader("Accept", "application/yang-data+json").
		SetHeader("Content-Type", "application/yang-patch+json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)
	log.Debug("url: ", url)
	log.Debug("NspIntentManagerDel Response: ", resp.String())
	log.Debug("Response: ", result)
	if err != nil {
		log.Error("NspIntentManagerDel is unsuccesful: ", err)
		return
	}
	return resp.String()
}

// RestConf xPath
// xPath := "/root/nokia-conf:configure/router=Base"
// xPath := "/root/nokia-conf:configure"
