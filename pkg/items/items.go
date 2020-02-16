package items

import (
	"fmt"

	"github.com/mbStavola/tarkov-api/pkg/client"
)

type ItemsAPI struct {
	client *client.TarkovClient
}

func NewItemsAPI(client *client.TarkovClient) ItemsAPI {
	return ItemsAPI{client}
}

func (api *ItemsAPI) get(path string, ret interface{}) error {
	path = fmt.Sprintf("items/%s", path)
	return api.client.Get(path, ret)
}

type AmmunitionType struct {
	Name    string `json:"name"`
	Caliber string `json:"caliber"`
	Image   string `json:"image"`
}

func (api ItemsAPI) AmmunitionTypes() ([]AmmunitionType, error) {
	path := "ammunition"

	var types []AmmunitionType
	if err := api.get(path, &types); err != nil {
		return nil, err
	}

	return types, nil
}

type Ammunition struct {
	Name                string  `json:"name"`
	Damage              string  `json:"damage"`
	PenetrationPower    int     `json:"penetration_power"`
	ArmorDamage         string  `json:"armor_damage"`
	FragmentationChance float64 `json:"fragmentation_chance"`
	RicochetChance      int     `json:"ricochet_chance"`
	Speed               int     `json:"speed"`
	SpecialEffects      string  `json:"special_effects"`
	SoldBy              string  `json:"sold_by"`
}

func (api ItemsAPI) Ammunition(caliber string) ([]Ammunition, error) {
	path := fmt.Sprintf("currencies/%s", caliber)

	var ammunition []Ammunition
	if err := api.get(path, &ammunition); err != nil {
		return nil, err
	}

	return ammunition, nil
}
