package pokeapi

type LocationResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type RespShallowLocations struct {
	Count    int              `json:"count"`
	Next     *string          `json:"next"`
	Previous *string          `json:"previous"`
	Results  []LocationResult `json:"results"`
}
