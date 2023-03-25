package entities

type Ordination string

const (
	Alphabetical Ordination = "alphabetical"
	Price        Ordination = "price"
	Custom       Ordination = "custom"
)

type AssetOrdination struct {
	Ordination Ordination `json:"ordination"`
}

func NewAssetOrdination(ordination Ordination) (AssetOrdination, error) {
	if ordination != Alphabetical && ordination != Price && ordination != Custom {
		return AssetOrdination{}, ErrAssetOrdinationInvalid
	}

	return AssetOrdination{Ordination: ordination}, nil
}
