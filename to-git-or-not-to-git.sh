# !/bin/bash
jq -r '.[] | select(.id == 170) | .name, .power, .gender' superhero