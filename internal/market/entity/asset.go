package entity

type Asset struct { // papel
	ID						string
	Name					string
	MarketVolume	int 
}

func NewAsset(id string, name string, marketVolume int) *Asset { // cria um novo papel
	return &Asset {
		ID:						id,
		Name: 				name,
		MarketVolume: marketVolume	
	}
}