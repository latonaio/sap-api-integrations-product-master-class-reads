package sap_api_output_formatter

type ProductMasterClassReads struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	APISchema          string `json:"api_schema"`
	ProductMasterClass string `json:"product_master_class"`
	Deleted            bool   `json:"deleted"`
}    
    
type ProductGeneral struct {
	Product                       string      `json:"Product"`
	ProductType                   string      `json:"ProductType"`
	CreationDate                  string      `json:"CreationDate"`
	LastChangeDate                string      `json:"LastChangeDate"`
	IsMarkedForDeletion           bool        `json:"IsMarkedForDeletion"`
	ProductGroup                  string      `json:"ProductGroup"`
	BaseUnit                      string      `json:"BaseUnit"`
	ProductHierarchy              string      `json:"ProductHierarchy"`
	ToProductClass                string      `json:"to_ProductClass"`
	ToProductCharc                string      `json:"to_ProductCharc"`
}

type ToProductClass struct {
	Product                  string `json:"Product"`
	ClassInternalID          string `json:"ClassInternalID"`
	KeyDate                  string `json:"KeyDate"`
	ChangeNumber             string `json:"ChangeNumber"`
	ClassType                string `json:"ClassType"`
	ToClassDetails           string `json:"to_ClassDetails"`
	ToCharc                  string `json:"to_Characteristics"`
}

type ToClassDetails struct {
	ClassInternalID               string      `json:"ClassInternalID"`
	ClassType                     string      `json:"ClassType"`
	ClassTypeName                 string      `json:"ClassTypeName"`
	Class                         string      `json:"Class"`
	ClassStatus                   string      `json:"ClassStatus"`
	ClassStatusName               string      `json:"ClassStatusName"`
	ClassGroup                    string      `json:"ClassGroup"`
	ClassGroupName                string      `json:"ClassGroupName"`
	CreationDate                  string      `json:"CreationDate"`
	LastChangeDate                string      `json:"LastChangeDate"`
	ValidityStartDate             string      `json:"ValidityStartDate"`
	ValidityEndDate               string      `json:"ValidityEndDate"`
	ClassLastChangedDateTime      string      `json:"ClassLastChangedDateTime"`
	KeyDate                       string      `json:"KeyDate"`
}

type ToProductCharc struct {
	Product         string `json:"Product"`
	ClassInternalID string `json:"ClassInternalID"`
	CharcInternalID string `json:"CharcInternalID"`
	KeyDate         string `json:"KeyDate"`
	ChangeNumber    string `json:"ChangeNumber"`
	ClassType       string `json:"ClassType"`
}
