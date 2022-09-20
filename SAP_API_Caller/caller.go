package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-master-recipe-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetMasterRecipe(masterRecipeGroup, masterRecipe string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(masterRecipeGroup, masterRecipe)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(masterRecipeGroup, masterRecipe string) {
	headerData, err := c.callMasterRecipeSrvAPIRequirementHeader("MasterRecipeHeader", masterRecipeGroup, masterRecipe)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	materialAssignmentData, err := c.callToMaterialAssignment(headerData[0].ToMaterialAssignment)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(materialAssignmentData)

	operationData, err := c.callToOperation(headerData[0].ToOperation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(operationData)

	phaseData, err := c.callToPhase(operationData[0].ToPhase)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(phaseData)

	phaseRelationshipData, err := c.callToPhaseRelationship(phaseData[0].ToPhaseRelationship)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(phaseRelationshipData)

}

func (c *SAPAPICaller) callMasterRecipeSrvAPIRequirementHeader(api, masterRecipeGroup, masterRecipe string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_MASTER_RECIPE", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, masterRecipeGroup, masterRecipe)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToMaterialAssignment(url string) ([]sap_api_output_formatter.ToMaterialAssignment, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToMaterialAssignment(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToOperation(url string) ([]sap_api_output_formatter.ToOperation, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToOperation(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPhase(url string) ([]sap_api_output_formatter.ToPhase, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPhase(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPhaseRelationship(url string) ([]sap_api_output_formatter.ToPhaseRelationship, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPhaseRelationship(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, masterRecipeGroup, masterRecipe string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("MasterRecipeGroup eq '%s' and MasterRecipe eq '%s'", masterRecipeGroup, masterRecipe)
	return params
}
