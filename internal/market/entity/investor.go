package entity 

type Investor struct { // investidor 
	ID string
	Name string
	AssetPosition []*InvestorAssetPosition
}

type InvestorAssetPosition struct { // papéis do investidor
	AssetID string
	Shares int
}

func NewInvestor(id string) *Investor{ // novo investidor
	return &Investor{
		ID: id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition * InvestorAssetPosition) { // adiciona um novo papel no array AssetPosition
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetID string, qtdShares int) { // atualiza as posições dos assets
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetID, qtdShares))
	}else {
		assetPosition.Shares += qtdShares
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition { // pega a posição do papel 
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil
}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition { // cia um novo papel
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares: shares,
	}
}

