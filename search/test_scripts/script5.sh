# returns the title/id of the track given a track fragment

#!/bin/sh
ID="~Everybody+(Backstreet's+Back)+(Radio+Edit)"
AUDIO=`base64 -i "example_track_segments/$ID".wav`
RESOURCE=localhost:3001/search
echo "{ \"Audio\":\"$AUDIO\" }" > input
curl -v -X POST -d @input $RESOURCE
