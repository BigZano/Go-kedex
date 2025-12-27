package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationData(locationName string) (RespExploreLocations, error) {
	fullURL := baseURL + "location-area/" + locationName

	if cachedData, found := c.cache.Get(fullURL); found {
		exploreResp := RespExploreLocations{}
		err := json.Unmarshal(cachedData, &exploreResp)
		if err != nil {
			return RespExploreLocations{}, err
		}
		return exploreResp, nil
	}
	res, err := http.Get(fullURL)
	if err != nil {
		return RespExploreLocations{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(res.Body)
		return RespExploreLocations{}, fmt.Errorf("bad status %d from %s\nBody: %s", res.StatusCode, fullURL, string(b))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespExploreLocations{}, err
	}

	c.cache.Add(fullURL, body)

	exploreResp := RespExploreLocations{}
	err = json.Unmarshal(body, &exploreResp)
	if err != nil {
		return RespExploreLocations{}, err
	}
	return exploreResp, nil
}
