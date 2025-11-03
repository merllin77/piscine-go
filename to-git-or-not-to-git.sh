# !/bin/bash
wget -O superhero https://platform.zone01.gr/assets/superhero/all.json
jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender' superhero