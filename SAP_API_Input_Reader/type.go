package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	SalesOrder struct {
		SalesOrder                string      `json:"document_no"`
		Plant                     string      `json:"deliver_to"`
		OrderQuantity             string      `json:"quantity"`
		PickedQuantity            string      `json:"picked_quantity"`
		NetPriceAmount            string      `json:"price"`
	    Batch                     string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             string      `json:"quantity"`
		CompletedQuantity    string      `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 string      `json:"quantity"`
			CompletedQuantity        string      `json:"completed_quantity"`
			ErroredQuantity          string      `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity string      `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                         string `json:"api_schema"`
	Accepter                        []string `json:"accepter"`
	Product                           string `json:"material_code"`
	Supplier                          string `json:"plant/supplier"`
	Stock                             string `json:"stock"`
	SalesOrderType                    string `json:"document_type"`
	SONumber                          string `json:"document_no"`
	ScheduleLineDeliveryDate          string `json:"planned_date"`
	ValidatedDate                     string `json:"validated_date"`
    Deleted                           bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       struct {
		Product             string      `json:"Product"`
		ProductType         string      `json:"ProductType"`
		CreationDate        string      `json:"CreationDate"`
		LastChangeDate      string      `json:"LastChangeDate"`
		IsMarkedForDeletion bool        `json:"IsMarkedForDeletion"`
		ProductGroup        string      `json:"ProductGroup"`
		BaseUnit            string      `json:"BaseUnit"`
		ProductHierarchy    string      `json:"ProductHierarchy"`
		ProductClass        struct {
			ClassInternalID     string `json:"ClassInternalID"`
			KeyDate             string `json:"KeyDate"`
			ChangeNumber        string `json:"ChangeNumber"`
			ClassType           string `json:"ClassType"`
			ClassDetails        struct {
				ClassType                string `json:"ClassType"`
				ClassTypeName            string `json:"ClassTypeName"`
				Class                    string `json:"Class"`
				ClassStatus              string `json:"ClassStatus"`
				ClassStatusName          string `json:"ClassStatusName"`
				ClassGroup               string `json:"ClassGroup"`
				ClassGroupName           string `json:"ClassGroupName"`
				CreationDate             string `json:"CreationDate"`
				LastChangeDate           string `json:"LastChangeDate"`
				ValidityStartDate        string `json:"ValidityStartDate"`
				ValidityEndDate          string `json:"ValidityEndDate"`
				ClassLastChangedDateTime string `json:"ClassLastChangedDateTime"`
				KeyDate                  string `json:"KeyDate"`
			} `json:"ClassDetails"`
		} `json:"ProductClass"`
		ProductCharacteristic struct {
			ClassInternalID string `json:"ClassInternalID"`
			CharcInternalID string `json:"CharcInternalID"`
			KeyDate         string `json:"KeyDate"`
			ChangeNumber    string `json:"ChangeNumber"`
			ClassType       string `json:"ClassType"`
		} `json:"ProductCharacteristic"`
	} `json:"Product"`
	APISchema   string   `json:"api_schema"`
	Accepter    []string `json:"accepter"`
	ProductCode string   `json:"product_code"`
	Deleted     bool     `json:"deleted"`
}
