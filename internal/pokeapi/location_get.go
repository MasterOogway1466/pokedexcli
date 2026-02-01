package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (RespLocation, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + locationName

	// 1. Check Cache
	if val, ok := c.cache.Get(url); ok {
		locationResp := RespLocation{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespLocation{}, err
		}
		return locationResp, nil
	}

	// 2. Request Data
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}

	// 3. Unmarshal & Cache
	locationResp := RespLocation{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocation{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
