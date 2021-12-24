package responses

type ProductGeneral struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Product                       string      `json:"Product"`
			ProductType                   string      `json:"ProductType"`
			CreationDate                  string      `json:"CreationDate"`
			LastChangeDate                string      `json:"LastChangeDate"`
			IsMarkedForDeletion           bool        `json:"IsMarkedForDeletion"`
			ProductGroup                  string      `json:"ProductGroup"`
			BaseUnit                      string      `json:"BaseUnit"`
			ProductHierarchy              string      `json:"ProductHierarchy"`
			ToProductClass struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ProductClass"`
			ToProductCharc struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ProductCharc"`
		} `json:"results"`
	} `json:"d"`
}
