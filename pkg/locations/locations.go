package locations

import (
	"fmt"

	"github.com/mbStavola/tarkov-api/pkg/client"
)

type LocationsAPI struct {
	client *client.TarkovClient
}

func NewLocationsAPI(client *client.TarkovClient) LocationsAPI {
	return LocationsAPI{client}
}

func (api *LocationsAPI) get(path string, ret interface{}) error {
	path = fmt.Sprintf("locations/%s", path)
	return api.client.Get(path, ret)
}

type Location struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	ReleaseStatus string `json:"release_status"`
	MainImage     string `json:"main_image"`
}

func (api LocationsAPI) Locations() ([]Location, error) {
	return api.LocationsByName("")
}

func (api LocationsAPI) LocationsByName(name string) ([]Location, error) {
	path := fmt.Sprintf("currencies/%s", name)

	var locations []Location
	if err := api.get(path, &locations); err != nil {
		return nil, err
	}

	return locations, nil
}
