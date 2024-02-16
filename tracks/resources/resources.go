package resources

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tracks/repository"

	"github.com/gorilla/mux"
)

// update or insert a track
func updateTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get the id parameter from the url
	id := vars["id"]
	var t repository.Track
	// the body contains a track in json {Id: id, Audio: audio}
	if err := json.NewDecoder(r.Body).Decode(&t); err == nil {
		if id == t.Id {
			// update the Tracks table or Insert a new one
			if n := repository.Update(t); n > 0 {
				fmt.Println("track updated")
				w.WriteHeader(204) /* No Content */
				return
			} else if t, n := repository.Insert(t); n > 0 {
				fmt.Println("track inserted: " + t.Id)
				w.WriteHeader(201) /* Created */
				return
			} else {
				w.WriteHeader(500) /* Internal Server Error */
				return
			}
		}
	}

	w.WriteHeader(400) /* Bad Request */
}

// list all tracks
func listTracks(w http.ResponseWriter, r *http.Request) {
	if tracks, count := repository.List(); count > 0 {
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(tracks) // show the list of tracks as a list of ids, or empty list if no tracks
	} else if count == 0 {
		w.WriteHeader(200) /* OK */
		// create an empty array
        data := []interface{}{}
		// return an empty array
		json.NewEncoder(w).Encode(data)

	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

// read a track by id
func readTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get the id from the url
	id := vars["id"]
	// find the track from the Tracks table
	if t, n := repository.Read(id); n > 0 {
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(t)
	} else if n == 0 {
		w.WriteHeader(404) /* Not Found */
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func deleteTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get the id from the url
	id := vars["id"]
	// delete a row from the Tracks table
	if n := repository.Delete(id); n > 0 {
		w.WriteHeader(204) /* No content */
	} else if n == 0 {
		w.WriteHeader(404) /* Not Found */
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* Insert Tracks*/
	r.HandleFunc("/tracks/{id}", updateTrack).Methods("PUT")
	/* Read Track */
	r.HandleFunc("/tracks/{id}", readTrack).Methods("GET")
	/* List Tracks */
	r.HandleFunc("/tracks", listTracks).Methods("GET")
	/* Delete Track */
	r.HandleFunc("/tracks/{id}", deleteTrack).Methods("DELETE")
	return r
}
