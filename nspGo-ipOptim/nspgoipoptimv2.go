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
	toolLogger.InitLogger("./logs/nspGo-ipOptim.log", 3)
}

// Find an object by Application ID
func (ipO *IpOptim) IpoV4GenericApplicationidIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/generic/application-id/{id}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Set an object Application ID
func (ipO *IpOptim) IpoV4GenericApplicationidIdAppIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/generic/application-id/{id}/{appId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find all objects being consumed by an object
func (ipO *IpOptim) IpoV4GenericConsumedUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/generic/consumed/{uuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find the unique identifier of an object given an external identifier
func (ipO *IpOptim) IpoV4GenericFindbyexternalidPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/generic/find-by-external-id")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find all tenants assigned to an object
func (ipO *IpOptim) IpoV4GenericTenantsUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/generic/tenants/{uuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an object by a unique identifier
func (ipO *IpOptim) IpoV4GenericUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/generic/{uuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Add an endpoint to a L2 Backhaul Service
func (ipO *IpOptim) IpoV4IetfEthtsvcL2backhaulServiceUuidEndpointPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ietf/ethtsvc/l2-backhaul/{serviceUuid}/endpoint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update endpoint of a L2 Backhaul service
func (ipO *IpOptim) IpoV4IetfEthtsvcL2backhaulServiceUuidEndpointEndpointUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ietf/ethtsvc/l2-backhaul/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch endpoint of an L3 VPN service
func (ipO *IpOptim) IpoV4IetfEthtsvcL2backhaulServiceUuidEndpointEndpointUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ietf/ethtsvc/l2-backhaul/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create L2 Backhaul Service
func (ipO *IpOptim) IpoV4IetfEthtsvcL2backhaulAutoPathPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ietf/ethtsvc/l2backhaul/{autoPath}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify L2 Backhaul service
func (ipO *IpOptim) IpoV4IetfEthtsvcL2backhaulServiceUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ietf/ethtsvc/l2backhaul/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch L2 Backhaul service
func (ipO *IpOptim) IpoV4IetfEthtsvcL2backhaulServiceUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ietf/ethtsvc/l2backhaul/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get TE link. A link has a UUID, attributes, a source (node, termination-point) and a destination (node, termination-point).
func (ipO *IpOptim) IpoV4IetfTeLinkLinkIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ietf/te/link/{linkId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch TE link by updating the configuration parameters. Supported configuration parameter(s): SRLG, Latency, AdminStatus
func (ipO *IpOptim) IpoV4IetfTeLinkLinkUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ietf/te/link/{linkUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get TE network. A network has a UUID, attributes, a list of nodes and a list of links.
func (ipO *IpOptim) IpoV4IetfTeNetworkNetworkIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ietf/te/network/{networkId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get TE networks. A list of networks where each network has a UUID.
func (ipO *IpOptim) IpoV4IetfTeNetworksGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ietf/te/networks")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get TE node. A node has a UUID, attributes and a list of termination-points.
func (ipO *IpOptim) IpoV4IetfTeNodeNodeIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ietf/te/node/{nodeId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get TE termination-point. A termination-point has a UUID and attributes.
func (ipO *IpOptim) IpoV4IetfTeTerminationpointTpIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ietf/te/termination-point/{tpId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get L2 Path
func (ipO *IpOptim) IpoV4L2BackhaulPathsL2pathsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/l2/backhaul/paths/l2-paths")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Fetch the template global names for all AMIs, grouped by AMI
func (ipO *IpOptim) IpoV4MediationAmiversiontemplatesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mediation/ami-version-templates")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Fetch the AMIs and the versions.
func (ipO *IpOptim) IpoV4MediationAmisversionsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mediation/amis-versions")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all Augmentation Meta Infos
func (ipO *IpOptim) IpoV4MediationAugmentationmetaGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mediation/augmentation-meta")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an Augmentation Meta
func (ipO *IpOptim) IpoV4MediationAugmentationmetaPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mediation/augmentation-meta")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an Augmentation Meta
func (ipO *IpOptim) IpoV4MediationAugmentationmetaIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mediation/augmentation-meta/{id}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an Augmentation meta
func (ipO *IpOptim) IpoV4MediationAugmentationmetaIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mediation/augmentation-meta/{id}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// This action is to fetch the meta information of the passthrough attributes that could be provided during a template invocation. The returned information would be a JSON describing the attributes, their types, and optional constraint information
func (ipO *IpOptim) IpoV4MediationMediationaugmentationPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mediation/mediation-augmentation")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Fetch RSVP or SR-TE LSP list which discovered by SDN through device management NFMP or MDM
func (ipO *IpOptim) IpoV4MplsLsplistPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp-list")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Asynchronous API to create MPLS LSP path (PCE Initiated only)
func (ipO *IpOptim) IpoV4MplsLsppathPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp-path")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Fetch MPLS LSP path list
func (ipO *IpOptim) IpoV4MplsLsppathlistPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp-path-list")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Override path profile for both PCE and PCC initiated LSPs, note that the path profile will not be overridden if a path ID doesn't exist or the corresponding LSP configuration is invalid
func (ipO *IpOptim) IpoV4MplsLsppathprofileoverridePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp-path-profile-override")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get MPLS LSP path
func (ipO *IpOptim) IpoV4MplsLsppathPathIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp-path/{pathId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Asynchronous API to delete MPLS LSP path(PCE Initiated only)
func (ipO *IpOptim) IpoV4MplsLsppathPathIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp-path/{pathId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Asynchronous API to patch PCE Initiated MPLS LSP path config and add/remove path profile override for a MPLS LSP path. Setting pathProfileId.profileId = -1 deletes overide and sending same pathProfileOverride will cause a reset if the profileOverrideState is FAILED.
func (ipO *IpOptim) IpoV4MplsLsppathPathIdPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp-path/{pathId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get MPLS LSP paths
func (ipO *IpOptim) IpoV4MplsLsppathsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp-paths")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Asynchronous API to create MPLS LSP paths(PCE Initiated only)
func (ipO *IpOptim) IpoV4MplsLsppathsPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp-paths")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Asynchronous API to delete MPLS LSP paths(PCE Initiated only). List maximum size is 500.
func (ipO *IpOptim) IpoV4MplsLsppathsDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp-paths")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get MPLS LSP paths
func (ipO *IpOptim) IpoV4MplsLsppathspaginatedLimitIndexGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp-paths-paginated/{limit}/{index}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get MPLS LSP paths on link
func (ipO *IpOptim) IpoV4MplsLspPathsOnlinkLinkIdRequestTypeGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp/paths/on-link/{linkId}/{requestType}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get MPLS LSP paths assigned with the provided profile and group ID pair
func (ipO *IpOptim) IpoV4MplsLspPathsProfilegroupProfileIdExtendedIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp/paths/profile-group/{profileId}/{extendedId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get RSVP or SR-TE LSP which discovered by SDN through device management NFMP or MDM
func (ipO *IpOptim) IpoV4MplsLspLspIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsp/{lspId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get RSVP or SR-TE LSPs which discovered by SDN through device management NFMP or MDM
func (ipO *IpOptim) IpoV4MplsLspsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/lsps")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Trigger global concurrent optimization on PCEP LSP paths
func (ipO *IpOptim) IpoV4MplsOptimizationPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/optimization")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Trigger resignal on a list of PCEP LSP paths
func (ipO *IpOptim) IpoV4MplsResignalPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/mpls/resignal")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all Network Elements
func (ipO *IpOptim) IpoV4NeGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ne")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query a unique identifier for a Network Element by System ID
func (ipO *IpOptim) IpoV4NeSystemSystemIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ne/system/{systemId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update Network Element capabilities
func (ipO *IpOptim) IpoV4NeIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ne/{id}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a network element by a unique identifier
func (ipO *IpOptim) IpoV4NeUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ne/{uuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get NSP latency parameters
func (ipO *IpOptim) IpoV4NspConfigurationLatencyGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/configuration/latency")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch NSP latency parameters
func (ipO *IpOptim) IpoV4NspConfigurationLatencyPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/configuration/latency")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// API used to delete twamp tests.
func (ipO *IpOptim) IpoV4NspConfigurationLatencySessionDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/configuration/latency/{session}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get NSP NRCP historical parameters.
func (ipO *IpOptim) IpoV4NspConfigurationNrcphistoricalGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/configuration/nrcp-historical")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch NSP NRCP historical parameters.
func (ipO *IpOptim) IpoV4NspConfigurationNrcphistoricalPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/configuration/nrcp-historical")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get SR policy configuration parameters.
func (ipO *IpOptim) IpoV4NspConfigurationSrpolicyconfigGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/configuration/sr-policy-config")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch SR policy configuration parameters.
func (ipO *IpOptim) IpoV4NspConfigurationSrpolicyconfigPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/configuration/sr-policy-config")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get TCA Policy configuration. Policies configuration include list of policies and collection enable flag
func (ipO *IpOptim) IpoV4NspConfigurationTcaconfigpolicyGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/configuration/tca-config-policy")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch a TCA configuration policy by updating the parameters
func (ipO *IpOptim) IpoV4NspConfigurationTcaconfigpolicyPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/configuration/tca-config-policy")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get NSP traffic collection parameters.
func (ipO *IpOptim) IpoV4NspConfigurationTrafficdatacollectionGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/configuration/traffic-data-collection")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch NSP traffic-collection parameters.
func (ipO *IpOptim) IpoV4NspConfigurationTrafficdatacollectionPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/configuration/traffic-data-collection")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get link connected to the given termination point. The link may be a source or destination on the termination point. A link has a UUID, attributes, a source (node, termination-point) and a destination (node, termination-point)
func (ipO *IpOptim) IpoV4NspNetL3LinkTpTpIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/link/tp/{tpId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get link. A link has a UUID, attributes, a source (node, termination-point) and a destination (node, termination-point)
func (ipO *IpOptim) IpoV4NspNetL3LinkLinkIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/link/{linkId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch a link by updating the configuration parameters. Configuration parameters supported are srlg-value, latency, te-metric, igp-metric, administrative-group, admin-status, measuredIpBw and measuredMplsBw
func (ipO *IpOptim) IpoV4NspNetL3LinkLinkUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/link/{linkUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Fetch link list, note that invalid input link uuids will be ignored
func (ipO *IpOptim) IpoV4NspNetL3LinksPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/links")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get NSP network. A network has a UUID, attributes, a list of nodes and a list of links
func (ipO *IpOptim) IpoV4NspNetL3NetworkNetworkIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/network/{networkId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get NSP networks. A list of networks where each network has a UUID
func (ipO *IpOptim) IpoV4NspNetL3NetworksGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/networks")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get NSP node. A node has a UUID, attributes and a list of termination-points
func (ipO *IpOptim) IpoV4NspNetL3NodeNodeIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/node/{nodeId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch a NSP node. Currently, only admin-status parameter is supported.
func (ipO *IpOptim) IpoV4NspNetL3NodeNodeIdPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/node/{nodeId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get NSP nodes by IGP Router ID. A node has a UUID, attributes and a list of termination-points
func (ipO *IpOptim) IpoV4NspNetL3NodesRouterRouterIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/nodes/router/{routerId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get NSP nodes by Site IP. A node has a UUID, attributes and a list of termination-points
func (ipO *IpOptim) IpoV4NspNetL3NodesSiteSiteIpGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/nodes/site/{siteIp}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get NSP termination-point. A termination-point has a UUID and attributes
func (ipO *IpOptim) IpoV4NspNetL3TerminationpointTpIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/termination-point/{tpId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get NSP reverse termination-point for a given UUID. If the UUID corresponds to a termination point on a point-to-point link, the returned termination point will be the termination point on the connected remote router. If the UUID corresponds to a termination point on a broadcast link, the returned termination point will be the reverse termination point connected to the source router
func (ipO *IpOptim) IpoV4NspNetL3TerminationpointTpIdReverseGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/termination-point/{tpId}/reverse")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get NSP termination-point for a given node ID and IP Address
func (ipO *IpOptim) IpoV4NspNetL3TerminationpointsNodeIdIpAddressGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/termination-points/{nodeId}/{ipAddress}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Clean up topology references of all IP domains, or a specific IP domain (set networkId to -1 for all IP domains)
func (ipO *IpOptim) IpoV4NspNetL3TopologyreferencesDomainNetworkIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/topology-references/domain/{networkId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Clean up topology references for a specific link identifier
func (ipO *IpOptim) IpoV4NspNetL3TopologyreferencesLinkLinkIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/net/l3/topology-references/link/{linkId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// LSP Diagnosis API to run additional calculations and return a result set that may aid in determining why an LSP is operationally down. Use cases that involved grouping or association such as Diversity or Bidirectionality may return limited results. Please note that this API is currently experimental and not yet comprehensive.
func (ipO *IpOptim) IpoV4NspPathtoolDiagnoselspLspPathIdPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/pathtool/diagnose-lsp/{lspPathId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Path Diagnosis API to invoke extra behaviour to try and determine why/what paths are found or not found. Please note that this API is currently experimental and not yet comprehensive.
func (ipO *IpOptim) IpoV4NspPathtoolDiagnosepathPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/pathtool/diagnose-path")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Path Finder API to validate whether a path is possible without actually creating the path
func (ipO *IpOptim) IpoV4NspPathtoolFindpathPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/nsp/pathtool/find-path")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create flows
func (ipO *IpOptim) IpoV4OpenflowFlowsPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/openflow/flows")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete flows
func (ipO *IpOptim) IpoV4OpenflowFlowsDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/openflow/flows")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update flow instruction
func (ipO *IpOptim) IpoV4OpenflowFlowsPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/openflow/flows")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query flows based on search criteria
func (ipO *IpOptim) IpoV4OpenflowFlowsSearchPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/openflow/flows/search")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query flows based on unique identifier
func (ipO *IpOptim) IpoV4OpenflowFlowsSearchbyidPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/openflow/flows/search-by-id")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all ports of an openflow switch
func (ipO *IpOptim) IpoV4OpenflowPortsDatapathIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/openflow/ports/{datapathId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all openflow switches
func (ipO *IpOptim) IpoV4OpenflowSwitchesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/openflow/switches")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all openflow switches in a router
func (ipO *IpOptim) IpoV4OpenflowSwitchesNeIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/openflow/switches/{neId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all tables in an openflow switch
func (ipO *IpOptim) IpoV4OpenflowTablesDatapathIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/openflow/tables/{datapathId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all physical links
func (ipO *IpOptim) IpoV4PhysicallinksGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/physicallinks")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Add a port to a multipoint link
func (ipO *IpOptim) IpoV4PhysicallinksMultipointportMultipointLinkIdPortIdPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/physicallinks/multipoint-port/{multipointLinkId}/{portId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Remove a port to a multipoint link
func (ipO *IpOptim) IpoV4PhysicallinksMultipointportMultipointLinkIdPortIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/physicallinks/multipoint-port/{multipointLinkId}/{portId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete a multipoint link that was defined in NSP
func (ipO *IpOptim) IpoV4PhysicallinksMultipointLinkIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/physicallinks/multipoint/{linkId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Define a multipoint link in NSP between multiple ports
func (ipO *IpOptim) IpoV4PhysicallinksMultipointNamePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/physicallinks/multipoint/{name}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a multipoint link by a unique identifier
func (ipO *IpOptim) IpoV4PhysicallinksMultipointUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/physicallinks/multipoint/{uuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all multipoint links
func (ipO *IpOptim) IpoV4PhysicallinksMultipointsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/physicallinks/multipoints")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an undiscovered physical link that was defined in NSP
func (ipO *IpOptim) IpoV4PhysicallinksLinkIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/physicallinks/{linkId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Define an undiscovered physical link in NSP
func (ipO *IpOptim) IpoV4PhysicallinksSrcIdDestIdPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/physicallinks/{srcId}/{destId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a physical link by a unique identifier
func (ipO *IpOptim) IpoV4PhysicallinksUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/physicallinks/{uuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all connection profile objects
func (ipO *IpOptim) IpoV4PolicyAllconnectionprofileNeIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/all-connection-profile/{neId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all IP Optical correlation policies
func (ipO *IpOptim) IpoV4PolicyIpopticalcorrelationpoliciyGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/ip-optical-correlation-policiy")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an IP Optical correlation policy
func (ipO *IpOptim) IpoV4PolicyIpopticalcorrelationpolicyPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/ip-optical-correlation-policy")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query an IP Optical correlation policy
func (ipO *IpOptim) IpoV4PolicyIpopticalcorrelationpolicyPolicyIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/ip-optical-correlation-policy/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update an IP Optical correlation policy
func (ipO *IpOptim) IpoV4PolicyIpopticalcorrelationpolicyPolicyIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/ip-optical-correlation-policy/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an IP Optical correlation policy
func (ipO *IpOptim) IpoV4PolicyIpopticalcorrelationpolicyPolicyIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/ip-optical-correlation-policy/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all L3VPN RT/RD policies
func (ipO *IpOptim) IpoV4PolicyRdrtrangesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/rd-rt-ranges")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a L3VPN RT/RD policy
func (ipO *IpOptim) IpoV4PolicyRdrtrangesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/rd-rt-ranges")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update a L3VPN RT/RD policy
func (ipO *IpOptim) IpoV4PolicyRdrtrangesPolicyIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/rd-rt-ranges/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an existing L3VPN RT/RD policy.
func (ipO *IpOptim) IpoV4PolicyRdrtrangesPolicyIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/rd-rt-ranges/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all router port protection group policies
func (ipO *IpOptim) IpoV4PolicyRouterportprotectiongrouppolicyGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/router-port-protection-group-policy")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a router port protection group policy
func (ipO *IpOptim) IpoV4PolicyRouterportprotectiongrouppolicyPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/router-port-protection-group-policy")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query a router port protection group policy
func (ipO *IpOptim) IpoV4PolicyRouterportprotectiongrouppolicyPolicyIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/router-port-protection-group-policy/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update a router port protection group policy
func (ipO *IpOptim) IpoV4PolicyRouterportprotectiongrouppolicyPolicyIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/router-port-protection-group-policy/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete a router port protection group policy
func (ipO *IpOptim) IpoV4PolicyRouterportprotectiongrouppolicyPolicyIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/router-port-protection-group-policy/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a Steering Parameter
func (ipO *IpOptim) IpoV4PolicySteeringparameterPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/steering-parameter")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete a Steering Parameter
func (ipO *IpOptim) IpoV4PolicySteeringparameterSteeringParameterNameDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/steering-parameter/{steeringParameterName}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all Steering Paramters
func (ipO *IpOptim) IpoV4PolicySteeringparametersGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/steering-parameters")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all tunnel selection policies.
func (ipO *IpOptim) IpoV4PolicyTunnelselectionsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/tunnel-selections")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a tunnel selection policy.
func (ipO *IpOptim) IpoV4PolicyTunnelselectionsPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/tunnel-selections")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query a tunnel selections policy by its unique identifier.
func (ipO *IpOptim) IpoV4PolicyTunnelselectionsPolicyIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/tunnel-selections/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update a tunnel selection policy.
func (ipO *IpOptim) IpoV4PolicyTunnelselectionsPolicyIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/tunnel-selections/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an existing tunnel selections policy.
func (ipO *IpOptim) IpoV4PolicyTunnelselectionsPolicyIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/policy/tunnel-selections/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find all ports in the network
func (ipO *IpOptim) IpoV4PortsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ports")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4PortsNeallbytenantNeIdTenantTenantUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ports/ne-all-by-tenant/{neId}/tenant/{tenantUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all ports of the provided network element
func (ipO *IpOptim) IpoV4PortsNeallNeIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ports/ne-all/{neId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4PortsNebytenantNeIdServicetypeServiceTypeTenantTenantUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ports/ne-by-tenant/{neId}/servicetype/{serviceType}/tenant/{tenantUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all ports of the provided network element of a service type
func (ipO *IpOptim) IpoV4PortsNeNeUuidServicetypeServiceTypeGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ports/ne/{neUuid}/servicetype/{serviceType}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find number of services on a specified port
func (ipO *IpOptim) IpoV4PortsServicecountonportPortUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ports/service-count-on-port/{portUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all services using a port
func (ipO *IpOptim) IpoV4PortsServicesonportPortUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ports/services-on-port/{portUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find all ports in the network that support a given service type
func (ipO *IpOptim) IpoV4PortsServicetypeServiceTypeGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ports/servicetype/{serviceType}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4PortsServicetypeServiceTypeTenantTenantUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ports/servicetype/{serviceType}/tenant/{tenantUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4PortsTenantTenantUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ports/tenant/{tenantUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update the port
func (ipO *IpOptim) IpoV4PortsIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ports/{id}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a port by unique identifier
func (ipO *IpOptim) IpoV4PortsPortUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/ports/{portUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Retrieve information about the currently-authenticated user from the perspective of the system
func (ipO *IpOptim) IpoV4SecurityAuthenticationGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/security/authentication")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// BETA - Run service constraint test
func (ipO *IpOptim) IpoV4ServicedebugConstrainttestServiceUuidPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/servicedebug/constraint-test/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Triggers re-computation for a L2 or L3 DCI Service. The DCI Service types automatically react to topological and service requirements and perform necessary differential actions for the service however this API can be used to manually invoke that execution. The process that is normally automatically triggered to differentiate the service state with the required state will be triggered
func (ipO *IpOptim) IpoV4ServicedebugDcirecomputeServiceUuidPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/servicedebug/dci-recompute/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Trigger the Transport Tech Zone Topology computation
func (ipO *IpOptim) IpoV4ServicedebugTriggerttzalgorithmPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/servicedebug/trigger-ttz-algorithm")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all service objects
func (ipO *IpOptim) IpoV4ServicesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a C-Line service
func (ipO *IpOptim) IpoV4ServicesClinesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/clines")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify a C-Line service
func (ipO *IpOptim) IpoV4ServicesClinesServiceUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/clines/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch a C-Line service
func (ipO *IpOptim) IpoV4ServicesClinesServiceUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/clines/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update the endpoint of a C-Line service
func (ipO *IpOptim) IpoV4ServicesClinesServiceUuidEndpointEndpointUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/clines/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch the endpoint of a C-Line service
func (ipO *IpOptim) IpoV4ServicesClinesServiceUuidEndpointEndpointUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/clines/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Retrieve inventory of service endpoints, tunnels, and physical links for the requested service UUID
func (ipO *IpOptim) IpoV4ServicesComponentsUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/components/{uuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an IP E-Access service
func (ipO *IpOptim) IpoV4ServicesEaccessPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/eaccess")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an IP E-Access service
func (ipO *IpOptim) IpoV4ServicesEaccessServiceUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/eaccess/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch an IP E-Access service
func (ipO *IpOptim) IpoV4ServicesEaccessServiceUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/eaccess/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an E-LAN VPN service
func (ipO *IpOptim) IpoV4ServicesElansPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/elans")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an E-LAN service
func (ipO *IpOptim) IpoV4ServicesElansServiceUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/elans/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch an E-LAN service
func (ipO *IpOptim) IpoV4ServicesElansServiceUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/elans/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Add an endpoint to an E-LAN service
func (ipO *IpOptim) IpoV4ServicesElansServiceUuidEndpointPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/elans/{serviceUuid}/endpoint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update the endpoint of an E-LAN service
func (ipO *IpOptim) IpoV4ServicesElansServiceUuidEndpointEndpointUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/elans/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch the endpoint of an E-LAN service
func (ipO *IpOptim) IpoV4ServicesElansServiceUuidEndpointEndpointUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/elans/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an IP or optical E-Line service
func (ipO *IpOptim) IpoV4ServicesElinesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/elines")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an IP or optical E-Line service
func (ipO *IpOptim) IpoV4ServicesElinesServiceUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/elines/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch an IP or optical E-Line service
func (ipO *IpOptim) IpoV4ServicesElinesServiceUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/elines/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Add an endpoint to an E-Line service
func (ipO *IpOptim) IpoV4ServicesElinesServiceUuidEndpointPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/elines/{serviceUuid}/endpoint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update the endpoint of an E-Line service
func (ipO *IpOptim) IpoV4ServicesElinesServiceUuidEndpointEndpointUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/elines/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch the endpoint of an E-Line service
func (ipO *IpOptim) IpoV4ServicesElinesServiceUuidEndpointEndpointUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/elines/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Sites within the specified L3-VPN service that are not correctly connected by VRF membership will be moved to a new service. Currently, only sites connected by Extranet route targets will be moved to a new service.
func (ipO *IpOptim) IpoV4ServicesFixl3vpnvrfmembershipServiceUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/fix-l3-vpn-vrf-membership/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an IP IES service
func (ipO *IpOptim) IpoV4ServicesIesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/ies")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an IP IES service
func (ipO *IpOptim) IpoV4ServicesIesServiceUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/ies/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch an IP IES service
func (ipO *IpOptim) IpoV4ServicesIesServiceUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/ies/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Add an endpoint to a IES service
func (ipO *IpOptim) IpoV4ServicesIesServiceUuidEndpointPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/ies/{serviceUuid}/endpoint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update the endpoint of an IES service
func (ipO *IpOptim) IpoV4ServicesIesServiceUuidEndpointEndpointUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/ies/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch the endpoint of an IES service
func (ipO *IpOptim) IpoV4ServicesIesServiceUuidEndpointEndpointUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/ies/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Add a loop back endpoint to an IES service
func (ipO *IpOptim) IpoV4ServicesIesServiceUuidLoopbackendpointPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/ies/{serviceUuid}/loopback-endpoint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update loop back endpoint of an IES service
func (ipO *IpOptim) IpoV4ServicesIesServiceUuidLoopbackendpointEndpointUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/ies/{serviceUuid}/loopback-endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch loop back endpoint of an IES service
func (ipO *IpOptim) IpoV4ServicesIesServiceUuidLoopbackendpointEndpointUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/ies/{serviceUuid}/loopback-endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Support for limited Data Center use cases as VXLAN Stitched connectivity type. Unsupported Proof of Concept with MPLS End-to-End connectivity type. An L3 DCI Service is a Network Function Interconnect implementation providing Layer 3 Data Center connectivity for EVPN services in between data center and WAN entities. Integrates various use cases with the Nokia Nuage Solution, Segment Routing and VXLAN
func (ipO *IpOptim) IpoV4ServicesL3dcivpnsPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-dci-vpns")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Add an L3 endpoint to an L3 DCI service
func (ipO *IpOptim) IpoV4ServicesL3dcisServiceUuidEndpointPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-dcis/{serviceUuid}/endpoint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch L3 endpoint of an L3 DCI service
func (ipO *IpOptim) IpoV4ServicesL3dcisServiceUuidEndpointEndpointUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-dcis/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Add a L3 loop back endpoint to an L3 DCI service
func (ipO *IpOptim) IpoV4ServicesL3dcisServiceUuidLoopbackendpointPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-dcis/{serviceUuid}/loopback-endpoint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch loop back endpoint of an L3 DCI service
func (ipO *IpOptim) IpoV4ServicesL3dcisServiceUuidLoopbackendpointEndpointUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-dcis/{serviceUuid}/loopback-endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Runs an audit to find all L3-VPN services with invalid VRF membership. Currently the audit only considers L3-VPN services that contains sites connected by Extranet route targets as invalid.
func (ipO *IpOptim) IpoV4ServicesL3vpnrtauditGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-vpn-rt-audit")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an L3 VPN service
func (ipO *IpOptim) IpoV4ServicesL3vpnsPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-vpns")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an L3 VPN service
func (ipO *IpOptim) IpoV4ServicesL3vpnsServiceUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-vpns/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch an L3 VPN service
func (ipO *IpOptim) IpoV4ServicesL3vpnsServiceUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-vpns/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Add an endpoint to an L3 VPN service
func (ipO *IpOptim) IpoV4ServicesL3vpnsServiceUuidEndpointPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-vpns/{serviceUuid}/endpoint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update endpoint of an L3 VPN service
func (ipO *IpOptim) IpoV4ServicesL3vpnsServiceUuidEndpointEndpointUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-vpns/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch endpoint of an L3 VPN service
func (ipO *IpOptim) IpoV4ServicesL3vpnsServiceUuidEndpointEndpointUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-vpns/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Add a loop back endpoint to an L3 VPN service
func (ipO *IpOptim) IpoV4ServicesL3vpnsServiceUuidLoopbackendpointPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-vpns/{serviceUuid}/loopback-endpoint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update loop back endpoint of an L3 VPN service
func (ipO *IpOptim) IpoV4ServicesL3vpnsServiceUuidLoopbackendpointEndpointUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-vpns/{serviceUuid}/loopback-endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch loop back endpoint of an L3 VPN service
func (ipO *IpOptim) IpoV4ServicesL3vpnsServiceUuidLoopbackendpointEndpointUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/l3-vpns/{serviceUuid}/loopback-endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a LAG service
func (ipO *IpOptim) IpoV4ServicesLagsPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/lags")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an LAG service
func (ipO *IpOptim) IpoV4ServicesLagsServiceUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/lags/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a normalized L3 VPN service
func (ipO *IpOptim) IpoV4ServicesNormalizedl3vpnsPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/normalized-l3-vpns")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify a normalized L3 VPN service
func (ipO *IpOptim) IpoV4ServicesNormalizedl3vpnsServiceUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/normalized-l3-vpns/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch a normalized L3 VPN service
func (ipO *IpOptim) IpoV4ServicesNormalizedl3vpnsServiceUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/normalized-l3-vpns/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Add an endpoint to a normalized L3 VPN service
func (ipO *IpOptim) IpoV4ServicesNormalizedl3vpnsServiceUuidEndpointPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/normalized-l3-vpns/{serviceUuid}/endpoint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create or modify an optical service
func (ipO *IpOptim) IpoV4ServicesOpticalPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/optical")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an optical service
func (ipO *IpOptim) IpoV4ServicesOpticalServiceUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/optical/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch an optical service
func (ipO *IpOptim) IpoV4ServicesOpticalServiceUuidPatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/optical/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query physical links used by service
func (ipO *IpOptim) IpoV4ServicesPhysicallinksonserviceServiceUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/physical-links-on-service/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4ServicesResourcecountTenantUuidResourceTypeGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/resource-count/{tenantUuid}/{resourceType}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query services using a tunnel
func (ipO *IpOptim) IpoV4ServicesServicesontunnelTunnelUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/services-on-tunnel/{tunnelUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Stitch a L2 Extension Service to a ELine or ELAN Service
func (ipO *IpOptim) IpoV4ServicesStitchl2extServiceUuidL2extServiceUuidServiceEndpointUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/stitch-l2-ext/{serviceUuid}/{l2extServiceUuid}/{serviceEndpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4ServicesTenantUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/tenant/{uuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a Tunnel
func (ipO *IpOptim) IpoV4ServicesTunnelTunnelUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/tunnel/{tunnelUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify a tunnel
func (ipO *IpOptim) IpoV4ServicesTunnelTunnelUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/tunnel/{tunnelUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all Tunnels
func (ipO *IpOptim) IpoV4ServicesTunnelsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/tunnels")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query service tunnels used by service
func (ipO *IpOptim) IpoV4ServicesTunnelsonserviceServiceUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/tunnels-on-service/{serviceUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete a service endpoint from a service
func (ipO *IpOptim) IpoV4ServicesServiceUuidEndpointEndpointUuidDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/{serviceUuid}/endpoint/{endpointUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all endpoints associated to a service object
func (ipO *IpOptim) IpoV4ServicesServiceUuidEndpointsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/{serviceUuid}/endpoints")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a service by a unique identifier
func (ipO *IpOptim) IpoV4ServicesUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/{uuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete a service
func (ipO *IpOptim) IpoV4ServicesUuidDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/services/{uuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get SR Policies
func (ipO *IpOptim) IpoV4SrpoliciesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/srpolicies")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create SR Policies
func (ipO *IpOptim) IpoV4SrpoliciesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/srpolicies")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Trigger resignal on a Candidate Path
func (ipO *IpOptim) IpoV4SrpoliciesCandidatepathresignalIdPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/srpolicies/candidate-path-resignal/{id}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch a Candidate Path admin status
func (ipO *IpOptim) IpoV4SrpoliciesCandidatepathCandidatePathIdAdminAdminStatePatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/srpolicies/candidate-path/{candidatePathId}/admin/{adminState}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update Candidate Path. This request requires Candidate Path to be shutdown first. Add/Delete segment list(s) on the Candidate Path can be utilized for this request.
func (ipO *IpOptim) IpoV4SrpoliciesCandidatepathsIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/srpolicies/candidate-paths/{id}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Get SR Policies on link
func (ipO *IpOptim) IpoV4SrpoliciesOnlinkLinkIdRequestTypeGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/srpolicies/on-link/{linkId}/{requestType}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete SR Policies
func (ipO *IpOptim) IpoV4SrpoliciesPoliciesDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/srpolicies/policies")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update SR Policy. This request requires SR Policy to be shutdown first. Add/delete Candidate Path(s) can be utilized by this request.
func (ipO *IpOptim) IpoV4SrpoliciesPoliciesIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/srpolicies/policies/{id}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Fetch SR Policy list
func (ipO *IpOptim) IpoV4SrpoliciesPolicylistPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/srpolicies/policy-list")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete SR Policy
func (ipO *IpOptim) IpoV4SrpoliciesPolicyIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/srpolicies/policy/{id}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Patch a SR Policy admin status
func (ipO *IpOptim) IpoV4SrpoliciesPolicySrPolicyIdAdminAdminStatePatch(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/srpolicies/policy/{srPolicyId}/admin/{adminState}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Patch(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Trigger resignal on a list of SR Policies
func (ipO *IpOptim) IpoV4SrpoliciesResignalPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/srpolicies/resignal")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Trigger resignal all SR Policies
func (ipO *IpOptim) IpoV4SrpoliciesResignalallPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/srpolicies/resignal-all")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// NSP server is master
func (ipO *IpOptim) IpoV4SystemIsMasterGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/system/isMaster")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// BETA: Forces NSP Plugin to form a specific connection
func (ipO *IpOptim) IpoV4SystemPluginconnectPluginNamePluginVIdPluginKeyPluginNameVSR_NRCPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/system/plugin-connect/{pluginName}/{pluginVId}/{pluginKey}?pluginName=VSR_NRC")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Triggers a data synchronization with the connected NMS
func (ipO *IpOptim) IpoV4SystemResyncnmsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/system/resync-nms")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Triggers a data synchronization with the connected NMS on a specific object
func (ipO *IpOptim) IpoV4SystemResyncobjectUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/system/resync-object/{uuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Retrieves the state of the system
func (ipO *IpOptim) IpoV4SystemStateGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/system/state")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Retrieves the version of the system
func (ipO *IpOptim) IpoV4SystemVersionGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/system/version")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all C-Line service creation templates
func (ipO *IpOptim) IpoV4TemplateClineservicesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/cline-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a C-Line service creation template
func (ipO *IpOptim) IpoV4TemplateClineservicesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/cline-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a C-Line service creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateClineservicesTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/cline-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify a C-Line service creation template
func (ipO *IpOptim) IpoV4TemplateClineservicesTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/cline-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete a C-Line service creation template
func (ipO *IpOptim) IpoV4TemplateClineservicesTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/cline-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all Custom Attributes templates
func (ipO *IpOptim) IpoV4TemplateCustomattributesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/custom-attributes")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a Custom Attributes template
func (ipO *IpOptim) IpoV4TemplateCustomattributesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/custom-attributes")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a Custom Attributes template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateCustomattributesTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/custom-attributes/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify a Custom Attributes template
func (ipO *IpOptim) IpoV4TemplateCustomattributesTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/custom-attributes/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete a Custom Attributes template
func (ipO *IpOptim) IpoV4TemplateCustomattributesTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/custom-attributes/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all E-LAN service creation templates
func (ipO *IpOptim) IpoV4TemplateElanservicesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/elan-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an E-LAN service creation template
func (ipO *IpOptim) IpoV4TemplateElanservicesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/elan-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an E-LAN service creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateElanservicesTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/elan-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an E-LAN service creation template
func (ipO *IpOptim) IpoV4TemplateElanservicesTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/elan-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an E-LAN service creation template
func (ipO *IpOptim) IpoV4TemplateElanservicesTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/elan-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all E-Line service creation templates
func (ipO *IpOptim) IpoV4TemplateElineservicesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/eline-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an E-Line service creation template
func (ipO *IpOptim) IpoV4TemplateElineservicesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/eline-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an E-Line service creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateElineservicesTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/eline-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an E-Line service creation template
func (ipO *IpOptim) IpoV4TemplateElineservicesTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/eline-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an E-Line service creation template
func (ipO *IpOptim) IpoV4TemplateElineservicesTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/eline-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all generic QoS profiles.
func (ipO *IpOptim) IpoV4TemplateGenericqosGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/generic-qos")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a Generic QoS Profile
func (ipO *IpOptim) IpoV4TemplateGenericqosPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/generic-qos")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all generic QoS profiles of the provided network element
func (ipO *IpOptim) IpoV4TemplateGenericqosprofilesNeIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/generic-qos-profiles/{neId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a generic QoS profile by its unique ID
func (ipO *IpOptim) IpoV4TemplateGenericqosGqpIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/generic-qos/{gqpId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update a Generic QoS Profile
func (ipO *IpOptim) IpoV4TemplateGenericqosGqpIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/generic-qos/{gqpId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete a Generic QoS Profile
func (ipO *IpOptim) IpoV4TemplateGenericqosGqpIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/generic-qos/{gqpId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all IES service creation templates
func (ipO *IpOptim) IpoV4TemplateIesservicesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/ies-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an IES service creation template
func (ipO *IpOptim) IpoV4TemplateIesservicesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/ies-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an IES service creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateIesservicesTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/ies-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an IES service creation template
func (ipO *IpOptim) IpoV4TemplateIesservicesTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/ies-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an IES service creation template
func (ipO *IpOptim) IpoV4TemplateIesservicesTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/ies-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all L2 DCI VPN service creation templates
func (ipO *IpOptim) IpoV4TemplateL2dcivpnservicesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l2-dci-vpn-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an L2 DCI VPN service creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateL2dcivpnservicesTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l2-dci-vpn-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an L2 DCI VPN service creation template
func (ipO *IpOptim) IpoV4TemplateL2dcivpnservicesTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l2-dci-vpn-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all L2 Backhaul service creation templates
func (ipO *IpOptim) IpoV4TemplateL2backhaulservicesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l2backhaul-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an L2 backhaul service creation template
func (ipO *IpOptim) IpoV4TemplateL2backhaulservicesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l2backhaul-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an L2 Backhaul service creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateL2backhaulservicesTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l2backhaul-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an L2 Backhaul service creation template
func (ipO *IpOptim) IpoV4TemplateL2backhaulservicesTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l2backhaul-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an L2 Backhaul service creation template
func (ipO *IpOptim) IpoV4TemplateL2backhaulservicesTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l2backhaul-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all L3 DCI VPN service creation templates
func (ipO *IpOptim) IpoV4TemplateL3dcivpnservicesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l3-dci-vpn-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an L3 DCI VPN service creation template
func (ipO *IpOptim) IpoV4TemplateL3dcivpnservicesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l3-dci-vpn-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an L3 DCI VPN service creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateL3dcivpnservicesTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l3-dci-vpn-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an L3 DCI VPN service creation template
func (ipO *IpOptim) IpoV4TemplateL3dcivpnservicesTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l3-dci-vpn-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an L3 DCI VPN service creation template
func (ipO *IpOptim) IpoV4TemplateL3dcivpnservicesTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l3-dci-vpn-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all L3 VPN service creation templates
func (ipO *IpOptim) IpoV4TemplateL3vpnservicesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l3-vpn-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an L3 VPN service creation template
func (ipO *IpOptim) IpoV4TemplateL3vpnservicesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l3-vpn-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an L3 VPN service creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateL3vpnservicesTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l3-vpn-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an L3 VPN service creation template
func (ipO *IpOptim) IpoV4TemplateL3vpnservicesTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l3-vpn-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an L3 VPN service creation template
func (ipO *IpOptim) IpoV4TemplateL3vpnservicesTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/l3-vpn-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all LAG service creation templates
func (ipO *IpOptim) IpoV4TemplateLagservicesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/lag-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a LAG service creation template
func (ipO *IpOptim) IpoV4TemplateLagservicesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/lag-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a LAG service creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateLagservicesTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/lag-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify a LAG service creation template
func (ipO *IpOptim) IpoV4TemplateLagservicesTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/lag-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete a LAG service creation template
func (ipO *IpOptim) IpoV4TemplateLagservicesTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/lag-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an Mediation profile Mapping
func (ipO *IpOptim) IpoV4TemplateMediationprofilemappingPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/mediation-profile-mapping")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a mediation profile mapping by its unique identifier
func (ipO *IpOptim) IpoV4TemplateMediationprofilemappingProfileIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/mediation-profile-mapping/{profileId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update a mediation profile mapping
func (ipO *IpOptim) IpoV4TemplateMediationprofilemappingProfileIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/mediation-profile-mapping/{profileId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an Mediation profile mapping
func (ipO *IpOptim) IpoV4TemplateMediationprofilemappingProfileIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/mediation-profile-mapping/{profileId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all mediation profile mappings
func (ipO *IpOptim) IpoV4TemplateMediationprofilemappingsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/mediation-profile-mappings")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query NFM-P system scripts. Some of these scripts may be used for service creation or modification.
func (ipO *IpOptim) IpoV4TemplateNfmptemplateGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/nfmp-template")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all OCH service creation templates
func (ipO *IpOptim) IpoV4TemplateOchservicesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/och-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an OCH service creation template
func (ipO *IpOptim) IpoV4TemplateOchservicesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/och-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an OCH service creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateOchservicesTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/och-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an OCH service creation template
func (ipO *IpOptim) IpoV4TemplateOchservicesTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/och-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an OCH service creation template
func (ipO *IpOptim) IpoV4TemplateOchservicesTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/och-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all ODU service creation templates
func (ipO *IpOptim) IpoV4TemplateOduservicesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/odu-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an ODU service creation template
func (ipO *IpOptim) IpoV4TemplateOduservicesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/odu-services")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an ODU service creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateOduservicesTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/odu-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an ODU service creation template
func (ipO *IpOptim) IpoV4TemplateOduservicesTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/odu-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an ODU service creation template
func (ipO *IpOptim) IpoV4TemplateOduservicesTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/odu-services/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all optical connectivity constraint templates
func (ipO *IpOptim) IpoV4TemplateOpticalconnectivityconstraintGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-connectivity-constraint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an optical connectivity constraint creation template
func (ipO *IpOptim) IpoV4TemplateOpticalconnectivityconstraintPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-connectivity-constraint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an optical connectivity constraint creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateOpticalconnectivityconstraintTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-connectivity-constraint/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an optical connectivity constraint creation template
func (ipO *IpOptim) IpoV4TemplateOpticalconnectivityconstraintTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-connectivity-constraint/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an optical connectivity constraint creation template
func (ipO *IpOptim) IpoV4TemplateOpticalconnectivityconstraintTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-connectivity-constraint/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all optical connectivity service templates
func (ipO *IpOptim) IpoV4TemplateOpticalconnectivityserviceGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-connectivity-service")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an optical connectivity service creation template
func (ipO *IpOptim) IpoV4TemplateOpticalconnectivityservicePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-connectivity-service")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an optical connectivity service creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateOpticalconnectivityserviceTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-connectivity-service/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an optical connectivity service creation template
func (ipO *IpOptim) IpoV4TemplateOpticalconnectivityserviceTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-connectivity-service/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an optical connectivity service creation template
func (ipO *IpOptim) IpoV4TemplateOpticalconnectivityserviceTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-connectivity-service/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all optical resilience constraint templates
func (ipO *IpOptim) IpoV4TemplateOpticalresilienceconstraintGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-resilience-constraint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an optical resilience constraint creation template
func (ipO *IpOptim) IpoV4TemplateOpticalresilienceconstraintPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-resilience-constraint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an optical resilience constraint creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateOpticalresilienceconstraintTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-resilience-constraint/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an optical resilience constraint creation template
func (ipO *IpOptim) IpoV4TemplateOpticalresilienceconstraintTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-resilience-constraint/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an optical resilience constraint creation template
func (ipO *IpOptim) IpoV4TemplateOpticalresilienceconstraintTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-resilience-constraint/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all optical routing constraint templates
func (ipO *IpOptim) IpoV4TemplateOpticalroutingconstraintGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-routing-constraint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create an optical routing constraint creation template
func (ipO *IpOptim) IpoV4TemplateOpticalroutingconstraintPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-routing-constraint")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find an optical routing constraint creation template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateOpticalroutingconstraintTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-routing-constraint/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify an optical routing constraint creation template
func (ipO *IpOptim) IpoV4TemplateOpticalroutingconstraintTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-routing-constraint/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an optical routing constraint creation template
func (ipO *IpOptim) IpoV4TemplateOpticalroutingconstraintTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/optical-routing-constraint/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all path profile templates
func (ipO *IpOptim) IpoV4TemplatePathprofilesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/path-profiles")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a path profile template
func (ipO *IpOptim) IpoV4TemplatePathprofilesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/path-profiles")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a path profile template by its unique identifier
func (ipO *IpOptim) IpoV4TemplatePathprofilesTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/path-profiles/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update a path profile template
func (ipO *IpOptim) IpoV4TemplatePathprofilesTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/path-profiles/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete an existing path profile template
func (ipO *IpOptim) IpoV4TemplatePathprofilesTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/path-profiles/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all QoS templates
func (ipO *IpOptim) IpoV4TemplateQosGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/qos")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a QoS template
func (ipO *IpOptim) IpoV4TemplateQosPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/qos")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all QoS policies
func (ipO *IpOptim) IpoV4TemplateQospoliciesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/qos-policies")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all QoS policies of the provided network element
func (ipO *IpOptim) IpoV4TemplateQospoliciesNeIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/qos-policies/{neId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a QoS template by its unique identifier
func (ipO *IpOptim) IpoV4TemplateQosTemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/qos/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify a QoS template
func (ipO *IpOptim) IpoV4TemplateQosTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/qos/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete a QoS template
func (ipO *IpOptim) IpoV4TemplateQosTemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/qos/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all router ID system ID mapping policy
func (ipO *IpOptim) IpoV4TemplateRouteridsystemidmappingGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/router-id-system-id-mapping")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a router ID system ID mapping policy
func (ipO *IpOptim) IpoV4TemplateRouteridsystemidmappingPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/router-id-system-id-mapping")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a router-id-system-id-mapping by its unique identifier
func (ipO *IpOptim) IpoV4TemplateRouteridsystemidmappingPolicyIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/router-id-system-id-mapping/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify a router-id-system-id-mapping policy
func (ipO *IpOptim) IpoV4TemplateRouteridsystemidmappingPolicyIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/router-id-system-id-mapping/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete a router-id-mapping policy
func (ipO *IpOptim) IpoV4TemplateRouteridsystemidmappingPolicyIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/router-id-system-id-mapping/{policyId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all system IP MPLS Configuration
func (ipO *IpOptim) IpoV4TemplateSystemipmplsconfigGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/system-ip-mpls-config")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update a system ip mpls configuration
func (ipO *IpOptim) IpoV4TemplateSystemipmplsconfigTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/system-ip-mpls-config/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all tunnel creation templates. Deprecated: Use get-all-tunnel-selections instead.
func (ipO *IpOptim) IpoV4TemplateTunnelcreationsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/tunnel-creations")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Modify a tunnel creation template. Deprecated: update-tunnel-selections instead.
func (ipO *IpOptim) IpoV4TemplateTunnelcreationsTemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/tunnel-creations/{templateId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all workflow profiles
func (ipO *IpOptim) IpoV4TemplateWorkflowprofileGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/workflow-profile")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Create a workflow profile mapping
func (ipO *IpOptim) IpoV4TemplateWorkflowprofilePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/workflow-profile")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a workflow profile mapping by its unique identifier
func (ipO *IpOptim) IpoV4TemplateWorkflowprofileProfileIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/workflow-profile/{profileId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Update a workflow profile mapping
func (ipO *IpOptim) IpoV4TemplateWorkflowprofileProfileIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/workflow-profile/{profileId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Delete a workflow profile
func (ipO *IpOptim) IpoV4TemplateWorkflowprofileProfileIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/template/workflow-profile/{profileId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4TenantsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4TenantsPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all Customers
func (ipO *IpOptim) IpoV4TenantsCustomerallGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants/customer-all")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a customer by customer-Id
func (ipO *IpOptim) IpoV4TenantsCustomerCustomerIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants/customer/{customerId}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4TenantsResyncProviderGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants/resync/{provider}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4TenantsTenantUuidGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants/{tenantUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4TenantsTenantUuidPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants/{tenantUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Put(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4TenantsTenantUuidDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants/{tenantUuid}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4TenantsTenantUuidResourcesGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants/{tenantUuid}/resources")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4TenantsTenantUuidResourcesPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants/{tenantUuid}/resources")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4TenantsTenantUuidResourcesDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants/{tenantUuid}/resources")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4TenantsTenantUuidUsergroupGroupNameGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants/{tenantUuid}/usergroup/{groupName}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4TenantsTenantUuidUsergroupGroupNameDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants/{tenantUuid}/usergroup/{groupName}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Delete(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4TenantsTenantUuidUsergroupGroupNameRoleRoleTypePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/tenants/{tenantUuid}/usergroup/{groupName}/role/{roleType}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Query all Usergroups
func (ipO *IpOptim) IpoV4UsergroupsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/usergroups")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

// Find a Usergroup by a group name
func (ipO *IpOptim) IpoV4UsergroupsGroupNameGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/usergroups/{groupName}")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}

//
func (ipO *IpOptim) IpoV4UsergroupsGroupNameTenantsGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_IP_OPTIM_BASE_URL + "/v4/usergroups/{groupName}/tenants")
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Get(url)

	if err != nil {
		log.Error("NspRestconfInventory is unsuccesful: ", err)
		return
	}
	log.Info("Received Response " + urlHost + " Response: ")
	log.Info(resp.String())

	return resp.String()

}
