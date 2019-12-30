#! /bin/bash
if [ "$ATLAS_AUTH" == "" ]; then
    echo "export ATLAS_AUTH={PUB_KEY}:{PRIVATE_KEY}"
    exit
fi

if [ "$ATLAS_GROUP" == "" ]; then
    echo "ATLAS_GROUP required"
    echo "export ATLAS_GROUP={GROUP_ID}"
    exit
fi

if [ "$ATLAS_RESOURCE" == "" ]; then
    echo "ATLAS_RESOURCE required, provide a replica set resource"
    echo "export ATLAS_RESOURCE=cluster-shard-x"
    exit
fi

rtn=$(curl -s --user "${ATLAS_AUTH}" --digest \
--header "Accept: application/json" \
--header "Content-Type: application/json" \
--request POST "https://cloud.mongodb.com/api/atlas/v1.0/groups/${ATLAS_GROUP}/logCollectionJobs" \
--data "
{
\"resourceType\": \"REPLICASET\",
\"resourceName\": \"$ATLAS_RESOURCE\",
\"redacted\": true,
\"sizeRequestedPerFileBytes\": 100000000,
\"logTypes\": [
\"FTDC\"
]
}" )

JOB_ID=$(echo $rtn | jq -r '.id')

if [ "$JOB_ID" == null ]; then
    echo $rtn | jq -r '.errorCode'
    exit
fi

STATUS="IN_PROGRESS"
while [ "$STATUS" == "IN_PROGRESS" ]; do
    echo "status: $STATUS"
    sleep 10
    STATUS=$(curl -s --user "${ATLAS_AUTH}" --digest \
--header "Accept: application/json" \
--header "Content-Type: application/json" \
--request GET "https://cloud.mongodb.com/api/atlas/v1.0/groups/${ATLAS_GROUP}/logCollectionJobs/${JOB_ID}?verbose=true&pretty=true" | jq -r '.status')
done

curl --user "${ATLAS_AUTH}" --digest \
--header "Accept: application/gzip" \
--header "Content-Type: application/gzip" \
--output diagnostic.tar.gz \
--request GET "https://cloud.mongodb.com/api/atlas/v1.0/groups/${ATLAS_GROUP}/logCollectionJobs/${JOB_ID}/download"
