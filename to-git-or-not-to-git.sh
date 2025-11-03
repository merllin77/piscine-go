#!/bin/bash
curl -s -k "https://platform.zone01.gr/assets/superhero/all.json" -o superhero
jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender' superhero