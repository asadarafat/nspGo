package main

import (
	"fmt"

	nspgoipoptim "local.com/nspgo/nspGo-ipOptim"
	nspgosession "local.com/nspgo/nspGo-session"
)

func main() {
	// init class session
	//
	p := nspgosession.Session{}
	p.LoadConfig()

	fmt.Println("nsp.nspOsIP :", p.IpAdressNspOs)
	fmt.Println("nsp.nspIprcIP :", p.IpAdressIprc)
	fmt.Println("nsp.Username :", p.Username)
	fmt.Println("nsp.Password :", p.Password)
	fmt.Println("nsp.linetoken :", p.Token)

	p.EncodeUserName()
	fmt.Println(p.EncodeUserName())

	p.GetRestToken()
	fmt.Println("nsp.linetoken_NEW :", p.Token)

	//// init class ip optim
	// ip := nspgoipoptim.IpOptim{}

	//// get network ietf
	// will throw the response to IpOptim.ResponseData and func returen
	// ip.GetNetworkIetf(p.IpAdressIprc, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress)

	// // print the Response Data
	// fmt.Println(ip.ResponseData)

	// print the function return
	// fmt.Println(ip.GetNetworkIetf(p.IpAdressIprc, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress))

	// fmt.Println(ip.GetL3Networks(p.IpAdressIprc, p.Token, p.Proxy.Enable, p.Proxy.ProxyAddress))
	nspgoipoptim.GenerateIpOptimGoCode()
	p.RevokeRestToken()

}
