#! /bin/bash

echo "=> Test --info"
go run atlas.go --info "atlas://${ATLAS_AUTH}"

echo "=> Test --loginfo"
go run atlas.go --loginfo "atlas://${ATLAS_AUTH}@5b75e8da0bd66b7ea13217a1/keyhole/$(date '+%Y-%m-%d')"
ls -l mongodb.*.gz
rm -f mongodb.*.gz
