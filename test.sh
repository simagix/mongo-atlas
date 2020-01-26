#! /bin/bash
# Copyright 2019 Kuei-chun Chen. All rights reserved.

echo "=> Test --info"
go run atlas.go --info "atlas://${ATLAS_AUTH}"

echo "=> Test --loginfo"
go run atlas.go --loginfo "atlas://${ATLAS_AUTH}@${ATLAS_GROUP}/${ATLAS_CLUSTER}/$(date '+%Y-%m-%d')"
ls -l mongodb.*.gz
rm -f mongodb.*.gz

echo "=> Test --pause"
go run atlas.go --pause "atlas://${ATLAS_AUTH}@${ATLAS_GROUP}/${ATLAS_CLUSTER}"

echo "=> Test --resume"
go run atlas.go --resume "atlas://${ATLAS_AUTH}@${ATLAS_GROUP}/${ATLAS_CLUSTER}"

