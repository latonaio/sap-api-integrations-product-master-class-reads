package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-product-master-classification-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetProductMasterClass(product string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(product)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) General(product string) {
	generalData, err := c.callProductMasterClassSrvAPIRequirementGeneral("A_ClfnProduct", product)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(generalData)

	productClassData, err := c.callToProductClass(generalData[0].ToProductClass)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(productClassData)

	productClassDetailsData, err := c.callToProductClassDetails(productClassData[0].ToProductClassDetails)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(productClassDetailsData)
	
	productCharcData, err := c.callToProductCharc(generalData[0].ToProductCharc)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(productCharcData)

}

func (c *SAPAPICaller) callProductMasterClassSrvAPIRequirementGeneral(api, product string) ([]sap_api_output_formatter.General, error) {
	url := strings.Join([]string{c.baseURL, "API_CLFN_PRODUCT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithGeneral(req, product)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToProductClass(url string) ([]sap_api_output_formatter.ToProductClass, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToProductClass(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToProductClassDetails(url string) (*sap_api_output_formatter.ToProductClassDetails, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToProductClassDetails(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToProductCharc(url string) ([]sap_api_output_formatter.ToProductCharc, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToProductCharc(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}


func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithGeneral(req *http.Request, product string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Product eq '%s'", product))
	req.URL.RawQuery = params.Encode()
}
