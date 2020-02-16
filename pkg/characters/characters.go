package characters

import (
	"fmt"

	"github.com/mbStavola/tarkov-api/pkg/client"
	"github.com/mbStavola/tarkov-api/pkg/locations"
)

type CharactersAPI struct {
	client *client.TarkovClient
}

func NewCharactersAPI(client *client.TarkovClient) CharactersAPI {
	return CharactersAPI{client}
}

type Boss struct {
	Name                  string             `json:"name"`
	Description           interface{}        `json:"description"`
	Image                 string             `json:"image"`
	Equipment             string             `json:"equipment"`
	Followers             string             `json:"followers"`
	SpawnChancePercentage int                `json:"spawn_chance_percentage"`
	Location              locations.Location `json:"location"`
}

func (api *CharactersAPI) Bosses() ([]Boss, error) {
	path := "bosses"

	var bosses []Boss
	if err := api.client.Get(path, &bosses); err != nil {
		return nil, err
	}

	return bosses, nil
}

func (api *CharactersAPI) BossesByLocation(locationName string) ([]Boss, error) {
	path := fmt.Sprintf("bosses/%s", locationName)

	var bosses []Boss
	if err := api.client.Get(path, &bosses); err != nil {
		return nil, err
	}

	return bosses, nil
}

type Dealer struct {
	Name        string `json:"name"`
	RealName    string `json:"real_name"`
	Description string `json:"description"`
	Biography   string `json:"biography"`
	Origin      string `json:"origin"`
	Image       string `json:"image"`
}

func (api *CharactersAPI) Dealers() ([]Dealer, error) {
	path := "dealers"

	var dealers []Dealer
	if err := api.client.Get(path, &dealers); err != nil {
		return nil, err
	}

	return dealers, nil
}
