#!/bin/bash
curl -s https://platform.zone01.gr/assets/superhero/all.json \
| jq --arg family_id "$HERO_ID" -r '.[] | select(.id == ($family_id | tonumber)) | .connections.relatives
| gsub("\n"; "\\n")
| gsub("\r"; "")'