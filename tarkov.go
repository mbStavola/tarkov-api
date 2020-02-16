package tarkov

import (
	"net/http"

	"github.com/mbStavola/tarkov-api/pkg/characters"
	"github.com/mbStavola/tarkov-api/pkg/client"
	"github.com/mbStavola/tarkov-api/pkg/economy"
	"github.com/mbStavola/tarkov-api/pkg/items"
	"github.com/mbStavola/tarkov-api/pkg/locations"
)

type TarkovAPI struct {
	client     *client.TarkovClient
	Characters characters.CharactersAPI
	Economy    economy.EconomyAPI
	Items      items.ItemsAPI
	Locations  locations.LocationsAPI
}

func NewTarkovAPI(sessionID string) TarkovAPI {
	return NewCustomTarkovAPI(nil, sessionID)
}

func NewCustomTarkovAPI(httpClient *http.Client, sessionID string) TarkovAPI {
	client := client.NewClient(httpClient, sessionID)

	return TarkovAPI{
		client:     &client,
		Characters: characters.NewCharactersAPI(&client),
		Economy:    economy.NewEconomyAPI(&client),
		Items:      items.NewItemsAPI(&client),
		Locations:  locations.NewLocationsAPI(&client),
	}
}
