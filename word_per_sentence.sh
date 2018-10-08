#!/bin/bash -e

req='{"somekey":"This is a sentence. But is this? I KNOW THIS IS!" }'

curl \
    -k \
    -X POST \
    -H "Content-type: application/json" \
    -H "Accept: application/json" \
    -d "$req" \
    "http://localhost:9090/word_count_per_sentence"
