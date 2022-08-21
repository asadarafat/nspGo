package main

import (
	"io/ioutil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/noirbizarre/gonja"
	"github.com/tidwall/gjson"
)

// nspgowfm "local.com/nspgo/nspGo-wfm"

func main() {

	////  for below sourceCode auto generated
	// nspgowfm.GenerateWfmGoCode()

	// init class session
	// p := nspgosession.Session{}
	// p.LoadConfig()

	// log.Info("nsp.nspOsIP :", p.IpAdressNspOs)
	// log.Info("nsp.nspIprcIP :", p.IpAdressIprc)
	// log.Info("nsp.Username :", p.Username)
	// log.Info("nsp.Password :", p.Password)
	// log.Info("nsp.linetoken :", p.Token)

	// p.EncodeUserName()
	// log.Info(p.EncodeUserName())

	// p.GetRestToken()
	// log.Info("nsp.linetoken_NEW :", p.Token)

	//// execute auto generated code
	// w := nspgowfm.Wfm{}

	// payload := " "
	// w.WfmV1WorkflowGet(p.IpAdressNspOs, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, []byte(payload))

	wfmJinjaTemplate := (`
	{% if codeHeader == "true" %}
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

	//
	// NSP 22.06
	//
 
	func (wfm *Wfm) InitLogger() {
		// init logConfig
		toolLogger := nspgotools.Tools{}
		toolLogger.InitLogger("./logs/nspGo-wfm.log", wfm.LogLevel)
	}
	{% endif %}
	{% if method == "Post" %}
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
	}
	{% elif method == "Put" %}
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
	}
	{% elif method == "Get" %}
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
	}
	{% elif method == "Patch" %}
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
	}
	{% elif method == "Delete" %}
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
	}
	{% else %}
    	// Jinja2 output {{method}}
	{% endif %}
	`)

	// Read swager file
	swagerJsonFile, err := ioutil.ReadFile("./nspGo-wfm/22-06-wfm-swagger.json")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	// Prepare generated file
	errRemove := os.Remove("./nspGo-wfm/asadWfmCodeGenerated.go")
	if err != nil {
		log.Println(errRemove)
	}

	f, err := os.OpenFile("./nspGo-wfm/asadWfmCodeGenerated.go",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	// Write code header to generated file
	tpl, err := gonja.FromString(wfmJinjaTemplate)
	if err != nil {
		panic(err)
	}
	out, err := tpl.Execute(
		gonja.Context{
			"codeHeader": "true"})
	// fmt.Println(out)
	if _, err := f.WriteString(out); err != nil {
		log.Println(err)
	}

	jsonOutPath := gjson.GetBytes(swagerJsonFile, "paths.@keys")
	for _, path := range jsonOutPath.Array() {
		jsonOutPathMethod := gjson.GetBytes(swagerJsonFile, "paths."+path.String()+".@keys")
		for _, method := range jsonOutPathMethod.Array() {
			descriptionJsonPath := ("paths." + path.String() + "." + method.String() + ".description")
			description := gjson.GetBytes(swagerJsonFile, descriptionJsonPath)
			// println(description.String())
			// println(path.String())
			// println(method.String())
			functionTitle := (strings.Title(
				strings.Replace(
					strings.Replace(
						strings.Title(strings.Replace(strings.Title(strings.Replace(path.String(), "-", "", 5)), "/", "", 5)), "{", "", 5), "}", "", 5)) + strings.Title(method.String()))

			// Write code function to generated file
			switch method.String() {
			case "post":
				tpl, err := gonja.FromString(wfmJinjaTemplate)

				if err != nil {
					panic(err)
				}

				out, err := tpl.Execute(
					gonja.Context{
						"codeHeader":   "false",
						"description":  description.String(),
						"functionName": functionTitle,
						// "method":       "Post",
						"method":  cases.Title(language.Und, cases.NoLower).String(method.String()), // make method string from post to Post
						"urlPath": path})

				if err != nil {
					panic(err)
				}
				// fmt.Println(out)
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
						"codeHeader":   "false",
						"description":  description.String(),
						"functionName": functionTitle,
						"method":       cases.Title(language.Und, cases.NoLower).String(method.String()),
						"urlPath":      path})

				if err != nil {
					panic(err)
				}
				// fmt.Println(out)
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
						"codeHeader":   "false",
						"description":  description.String(),
						"functionName": functionTitle,
						"method":       cases.Title(language.Und, cases.NoLower).String(method.String()),
						"urlPath":      path})

				if err != nil {
					panic(err)
				}
				// fmt.Println(out)
				if _, err := f.WriteString(out); err != nil {
					log.Println(err)
				}

			case "patch":
				tpl, err := gonja.FromString(wfmJinjaTemplate)

				if err != nil {
					panic(err)
				}

				out, err := tpl.Execute(
					gonja.Context{
						"codeHeader":   "false",
						"description":  description.String(),
						"functionName": functionTitle,
						"method":       cases.Title(language.Und, cases.NoLower).String(method.String()),
						"urlPath":      path})

				if err != nil {
					panic(err)
				}
				// fmt.Println(out)
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
						"codeHeader":   "false",
						"description":  description.String(),
						"functionName": functionTitle,
						"method":       cases.Title(language.Und, cases.NoLower).String(method.String()),
						"urlPath":      path})

				if err != nil {
					panic(err)
				}
				// fmt.Println(out)
				if _, err := f.WriteString(out); err != nil {
					log.Println(err)
				}

			}
		}
	}
}
