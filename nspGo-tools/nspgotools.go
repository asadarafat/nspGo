package nspgotools

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/noirbizarre/gonja"
	// "github.com/scrapli/scrapligo/driver/base"
	// "github.com/scrapli/scrapligo/netconf"
	// "github.com/scrapli/scrapligo/transport"
	log "github.com/sirupsen/logrus"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Tools struct {
	JinjaTemplate      string
	JinjaParameterFill string
	LogFileName        string
	JinjaOutput        string
	// NetconfDriver      *netconf.Driver
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

// log level
// // A constant exposing all logging levels
// var AllLevels = []Level{
// 	PanicLevel, 0
// 	FatalLevel, 1
// 	ErrorLevel, 2
// 	WarnLevel,  3
// 	InfoLevel,  4
// 	DebugLevel, 5
// 	TraceLevel, 6
// }

// const (
// 	// PanicLevel level, highest level of severity. Logs and then calls panic with the
// 	// message passed to Debug, Info, ...
// 	PanicLevel Level = iota
// 	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
// 	// logging level is set to Panic.
// 	FatalLevel
// 	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
// 	// Commonly used for hooks to send errors to an error tracking service.
// 	ErrorLevel
// 	// WarnLevel level. Non-critical entries that deserve eyes.
// 	WarnLevel
// 	// InfoLevel level. General operational entries about what's going on inside the
// 	// application.
// 	InfoLevel
// 	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
// 	DebugLevel
// 	// TraceLevel level. Designates finer-grained informational events than the Debug.
// 	TraceLevel
// )

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

	// log.SetFormatter(&nested.Formatter{})

	log.SetFormatter(&log.TextFormatter{
		DisableQuote:  true,
		DisableColors: false,
		FullTimestamp: true})

	// log.SetFormatter(&log.JSONFormatter{})
}

// follow this guide
// https://github.com/scrapli/scrapligo/blob/v0.1.2/examples/netconf/main.go

func (tool *Tools) NetconfClientGet(neId, username string, password string, filter string) {

	// d, _ := netconf.NewNetconfDriver(
	// 	neId,
	// 	// base.WithPort(21830),
	// 	base.WithAuthStrictKey(false),
	// 	base.WithAuthUsername(username),
	// 	base.WithAuthPassword(password),
	// 	base.WithTransportType(transport.StandardTransportName),
	// )

	// err := d.Open()
	// if err != nil {
	// 	fmt.Printf("failed to open driver; error: %+v\n", err)
	// 	return
	// }
	// defer d.Close()

	// // filter := `
	// // <state xmlns="urn:nokia.com:sros:ns:yang:sr:state">
	// // <system><version><version-number/></version></system>
	// // </state>`

	// // filter := `
	// // <state xmlns="urn:nokia.com:sros:ns:yang:sr:state">
	// // </state>`

	// g, err := d.Get(netconf.WithNetconfFilter(filter))
	// if err != nil {
	// 	fmt.Printf("failed to get with filter; error: %+v\n", err)
	// 	return
	// }
	// log.Debug("NetcConf Edit-Config  "+neId+" Debug: ", string(g.ChannelInput))
	// log.Info("NetcConf Edit-Config  "+neId+" Response: ", g.Result)
}

func (tool *Tools) NetconfClientEditConfig(neId, username string, password string, payload string) {
	// d, _ := netconf.NewNetconfDriver(
	// 	neId,
	// 	// base.WithPort(21830),
	// 	base.WithAuthStrictKey(false),
	// 	base.WithAuthUsername(username),
	// 	base.WithAuthPassword(password),
	// 	base.WithTransportType(transport.StandardTransportName),
	// )

	// err := d.Open()
	// if err != nil {
	// 	fmt.Printf("failed to open driver; error: %+v\n", err)
	// 	return
	// }
	// defer d.Close()
	// // edit := `
	// // <config>
	// // <configure xmlns="urn:nokia.com:sros:ns:yang:sr:conf">
	// // 	<service>
	// // 		<sdp>
	// // 			<sdp-id>1</sdp-id>
	// // 			<admin-state>enable</admin-state>
	// // 			<delivery-type>mpls</delivery-type>
	// // 			<ldp>true</ldp>
	// // 			<far-end>
	// // 				<ip-address>99.99.99.1</ip-address>
	// // 			</far-end>
	// // 		</sdp>
	// // 		<vpls>
	// // 			<service-name>service-11</service-name>
	// // 			<description>This Is PW-Labels-01 PlaceHolder-TiMOS-B-21.10.R1</description>
	// // 			<service-id>101</service-id>
	// // 			<customer>1</customer>
	// // 			<spoke-sdp>
	// // 				<sdp-bind-id>1:11</sdp-bind-id>
	// // 				<description>This Is MPLS-Params-01 PlaceHolder-1-10.2.31.2</description>
	// // 			</spoke-sdp>
	// // 			<sap>
	// // 				<sap-id>1/1/c1/1:1</sap-id>
	// // 				<description>This Is MPLS-Params-02 PlaceHolder-1-7750 SR-1</description>
	// // 			</sap>
	// // 		</vpls>
	// // 	</service>
	// // </configure>
	// // </config>
	// // `
	// // r, err := d.EditConfig("candidate", edit)

	// e, EditConfigErr := d.EditConfig("candidate", payload)
	// if EditConfigErr != nil {
	// 	fmt.Printf("failed to edit config; error: %+v\n", err)
	// 	return
	// }
	// log.Debug("NetcConf Edit-Config "+neId+" Debug: ", string(e.ChannelInput))
	// log.Info("NetcConf Edit-Config "+neId+" Response: ", e.Result)
}

func (tool *Tools) NetconfClientEditCommit(neId, username string, password string) {
	// d, _ := netconf.NewNetconfDriver(
	// 	neId,
	// 	// base.WithPort(21830),
	// 	base.WithAuthStrictKey(false),
	// 	base.WithAuthUsername(username),
	// 	base.WithAuthPassword(password),
	// 	base.WithTransportType(transport.StandardTransportName),
	// )

	// err := d.Open()
	// if err != nil {
	// 	fmt.Printf("failed to open driver; error: %+v\n", err)
	// 	return
	// }

	// defer d.Close()

	// c, CommitErr := d.Commit()
	// if CommitErr != nil {
	// 	fmt.Printf("failed to edit config; error: %+v\n", err)
	// 	return
	// }
	// log.Debug("NetcConf Commit "+neId+" Debug: ", string(c.ChannelInput))
	// log.Info("NetcConf Commit "+neId+" Response: ", c.Result)

}

func (tool *Tools) WriteDataToFile(data []byte, filename string) {
	filePath, _ := (os.Getwd())
	log.Debugf("filePath: ", filePath+"/nspGo-topoViewer/html-template/static/js/"+filename)
	file, err := os.Create(filePath + "/nspGo-topoViewer/html-template/static/js/" + filename)
	if err != nil {
		return
	}
	defer file.Close()
	file.Write(data)
}

func (tool *Tools) DownloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
