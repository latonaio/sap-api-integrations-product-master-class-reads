package main

import (
	sap_api_caller "sap-api-integrations-product-master-class-reads/SAP_API_Caller"
	"sap-api-integrations-product-master-class-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Product_Master_Class_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"ProductGeneral",
		}
	}

	caller.AsyncGetProductMasterClass(
		inoutSDC.ProductClass.Product,

		accepter,
	)
}
