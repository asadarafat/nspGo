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
}

func init() {
	// init logConfig
	toolLogger := nspgotools.Tools{}
	toolLogger.InitLogger("./logs/nspGo-wfm.log", 1)
}

// Query all actions
func (wfm *Wfm) WfmV1ActionGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action")
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

// Create an action using Mistral DSL
func (wfm *Wfm) WfmV1ActionPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action")
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

// Query all action-executions
func (wfm *Wfm) WfmV1ActionexecutionGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action-execution")
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

// Run an action-execution
func (wfm *Wfm) WfmV1ActionexecutionPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action-execution")
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

// fetch the mistral_http action-executions
func (wfm *Wfm) WfmV1ActionexecutionAsyncWaitingGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action-execution/asyncWaiting")
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

// Query action-executions by task
func (wfm *Wfm) WfmV1ActionexecutionTaskIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action-execution/task/{id}")
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

// Query for an action-execution
func (wfm *Wfm) WfmV1ActionexecutionIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action-execution/{id}")
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

// Update an action-execution
func (wfm *Wfm) WfmV1ActionexecutionIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action-execution/{id}")
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

// Delete an action-execution
func (wfm *Wfm) WfmV1ActionexecutionIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action-execution/{id}")
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

// Create an action using Mistral DSL
func (wfm *Wfm) WfmV1ActionDefinitionPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action/definition")
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

// Bulk export actions with filter
func (wfm *Wfm) WfmV1ActionExportGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action/export")
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

// Validate an action definition using Mistral v2 DSL YAML
func (wfm *Wfm) WfmV1ActionValidatePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action/validate")
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

// Query for an action
func (wfm *Wfm) WfmV1ActionIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action/{id}")
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

// Delete an action
func (wfm *Wfm) WfmV1ActionIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action/{id}")
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

// Clone an existing action
func (wfm *Wfm) WfmV1ActionIdClonePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action/{id}/clone")
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

// Query for action definition
func (wfm *Wfm) WfmV1ActionIdDefinitionGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action/{id}/definition")
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

// Update action
func (wfm *Wfm) WfmV1ActionIdDefinitionPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action/{id}/definition")
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

// Export an action to a yaml file
func (wfm *Wfm) WfmV1ActionIdExportGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/action/{id}/export")
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

// Query all workflows
func (wfm *Wfm) WfmV1AdminWorkflowGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/admin/workflow")
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

// Put workflow user groups
func (wfm *Wfm) WfmV1AdminWorkflowUsergroupsPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/admin/workflow/usergroups")
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

// Delete workflow user groups
func (wfm *Wfm) WfmV1AdminWorkflowUsergroupsDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/admin/workflow/usergroups")
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

// Query all cron-triggers
func (wfm *Wfm) WfmV1CrontriggerGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/cron-trigger")
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

// Create a cron-trigger
func (wfm *Wfm) WfmV1CrontriggerPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/cron-trigger")
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

// Query for a cron-trigger by Workflow ID
func (wfm *Wfm) WfmV1CrontriggerWorkflowIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/cron-trigger/workflow/{id}")
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

// Query for a cron-trigger
func (wfm *Wfm) WfmV1CrontriggerIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/cron-trigger/{id}")
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

// Delete a cron-trigger
func (wfm *Wfm) WfmV1CrontriggerIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/cron-trigger/{id}")
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

// Query all executions
func (wfm *Wfm) WfmV1ExecutionGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/execution")
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

// Create an execution
func (wfm *Wfm) WfmV1ExecutionPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/execution")
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

// Create a synchronous execution
func (wfm *Wfm) WfmV1ExecutionSynchronousPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/execution/synchronous")
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

// Query executions for a workflow
func (wfm *Wfm) WfmV1ExecutionWorkflowIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/execution/workflow/{id}")
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

// Query for an execution
func (wfm *Wfm) WfmV1ExecutionIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/execution/{id}")
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

// Update an execution
func (wfm *Wfm) WfmV1ExecutionIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/execution/{id}")
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

// Delete an execution
func (wfm *Wfm) WfmV1ExecutionIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/execution/{id}")
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

// Gather Quick View information for Execution
func (wfm *Wfm) WfmV1ExecutionIdQuickviewGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/execution/{id}/quickview")
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

// Test a filter
func (wfm *Wfm) WfmV1FiltertestPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/filter-test")
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

// Import a Git Workflow
func (wfm *Wfm) WfmV1GitIntegrationWorkflowImportPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/gitIntegration/workflow/import")
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

// Update a Git workflow definition
func (wfm *Wfm) WfmV1GitIntegrationWorkflowIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/gitIntegration/workflow/{id}")
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

// Query all jinja templates
func (wfm *Wfm) WfmV1JinjatemplateGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/jinja-template")
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

// Create a jinja template
func (wfm *Wfm) WfmV1JinjatemplatePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/jinja-template")
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

// Bulk export jinja templates
func (wfm *Wfm) WfmV1JinjatemplateExportGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/jinja-template/export")
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

// Parse Jinja Template YAML Definition
func (wfm *Wfm) WfmV1JinjatemplateValidatePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/jinja-template/validate")
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

// Query for a jinja template
func (wfm *Wfm) WfmV1JinjatemplateIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/jinja-template/{id}")
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

// Update a jinja template
func (wfm *Wfm) WfmV1JinjatemplateIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/jinja-template/{id}")
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

// Delete a jinja template
func (wfm *Wfm) WfmV1JinjatemplateIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/jinja-template/{id}")
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

// Query for a jinja template definition
func (wfm *Wfm) WfmV1JinjatemplateIdDefinitionGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/jinja-template/{id}/definition")
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

// Export a jinja template to a yaml file
func (wfm *Wfm) WfmV1JinjatemplateIdExportGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/jinja-template/{id}/export")
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

// Query for a jinja template content
func (wfm *Wfm) WfmV1JinjatemplateIdTemplateGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/jinja-template/{id}/template")
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

// Query all kafka triggers
func (wfm *Wfm) WfmV1KafkatriggerGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/kafka-trigger")
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

// Create a kafka trigger
func (wfm *Wfm) WfmV1KafkatriggerPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/kafka-trigger")
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

// Parse Kafka Trigger YAML Definition
func (wfm *Wfm) WfmV1KafkatriggerValidatePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/kafka-trigger/validate")
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

// Query all kafka triggers for a given workflow id
func (wfm *Wfm) WfmV1KafkatriggerWorkflowIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/kafka-trigger/workflow/{id}")
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

// Query for a kafka trigger
func (wfm *Wfm) WfmV1KafkatriggerIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/kafka-trigger/{id}")
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

// Update a kafka trigger
func (wfm *Wfm) WfmV1KafkatriggerIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/kafka-trigger/{id}")
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

// Delete a kafka trigger
func (wfm *Wfm) WfmV1KafkatriggerIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/kafka-trigger/{id}")
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

// Reset a kafka trigger counter
func (wfm *Wfm) WfmV1KafkatriggerIdCounterPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/kafka-trigger/{id}/counter")
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

// Export kafka trigger to a yaml file
func (wfm *Wfm) WfmV1KafkatriggerIdExportGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/kafka-trigger/{id}/export")
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
func (wfm *Wfm) WfmV1NotificationProducerFilterUpdatePut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/notification/producer/filter/update")
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
func (wfm *Wfm) WfmV1NotificationProducerFilterValidationPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/notification/producer/filter/validation")
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
func (wfm *Wfm) WfmV1NotificationProducerPropertiesValidationPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/notification/producer/properties/validation")
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
func (wfm *Wfm) WfmV1NotificationProducerPublishingStartPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/notification/producer/publishing/start")
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
func (wfm *Wfm) WfmV1NotificationProducerPublishingStopPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/notification/producer/publishing/stop")
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

// Gather Statistics for Dashboard
func (wfm *Wfm) WfmV1StatisticsDashboardGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/statistics/dashboard")
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

// Gather Statistics for Execution Dashboard
func (wfm *Wfm) WfmV1StatisticsExecutionIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/statistics/execution/{id}")
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

// Gather Statistics for Task Dashboard
func (wfm *Wfm) WfmV1StatisticsTaskIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/statistics/task/{id}")
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

// Gather Statistics for Workflow Dashboard
func (wfm *Wfm) WfmV1StatisticsWorkflowIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/statistics/workflow/{id}")
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

// Query System Status
func (wfm *Wfm) WfmV1SystemstatusGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/system-status")
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

// Query all tasks
func (wfm *Wfm) WfmV1TaskGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/task")
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

// Query tasks by execution
func (wfm *Wfm) WfmV1TaskExecutionIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/task/execution/{id}")
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

// Query for a task
func (wfm *Wfm) WfmV1TaskIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/task/{id}")
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

// Re-runs a failed task or resumes task using with-items from the failed index
func (wfm *Wfm) WfmV1TaskIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/task/{id}")
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

// Query all workflows
func (wfm *Wfm) WfmV1WorkflowGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow")
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

// Create a workflow using Mistral DSL with escaped characters
func (wfm *Wfm) WfmV1WorkflowPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow")
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

// Query all workflow environments
func (wfm *Wfm) WfmV1WorkflowenvironmentGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow-environment")
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

// Create a workflow environment
func (wfm *Wfm) WfmV1WorkflowenvironmentPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow-environment")
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

// Parse Workflow Environment YAML Definition
func (wfm *Wfm) WfmV1WorkflowenvironmentParseDefinitionPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow-environment/parseDefinition")
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

// Validate Workflow Environment YAML Definition
func (wfm *Wfm) WfmV1WorkflowenvironmentValidatePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow-environment/validate")
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

// Query for a workflow environment
func (wfm *Wfm) WfmV1WorkflowenvironmentIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow-environment/{id}")
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

// Update a workflow environment
func (wfm *Wfm) WfmV1WorkflowenvironmentIdPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow-environment/{id}")
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

// Delete a workflow environment
func (wfm *Wfm) WfmV1WorkflowenvironmentIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow-environment/{id}")
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

// Query for a workflow environment definition
func (wfm *Wfm) WfmV1WorkflowenvironmentIdDefinitionGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow-environment/{id}/definition")
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

// Export a workflow environment to a yaml file
func (wfm *Wfm) WfmV1WorkflowenvironmentIdExportGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow-environment/{id}/export")
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

// Create a workflow using Mistral DSL
func (wfm *Wfm) WfmV1WorkflowDefinitionPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/definition")
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

// Bulk export workflows with filter
func (wfm *Wfm) WfmV1WorkflowExportGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/export")
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

// Import a Workflow package
func (wfm *Wfm) WfmV1WorkflowImportPackagePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/importPackage")
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

// Parse Workflow YAML Definition
func (wfm *Wfm) WfmV1WorkflowParseDefinitionPost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/parseDefinition")
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

// Validate a workflow definition using Mistral v2 DSL YAML
func (wfm *Wfm) WfmV1WorkflowValidatePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/validate")
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

// Query for a workflow
func (wfm *Wfm) WfmV1WorkflowIdGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/{id}")
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

// Delete a workflow
func (wfm *Wfm) WfmV1WorkflowIdDelete(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/{id}")
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

// Clone an existing workflow and corresponding artifacts
func (wfm *Wfm) WfmV1WorkflowIdClonePost(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/{id}/clone")
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

// Query for a workflow definition
func (wfm *Wfm) WfmV1WorkflowIdDefinitionGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/{id}/definition")
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

// Update a workflow
func (wfm *Wfm) WfmV1WorkflowIdDefinitionPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/{id}/definition")
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

// Export a workflow to a yaml file
func (wfm *Wfm) WfmV1WorkflowIdExportGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/{id}/export")
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

// Gather Quick View information on the latest Execution of a Workflow
func (wfm *Wfm) WfmV1WorkflowIdQuickviewGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/{id}/quickview")
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

// Update a workflow readme
func (wfm *Wfm) WfmV1WorkflowIdReadmePut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/{id}/readme")
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

// Update workflow status
func (wfm *Wfm) WfmV1WorkflowIdStatusPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/{id}/status")
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

// Query for yang and UI schema for a workflow
func (wfm *Wfm) WfmV1WorkflowIdUiGet(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/{id}/ui")
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

// Create a UI object for a workflow
func (wfm *Wfm) WfmV1WorkflowIdUiPut(urlHost string, token string, proxyEnable string, proxyAddress string, payload []byte) (result string) {
	client := resty.New()
	client.SetTimeout(6000 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	if proxyEnable == "true" {
		client.SetProxy(proxyAddress)
	}

	url := ("https://" + urlHost + nspgoconstants.GLBL_NSP_WFM_BASE_URL + "/v1/workflow/{id}/ui")
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
