package nspgotools

import (
	"fmt"
	"io"
	"os"

	"github.com/noirbizarre/gonja"
	"github.com/scrapli/scrapligo/driver/base"
	"github.com/scrapli/scrapligo/netconf"
	"github.com/scrapli/scrapligo/transport"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Tools struct {
	JinjaTemplate      string
	JinjaParameterFill string
	LogFileName        string
	JinjaOutput        string
	NetconfDriver      *netconf.Driver
}

func (tool *Tools) LoadTemplateJinja(template string) {
	tpl, err := gonja.FromString(template)

	if err != nil {
		panic(err)
	}
	// Now you can render the template with the given
	// pongo2.Context how often you want to.
	tool.JinjaOutput, err = tpl.Execute(gonja.Context{"variable": "variableInput", "a": "a"})
	if err != nil {
		panic(err)
	}
}

func (tool *Tools) InitLogger(filePath string, level uint32) {
	mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		// mw := io.MultiWriter(&lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
	log.SetLevel(log.Level(level))
	log.SetOutput(mw)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true})
}

// follow this guide
// https://github.com/scrapli/scrapligo/blob/v0.1.2/examples/netconf/main.go

func (tool *Tools) NetconfClientGet(neId, username string, password string, filter string) {

	d, _ := netconf.NewNetconfDriver(
		neId,
		// base.WithPort(21830),
		base.WithAuthStrictKey(false),
		base.WithAuthUsername(username),
		base.WithAuthPassword(password),
		base.WithTransportType(transport.StandardTransportName),
	)

	err := d.Open()
	if err != nil {
		fmt.Printf("failed to open driver; error: %+v\n", err)
		return
	}
	defer d.Close()

	// filter := `
	// <state xmlns="urn:nokia.com:sros:ns:yang:sr:state">
	// <system><version><version-number/></version></system>
	// </state>`

	// filter := `
	// <state xmlns="urn:nokia.com:sros:ns:yang:sr:state">
	// </state>`

	g, err := d.Get(netconf.WithNetconfFilter(filter))
	if err != nil {
		fmt.Printf("failed to get with filter; error: %+v\n", err)
		return
	}
	log.Debug("NetcConf Edit-Config  "+neId+" Debug: ", string(g.ChannelInput))
	log.Info("NetcConf Edit-Config  "+neId+" Response: ", g.Result)
}

func (tool *Tools) NetconfClientEditConfig(neId, username string, password string, payload string) {
	d, _ := netconf.NewNetconfDriver(
		neId,
		// base.WithPort(21830),
		base.WithAuthStrictKey(false),
		base.WithAuthUsername(username),
		base.WithAuthPassword(password),
		base.WithTransportType(transport.StandardTransportName),
	)

	err := d.Open()
	if err != nil {
		fmt.Printf("failed to open driver; error: %+v\n", err)
		return
	}
	defer d.Close()
	// edit := `
	// <config>
	// <configure xmlns="urn:nokia.com:sros:ns:yang:sr:conf">
	// 	<service>
	// 		<sdp>
	// 			<sdp-id>1</sdp-id>
	// 			<admin-state>enable</admin-state>
	// 			<delivery-type>mpls</delivery-type>
	// 			<ldp>true</ldp>
	// 			<far-end>
	// 				<ip-address>99.99.99.1</ip-address>
	// 			</far-end>
	// 		</sdp>
	// 		<vpls>
	// 			<service-name>service-11</service-name>
	// 			<description>This Is PW-Labels-01 PlaceHolder-TiMOS-B-21.10.R1</description>
	// 			<service-id>101</service-id>
	// 			<customer>1</customer>
	// 			<spoke-sdp>
	// 				<sdp-bind-id>1:11</sdp-bind-id>
	// 				<description>This Is MPLS-Params-01 PlaceHolder-1-10.2.31.2</description>
	// 			</spoke-sdp>
	// 			<sap>
	// 				<sap-id>1/1/c1/1:1</sap-id>
	// 				<description>This Is MPLS-Params-02 PlaceHolder-1-7750 SR-1</description>
	// 			</sap>
	// 		</vpls>
	// 	</service>
	// </configure>
	// </config>
	// `
	// r, err := d.EditConfig("candidate", edit)

	e, EditConfigErr := d.EditConfig("candidate", payload)
	if EditConfigErr != nil {
		fmt.Printf("failed to edit config; error: %+v\n", err)
		return
	}
	log.Debug("NetcConf Edit-Config "+neId+" Debug: ", string(e.ChannelInput))
	log.Info("NetcConf Edit-Config "+neId+" Response: ", e.Result)
}

func (tool *Tools) NetconfClientEditCommit(neId, username string, password string) {
	d, _ := netconf.NewNetconfDriver(
		neId,
		// base.WithPort(21830),
		base.WithAuthStrictKey(false),
		base.WithAuthUsername(username),
		base.WithAuthPassword(password),
		base.WithTransportType(transport.StandardTransportName),
	)

	err := d.Open()
	if err != nil {
		fmt.Printf("failed to open driver; error: %+v\n", err)
		return
	}

	defer d.Close()

	c, CommitErr := d.Commit()
	if CommitErr != nil {
		fmt.Printf("failed to edit config; error: %+v\n", err)
		return
	}
	log.Debug("NetcConf Commit "+neId+" Debug: ", string(c.ChannelInput))
	log.Info("NetcConf Commit "+neId+" Response: ", c.Result)

}
