package responses

type ToClassDetails struct {
	D struct {
		Metadata struct {
			ID   string `json:"id"`
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
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
	} `json:"d"`
}
