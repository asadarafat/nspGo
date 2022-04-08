package main

import (
	"fmt"

	nspgotools "local.com/nspgo/nspGo-tools"
)

func main() {
	// init class
	//
	t := nspgotools.Tools{}

	template := (`
		{
			"ietf-yang-patch:yang-patch": {
				"patch-id": "as",
				"edit": [{
					"edit-id": "svc-01",
					"operation": "merge",
					"target": "nokia-conf:configure",
					"value": {
						"nokia-conf:configure": {
							"port": [{
									"port-id": "1/1/c1/1",
									"ethernet": [{
										"mode": "hybrid"
									}]
								},
								{
									"port-id": "1/1/c2/1",
									"ethernet": [{
										"mode": "hybrid"
									}]
								}
							],
							"service": {
								"sdp": [{
									"sdp-id": "1",
									"admin-state": "enable",
									"delivery-type": "mpls",
									"ldp": "true",
									"far-end": {
										"ip-address": "99.99.99.1"
									}
								}],
								"vpls": [{% for n in range(1,4000) %}{
									"service-name": "service{{n}}",
									"service-id": "n",
									"description": "<% mplsParam01 %>",
									"customer": "1",
									"spoke-sdp": [{
										"sdp-bind-id": "1:{{n}}",
										"description": "This Is MPLS-Params-01 PlaceHolder-{{n}}-<% mplsParam02 %>"
									}],
									"sap": [{
										"sap-id": "1/1/c2/1:{{n}}",
										"description": "This Is MPLS-Params-02 PlaceHolder-{{n}}-<% pwLabels %>"
									}]
								}{%- if not loop.last -%} , {% endif %}
								{% endfor %}]
							}
						}
					}
				}]
			}
		}
	`)

	// template := (`
	// <config>
	// <configure xmlns="urn:nokia.com:sros:ns:yang:sr:conf">
	// 	<service>
	// 		<sdp>
	// 			<sdp-id>3</sdp-id>
	// 			<admin-state>enable</admin-state>
	// 			<delivery-type>mpls</delivery-type>
	// 			<ldp>true</ldp>
	// 			<far-end>
	// 				<ip-address>99.99.99.1</ip-address>
	// 			</far-end>
	// 		</sdp>
	// 		<vpls>
	// 			<service-name>service-12</service-name>
	// 			<description>This Is PW-Labels-01 PlaceHolder-TiMOS-B-21.10.R1</description>
	// 			<service-id>12</service-id>
	// 			<customer>1</customer>
	// 			<spoke-sdp>
	// 				<sdp-bind-id>1:12</sdp-bind-id>
	// 				<description>This Is MPLS-Params-01 PlaceHolder-1-10.2.31.2</description>
	// 			</spoke-sdp>
	// 			<sap>
	// 				<sap-id>1/1/c1/1:12</sap-id>
	// 				<description>This Is MPLS-Params-02 PlaceHolder-1-7750 SR-1</description>
	// 			</sap>
	// 		</vpls>
	// 	</service>
	// </configure>
	// </config>
	// `)

	// variableJinja := ("nama")
	// variableJinajaIput := ("asad")

	// var X string
	t.LoadTemplateJinja(template)

	// getFilter := `
	// <state xmlns="urn:nokia.com:sros:ns:yang:sr:state">
	// <system><version><version-number/></version></system>
	// </state>`

	// // t.NetconfClient("10.2.31.2", "admin", "admin", "get", getFilter)
	// t.NetconfClientEditConfig("10.2.31.2", "admin", "admin", t.JinjaOutput)
	// t.NetconfClientEditCommit("10.2.31.2", "admin", "admin")
	// log.Info(t.JinjaOutput)

	fmt.Println(t.JinjaOutput)
}
