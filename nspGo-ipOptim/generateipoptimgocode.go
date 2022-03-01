package nspgoipoptim

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/noirbizarre/gonja"
	"github.com/tidwall/gjson"
)

func GenerateIpOptimGoCode() {
	err := os.Remove("./nspGo-ipOptim/nspgoipoptimv2.go")
	if err != nil {
		log.Println(err)
	}

	// write to file
	f, err := os.OpenFile("./nspGo-ipOptim/nspgoipoptimv2.go",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(
		`
	package nspgoipoptim
	// asadArafat
	import (
		"crypto/tls"
		"time"
	
		"github.com/go-resty/resty/v2"
		log "github.com/sirupsen/logrus"
		nspgoconstants "local.com/nspgo/nspGo-constants"
		nspgotools "local.com/nspgo/nspGo-tools"

	)

	func init() {
		// init logConfig
		toolLogger := nspgotools.Tools{}
		toolLogger.InitLogger("./logs/nspGo-ipOptim.log")
	}

	`); err != nil {
		log.Println(err)
	}

	ipOptimJinjaTemplate := (`

	// {{description}}
	func (ipO *IpOptim) Ipo{{functionName}}(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
		client := resty.New()
		client.SetTimeout(6000 * time.Second)
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		if proxyEnable == "true" {
			client.SetProxy(proxyAddress)
		}
	
		url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "{{urlPath}}")
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

	ipOptimSwagger, err := http.Get("https://network.developer.nokia.com/static/swagger/v2.0/source-json/21.11/sdn_r2111_v4_swagger.json")
	if err != nil {
		log.Fatalf("unable to dowload swagger file: %v", err)
	}
	defer ipOptimSwagger.Body.Close()

	jsonFile, err := ioutil.ReadAll(ipOptimSwagger.Body)
	if err != nil {
		log.Fatal(err)
	}

	// jsonFile, err := ioutil.ReadFile("/home/suuser/nspGo/nspGo-wfm/21-11-wfm-swagger.json")
	// if err != nil {
	// 	log.Fatalf("unable to read file: %v", err)
	// }

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
						strings.Title(
							strings.Replace(
								strings.Title(
									strings.Replace(
										strings.Replace(
											strings.Replace(path.String(), "=", "", 10), "?", "", 10), "-", "", 10)), "/", "", 10)), "{", "", 10), "}", "", 10)) + strings.Title(method.String()))
			switch method.String() {
			case "post":
				tpl, err := gonja.FromString(ipOptimJinjaTemplate)

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
				tpl, err := gonja.FromString(ipOptimJinjaTemplate)

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
				tpl, err := gonja.FromString(ipOptimJinjaTemplate)

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
				tpl, err := gonja.FromString(ipOptimJinjaTemplate)

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

			case "patch":
				tpl, err := gonja.FromString(ipOptimJinjaTemplate)

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
