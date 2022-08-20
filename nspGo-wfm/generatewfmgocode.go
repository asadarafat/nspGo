package nspgowfm

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/noirbizarre/gonja"
	"github.com/tidwall/gjson"
)

func GenerateWfmGoCode() {
	err := os.Remove("./nspGo-wfm/nspgowfm.go")
	if err != nil {
		log.Println(err)
	}

	// write to file
	f, err := os.OpenFile("./nspGo-wfm/nspgowfm.go",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(
		`
	package nspgowfm
	import (
		"crypto/tls"
		"time"
	
		"github.com/go-resty/resty/v2"
		log "github.com/sirupsen/logrus"
		nspgoconstants "local.com/nspgo/nspGo-constants"
		nspgotools "local.com/nspgo/nspGo-tools"

	)
	type Wfm struct {
		Payload      []byte
		ResponseData []byte
		LogLevel     uint32
	}

	func (wfm *Wfm) InitLogger() {
		// init logConfig
		toolLogger := nspgotools.Tools{}
		toolLogger.InitLogger("./logs/nspGo-wfm.log", wfm.LogLevel)
	}

	`); err != nil {
		log.Println(err)
	}

	wfmJinjaTemplate := (`

	// {{description}}
	func (wfm *Wfm) Wfm{{functionName}}(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
		client := resty.New()
		client.SetTimeout(6000 * time.Second)
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		if proxyEnable == "true" {
			client.SetProxy(proxyAddress)
		}
	
		url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "{{urlPath}}")
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("authorization", "Bearer "+token).
			SetBody(payload).
			{{method}}(url)
	
		if err != nil {
			log.Error("NspRestconfInventory is unsuccesful: ", err)
			return
		}
		log.Info("Received Response "+urlHost+" Response: ")
		log.Info(resp.String())

		return resp.String()
	
	}`)

	// jsonFile, err := ioutil.ReadFile("/home/suuser/nspGo/nspGo-wfm/21-11-wfm-swagger.json")
	jsonFile, err := ioutil.ReadFile("./nspGo-wfm/21-11-wfm-swagger.json")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	//fmt.Println(string(body))

	jsonOutPath := gjson.GetBytes(jsonFile, "paths.@keys")
	for _, path := range jsonOutPath.Array() {
		jsonOutPathMethod := gjson.GetBytes(jsonFile, "paths."+path.String()+".@keys")
		for _, method := range jsonOutPathMethod.Array() {
			descriptionJsonPath := ("paths." + path.String() + "." + method.String() + ".description")
			description := gjson.GetBytes(jsonFile, descriptionJsonPath)
			// println(description.String())
			// println(paths.String())
			// println(method.String())
			functionTitle := (strings.Title(
				strings.Replace(
					strings.Replace(
						strings.Title(strings.Replace(strings.Title(strings.Replace(path.String(), "-", "", 5)), "/", "", 5)), "{", "", 5), "}", "", 5)) + strings.Title(method.String()))
			switch method.String() {
			case "post":
				tpl, err := gonja.FromString(wfmJinjaTemplate)

				if err != nil {
					panic(err)
				}
				out, err := tpl.Execute(
					gonja.Context{
						"description":  description.String(),
						"functionName": functionTitle,
						"method":       strings.Title(method.String()),
						"urlPath":      path})

				if err != nil {
					panic(err)
				}
				if _, err := f.WriteString(out); err != nil {
					log.Println(err)
				}

			case "put":
				tpl, err := gonja.FromString(wfmJinjaTemplate)

				if err != nil {
					panic(err)
				}
				out, err := tpl.Execute(
					gonja.Context{
						"description":  description.String(),
						"functionName": functionTitle,
						"method":       strings.Title(method.String()),
						"urlPath":      path})

				if err != nil {
					panic(err)
				}
				if _, err := f.WriteString(out); err != nil {
					log.Println(err)
				}

			case "get":
				tpl, err := gonja.FromString(wfmJinjaTemplate)

				if err != nil {
					panic(err)
				}
				out, err := tpl.Execute(
					gonja.Context{
						"description":  description.String(),
						"functionName": functionTitle,
						"method":       strings.Title(method.String()),
						"urlPath":      path})

				if err != nil {
					panic(err)
				}
				if _, err := f.WriteString(out); err != nil {
					log.Println(err)
				}

			case "delete":
				tpl, err := gonja.FromString(wfmJinjaTemplate)

				if err != nil {
					panic(err)
				}
				out, err := tpl.Execute(
					gonja.Context{
						"description":  description.String(),
						"functionName": functionTitle,
						"method":       strings.Title(method.String()),
						"urlPath":      path})

				if err != nil {
					panic(err)
				}
				if _, err := f.WriteString(out); err != nil {
					log.Println(err)
				}
			}
		}
	}
}
