package Handler

import (
	"encoding/json"
	"net/http"
)



func Get_JSONData(url string) interface{} {
	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		return nil
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		return nil
	}
	return artists
}
