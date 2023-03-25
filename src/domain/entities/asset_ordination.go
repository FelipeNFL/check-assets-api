package entities

type Ordination string

const (
	Alphabetical Ordination = "alphabetical"
	Price        Ordination = "price"
	Custom       Ordination = "custom"
)

type Direction string

const (
	Asc  Direction = "asc"
	Desc Direction = "desc"
)

type AssetOrdination struct {
	Ordination  Ordination `json:"ordination"`
	CustomOrder []string   `json:"custom_order,omitempty"`
}

func NewAssetOrdination(ordination string, customOrder []string) (AssetOrdination, error) {
	ordinationParsed := Ordination(ordination)

	if ordinationParsed != Alphabetical && ordinationParsed != Price && ordinationParsed != Custom {
		return AssetOrdination{}, ErrAssetOrdinationInvalid
	}

	return AssetOrdination{Ordination: ordinationParsed, CustomOrder: customOrder}, nil
}
