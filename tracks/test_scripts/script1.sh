# creates/updates a track

#!/bin/sh
ID="Everybody+(Backstreet's+Back)+(Radio+Edit)"
# ID="test"
ESCAPED=`perl -e "use URI::Escape; print uri_escape(\"$ID\")"`
AUDIO=`base64 -i "example_tracks/$ID".wav`
RESOURCE=localhost:3000/tracks/$ESCAPED
echo "{ \"Id\":\"$ID\", \"Audio\":\"$AUDIO\" }" > input
curl -v -X PUT -d @input $RESOURCE
