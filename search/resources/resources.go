package resources

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

// used to return the id of the recognized track
type Track struct {
	Id string
}

// >>> CHANGE THE API KEY HERE <<< (import from config file, etc)
const (
	API_KEY = "test"
)

// returns the title of a track by audio
func searchTrack(w http.ResponseWriter, r *http.Request) {
	t := map[string]interface{}{} // t will store the request data (the audio of the track segment)
	if err := json.NewDecoder(r.Body).Decode(&t); err == nil {
		// audio is the audio bas64 encoded
		if audio, ok := (t["Audio"].(string)); ok { // check if audio is a string
			if t["Audio"].(string) != "" { // the audio string is not empty

				// audd parameters
				data := url.Values{
					"audio":     {audio}, // send the audio in base64
					"api_token": {API_KEY},
				}

				// find the track using the audd api
				if response, err := http.PostForm("https://api.audd.io/", data); err == nil {
					defer response.Body.Close()

					// get the body of the response from audd
					if body, err := io.ReadAll(response.Body); err == nil {

						// the response body may look different
						result1 := map[string]interface{}{}            // as map(string => value)
						result2 := map[string]map[string]interface{}{} // as string => map (string => value)

						json.Unmarshal(body, &result1)
						json.Unmarshal(body, &result2)

						status := result1["status"]
						fmt.Println("status: ", status)

						if status == "success" {
							if result1["result"] == nil {
								// successful request but track not found
								fmt.Println("Audd: 404 Track Not Found")
								w.WriteHeader(404) // Not Found
								return
							} else {
								// get the title of the recognized track
								title := result2["result"]["title"].(string)
								track := Track{Id: title}

								// return the title of the track as a json
								w.WriteHeader(200) // OK
								fmt.Println("title: ", title)
								json.NewEncoder(w).Encode(track)
								return
							}
						} else {
							// print the audd api error code
							fmt.Println("error_code: ", result2["error"]["error_code"])

						}
					}

				}

				w.WriteHeader(500) // Internal Server Error
				return

			}

		}
	}

	w.WriteHeader(400) // Bad Request
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* Return the Track title/id */
	r.HandleFunc("/search", searchTrack).Methods("POST")
	return r
}
