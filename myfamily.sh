#!/bin/bash
curl -s https://platform.zone01.gr/assets/superhero/all.json \
| jq -r --arg family_id "$HERO_ID" '.[] | select(.id == ($family_id | tonumber)) | .connections.relatives' | tr ',' ';'