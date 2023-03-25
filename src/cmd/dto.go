package cmd

type CreateNewAssetDTO struct {
	Code string `json:"code"`
}

type SaveAssetOrdinationDTO struct {
	Ordination  string   `json:"ordination"`
	CustomOrder []string `json:"custom_order,omitempty"`
}
