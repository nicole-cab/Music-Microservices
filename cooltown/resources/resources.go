package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

// to return the audio base64 encoded
type Track struct {
	Audio string
}

// creates a http request to the other microservices
func httpRequest(w http.ResponseWriter, url string, method string, data []byte) map[string]interface{} {

	httpPostUrl := url
	fmt.Println("HTTP URL:", httpPostUrl)

	// make the http request
	request, error := http.NewRequest(method, httpPostUrl, bytes.NewBuffer(data))

	if error != nil {
		w.WriteHeader(500)
		return nil
	}

	// set the header
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		w.WriteHeader(500)
		return nil
	}
	defer response.Body.Close()

	// get the status of the request
	if response.StatusCode == 200 {
		if body, err := io.ReadAll(response.Body); err == nil {

			result := map[string]interface{}{}
			json.Unmarshal(body, &result)

			// return the response body
			return result
		}

	} else if response.StatusCode == 400 { // Bad Request
		w.WriteHeader(400)
		return nil

	} else if response.StatusCode == 404 { // Not Found
		w.WriteHeader(404)
		return nil

	}

	w.WriteHeader(500) // internal server error
	return nil

}

// connects to the search and tracks microservices to return the original track title
func getTrack(w http.ResponseWriter, r *http.Request) {
	t := map[string]interface{}{} // t stores the request data
	if err := json.NewDecoder(r.Body).Decode(&t); err == nil {

		// audio is the audio bas64 encoded
		if audio, ok := (t["Audio"].(string)); ok {
			if t["Audio"].(string) != "" {

				// get the id of the track using the search microservice
				var jsonData = []byte(`{
					"Audio":"` + audio + `"}`)

				result := httpRequest(w, "http://localhost:3001/search", "POST", jsonData) // will write the corresponding status code

				if result == nil {
					return
				}

				id := result["Id"]
				fmt.Println("id: " + id.(string))

				// replace the spaces with "+"
				newId := strings.ReplaceAll(id.(string), " ", "+")
				fmt.Println("new id: " + newId)

				// escape the id string
				newId = url.QueryEscape(newId)
				fmt.Println("new id escaped: " + newId)

				// get the audio in base64 using the tracks microservice by id
				result2 := httpRequest(w, "http://localhost:3000/tracks/"+newId, "GET", nil) // will write the corresponding status code

				if result2 == nil {
					return
				}

				audio2 := result2["Audio"] // audio in base64 of the full track

				track := Track{Audio: audio2.(string)}

				w.WriteHeader(200)               // 200 OK
				json.NewEncoder(w).Encode(track) // return the full audio as a json
				return

			}
		}
		w.WriteHeader(400) // Bad Request
		return

	}

	w.WriteHeader(500) // Internal Server Error

}

func Router() http.Handler {
	r := mux.NewRouter()
	/* recognizes the track fragment audio and returns the full track audio */
	r.HandleFunc("/cooltown", getTrack).Methods("POST")
	return r
}
