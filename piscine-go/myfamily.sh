: ${HERO_ID:=1}
curl -s "https://platform.zone01.gr/assets/superhero/all.json" | \
jq --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives' | \
sed 's/"//g' | tr -d '\n'
