package main

import (
	log "github.com/sirupsen/logrus"
	nspgosession "local.com/nspgo/nspGo-session"
	nspgowfm "local.com/nspgo/nspGo-wfm"
)

func main() {

	////  for below sourceCode auto generated
	nspgowfm.GenerateWfmGoCode()

	// init class session
	p := nspgosession.Session{}
	p.LoadConfig()

	log.Info("nsp.nspOsIP :", p.IpAdressNspOs)
	log.Info("nsp.nspIprcIP :", p.IpAdressIprc)
	log.Info("nsp.Username :", p.Username)
	log.Info("nsp.Password :", p.Password)
	log.Info("nsp.linetoken :", p.Token)

	p.EncodeUserName()
	log.Info(p.EncodeUserName())

	p.GetRestToken()
	log.Info("nsp.linetoken_NEW :", p.Token)

	//// execute auto generated code
	// w := nspgowfm.Wfm{}

	// payload := " "
	// w.WfmV1WorkflowGet(p.IpAdressNspOs, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress, []byte(payload))
	//asad
}
