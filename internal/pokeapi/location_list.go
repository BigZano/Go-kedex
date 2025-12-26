package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	fullURL := baseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}
	if cachedData, found := c.cache.Get(fullURL); found {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(cachedData, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}
	res, err := http.Get(fullURL)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(fullURL, body)

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return locationsResp, nil
}
