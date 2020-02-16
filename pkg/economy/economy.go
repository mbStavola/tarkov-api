package economy

import (
	"github.com/mbStavola/tarkov-api/pkg/client"
)

type EconomyAPI struct {
	client *client.TarkovClient
}

func NewEconomyAPI(client *client.TarkovClient) EconomyAPI {
	return EconomyAPI{client}
}

type Currency struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Image        string `json:"image"`
	Symbol       string `json:"symbol"`
}

func (api EconomyAPI) Currencies() ([]Currency, error) {
	path := "currencies"

	var currencies []Currency
	if err := api.client.Get(path, &currencies); err != nil {
		return nil, err
	}

	return currencies, nil
}
