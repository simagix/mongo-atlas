#! /bin/bash
# Copyright 2019 Kuei-chun Chen. All rights reserved.

echo "=> Test --request POST"
go run atlas.go --request POST "${ATLAS_AUTH}@${ATLAS_GROUP}" '
{
    "name": "matlas",
    "numShards": 1,
    "providerSettings": {
        "providerName": "AWS",
        "instanceSizeName": "M10",
        "regionName": "US_EAST_1"
    },
    "clusterType" : "REPLICASET",
    "backupEnabled": false
}'

echo "=> Test --info"
go run atlas.go --info "${ATLAS_AUTH}"

echo "=> Test --loginfo"
go run atlas.go --loginfo "${ATLAS_AUTH}@${ATLAS_GROUP}/${ATLAS_CLUSTER}/$(date '+%Y-%m-%d')"
ls -l mongodb.*.gz
rm -f mongodb.*.gz

echo "=> Test --pause"
go run atlas.go --pause "${ATLAS_AUTH}@${ATLAS_GROUP}/${ATLAS_CLUSTER}"

echo "=> Test --resume"
go run atlas.go --resume "${ATLAS_AUTH}@${ATLAS_GROUP}/${ATLAS_CLUSTER}"

echo "=> Test --request DELETE"
# go run atlas.go --request DELETE "${ATLAS_AUTH}@${ATLAS_GROUP}/matlas"

