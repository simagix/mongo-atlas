#! /bin/bash
# Copyright 2019 Kuei-chun Chen. All rights reserved.

if [ "${ATLAS_AUTH}" == "" ];
    echo "export ATLAS_CLUSTER={pub_key}:{pri_key}"
    exit
fi

echo "=> Test --info"
go run atlas.go --info "${ATLAS_AUTH}"

if [ "${ATLAS_CLUSTER}" == "" ];
    export ATLAS_CLUSTER=Cluster0
fi

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
    "mongoDBMajorVersion" : "4.2",
    "backupEnabled": false
}'

echo "=> Test --loginfo"
go run atlas.go --loginfo "${ATLAS_AUTH}@${ATLAS_GROUP}/${ATLAS_CLUSTER}/$(date '+%Y-%m-%d')"
ls -l mongodb.*.gz
rm -f mongodb.*.gz

echo "=> Test --ftdc REPLICASET"
go run atlas.go --ftdc "${ATLAS_AUTH}@${ATLAS_GROUP}/${ATLAS_CLUSTER}-shard-0-node-0"
ls -l ${ATLAS_CLUSTER}-shard-0-node-0-diagnostic.tar.gz
rm -f ${ATLAS_CLUSTER}-shard-0-node-0-diagnostic.*.gz

echo "=> Test --ftdc REPLICASET"
go run atlas.go --ftdc "${ATLAS_AUTH}@${ATLAS_GROUP}/${ATLAS_CLUSTER}-shard-0"
ls -l ${ATLAS_CLUSTER}-shard-0-diagnostic.tar.gz
rm -f ${ATLAS_CLUSTER}-shard-0-diagnostic.*.gz

echo "=> Test --ftdc CLUSTER"
go run atlas.go --ftdc "${ATLAS_AUTH}@${ATLAS_GROUP}/${ATLAS_CLUSTER}?size=5000000"
ls -l ${ATLAS_CLUSTER}-diagnostic.tar.gz
rm -f ${ATLAS_CLUSTER}-diagnostic.*.gz

echo "=> Test --pause"
go run atlas.go --pause "${ATLAS_AUTH}@${ATLAS_GROUP}/${ATLAS_CLUSTER}"

echo "=> Test --resume"
go run atlas.go --resume "${ATLAS_AUTH}@${ATLAS_GROUP}/${ATLAS_CLUSTER}"

echo "=> Test --request DELETE"
# go run atlas.go --request DELETE "${ATLAS_AUTH}@${ATLAS_GROUP}/${ATLAS_CLUSTER}"

