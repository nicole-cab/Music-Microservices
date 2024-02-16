# Music Microservices App

An audio recognition app which comprises a set of microservices designed to handle various functions related to music tracks similar to Shazam.

## Features:

#### Tracks Microservice:

- Create a Track: Insert a track to the database.
- List all Tracks: Show all tracks in the database.
- Read a Track: Read a track from the database.
- Delete a Track: Remove a track from the database.

#### Search Microservice:

- Search a Track: Given a track segment use the AudD Music Recognition API to return the title of the track.

#### Cooltown Microservice:

- Retrieve A Full Track: Given a track segment, retrieve the full track from the database. It uses the Search Microservice to recognize the track (gets the track id/title), then it uses the Tracks Microservice to retrieve the full track.

## Technologies Used:

- Go
- SQLite
- REST APIs

## Examples:

#### Tracks Microservice:

- Create a Track (script1.sh)
![Create Track Screenshot](https://github.com/nicole-cab/music-microservices/blob/main/screenshots/create_track.png?raw=true)

- List all Tracks (script2.sh)
![List Tracks Screenshot](https://github.com/nicole-cab/music-microservices/blob/main/screenshots/list_tracks.png?raw=true)

- Read a Track (script3.sh)
![Read Track Screenshot](https://github.com/nicole-cab/music-microservices/blob/main/screenshots/read_track.png?raw=true)

- Delete a Track (script4.sh)
![Delete Track Screenshot](https://github.com/nicole-cab/music-microservices/blob/main/screenshots/delete_track.png?raw=true)

#### Search Microservice:

- Search a Track (script5.sh)
![ Search Track Screenshot](https://github.com/nicole-cab/music-microservices/blob/main/screenshots/search_track.png?raw=true)

#### Cooltown Microservice: (script6.sh)

- Retrieve A Full Track
![Retrueve Full Track Screenshot](https://github.com/nicole-cab/music-microservices/blob/main/screenshots/recognize_track.png?raw=true)
