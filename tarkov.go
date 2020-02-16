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
	client *client.TarkovClient
}

func NewTarkovAPI(sessionID string) TarkovAPI {
	client := client.NewClient(nil, sessionID)
	return TarkovAPI{
		client: &client,
	}
}

func NewCustomTarkovAPI(httpClient *http.Client, sessionID string) TarkovAPI {
	client := client.NewClient(httpClient, sessionID)
	return TarkovAPI{
		client: &client,
	}
}

func (api *TarkovAPI) Economy() economy.EconomyAPI {
	return economy.NewEconomyAPI(api.client)
}

func (api *TarkovAPI) Characters() characters.CharactersAPI {
	return characters.NewCharactersAPI(api.client)
}

func (api *TarkovAPI) Locations() locations.LocationsAPI {
	return locations.NewLocationsAPI(api.client)
}

func (api *TarkovAPI) Items() items.ItemsAPI {
	return items.NewItemsAPI(api.client)
}
