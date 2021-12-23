package responses

type ToProductCharc struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Product         string `json:"Product"`
			ClassInternalID string `json:"ClassInternalID"`
			CharcInternalID string `json:"CharcInternalID"`
			KeyDate         string `json:"KeyDate"`
			ChangeNumber    string `json:"ChangeNumber"`
			ClassType       string `json:"ClassType"`
		} `json:"results"`
	} `json:"d"`
}
