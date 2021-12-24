package responses

type ToProductClass struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Product           string `json:"Product"`
			ClassInternalID   string `json:"ClassInternalID"`
			KeyDate           string `json:"KeyDate"`
			ChangeNumber      string `json:"ChangeNumber"`
			ClassType         string `json:"ClassType"`
			ToClassDetails struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ClassDetails"`
			ToCharc struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Characteristics"`
		} `json:"results"`
	} `json:"d"`
}
