package nspgosession

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	nspgoconstants "local.com/nspgo/nspGo-constants"
	nspgotools "local.com/nspgo/nspGo-tools"
)

type Session struct {
	IpAdressNspOs string
	IpAdressIprc  string
	Username      string
	Password      string
	Base64Str     string
	Token         string
	Proxy         Proxy
	LogLevel      uint32
}

type Proxy struct {
	Enable       string
	ProxyAddress string
}

func (s *Session) InitLogger() {
	// init logConfig
	toolLogger := nspgotools.Tools{}
	toolLogger.InitLogger("./logs/nspGo-Session.log", s.LogLevel)
}

func (s *Session) LoadConfig() {
	viper.SetConfigName("nspConfig") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./nspGo-session/cmd") // config file path
	viper.AutomaticEnv()                       // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	// Set default value
	viper.SetDefault("nsp.linetoken", "DefaultLineTokenValue")

	// pass to struct data
	//env := viper.GetString("nsp.env")
	s.IpAdressNspOs = viper.GetString("nsp.nspOsIP")
	s.IpAdressIprc = viper.GetString("nsp.nspIprcIP")
	s.Username = viper.GetString("nsp.username")
	s.Password = viper.GetString("nsp.Password")
	s.Token = viper.GetString("nsp.linetoken")
	s.Proxy.Enable = viper.GetString("nsp.proxy.enable")
	s.Proxy.ProxyAddress = viper.GetString("nsp.proxy.proxyAddress")

	// Print
	// fmt.Println("---------- Example ----------")
	// fmt.Println("nsp.env :", env)
	// fmt.Println("nsp.nspOsIP :", s.IpAdressNspOs)
	// fmt.Println("nsp.nspIprcIP :", s.IpAdressIprc)
	// fmt.Println("nsp.linetoken :", s.IpAdressIprc)

}

func (s *Session) EncodeUserName() string {
	var plainCredentials strings.Builder
	plainCredentials.WriteString(s.Username)
	plainCredentials.WriteString(":")
	plainCredentials.WriteString(s.Password)

	//fmt.Println(plainCredentials.String())

	var base64Credentials string = base64.StdEncoding.EncodeToString([]byte(plainCredentials.String()))
	//fmt.Println(base64Credentials)
	s.Base64Str = base64Credentials
	return s.Base64Str
}

func (s *Session) GetExample() {
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")

	if err != nil {
		fmt.Println(err)
		return
	}
	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
}

func (s *Session) GetRestToken() {
	client := resty.New()
	client.SetTimeout(36000000 * time.Minute)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// client.SetProxy("http://138.203.39.123:80")
	if s.Proxy.Enable == "true" {
		client.SetProxy(s.Proxy.ProxyAddress)
	}

	// POST JSON string
	// No need to set content type, if you have client level setting
	//
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Basic "+s.Base64Str).
		SetBody(`{ "grant_type": "client_credentials" }`).
		//Post("https://172.23.160.37/rest-gateway/rest/api/v1/auth/token")
		Post("https://" + s.IpAdressNspOs + nspgoconstants.GLBL_NSP_SESSION_URL_TOKEN)

	log.Info("get token is success: ", resp.IsSuccess())
	log.Debug("get token debugL ", resp)

	if err != nil {
		log.Error("token revoke not succesful", err)
		return
	}

	// fmt.Println("  resp type       :")
	// fmt.Println(reflect.TypeOf(resp))

	//// print the response body
	//
	// s.Token = string(resp.Body())
	// fmt.Println("  Token       :\n", s.Token)
	// fmt.Println("Response Body Access Token:")
	// fmt.Println(jsonparser.GetString([]byte(string(resp.Body())), "access_token"))

	//// jsonparser.GetString return two variable "VEtOLWFkbWluYzkyM2RlNjMtZjJhNy00ZGUxLThlMmUtNGUxZjBiMzcyMDM3" type string and <nil> type error"
	//// the code only interested is the first one.
	//
	// var token string
	// fmt.Println(reflect.TypeOf(jsonparser.GetString([]byte(string(resp.Body())), "token_type")))
	// token, _ = jsonparser.GetString([]byte(string(resp.Body())), "access_token")
	// fmt.Println(token)

	s.Token, _ = jsonparser.GetString([]byte(string(resp.Body())), "access_token")
}

func (s *Session) RevokeRestToken() {
	client := resty.New()
	client.SetTimeout(20 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if s.Proxy.Enable == "true" {
		client.SetProxy(s.Proxy.ProxyAddress)
	}

	// POST JSON string
	// No need to set content type, if you have client level setting
	//
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("authorization", "Basic "+s.Base64Str).
		SetBody("token=" + s.Base64Str + "&token_type_hint=token").
		//Post("https://172.23.160.37/rest-gateway/rest/api/v1/auth/revocation")
		Post("https://" + s.IpAdressNspOs + nspgoconstants.GLBL_NSP_SESSION_URL_TOKEN_REVOKE)

	log.Info("revoke token is success: ", resp.IsSuccess())

	if err != nil {
		log.Error("token revoke not succesful", err)
		return
	}
}
