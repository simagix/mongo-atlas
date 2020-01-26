# MongoDB Atlas API
Invoke MongoDB Atlas API using Golang.

## Supported Functions
- Clusters summary
- Create/Terminate a cluster
- Start/Pause a cluster
- Download FTDC
- Download Logs

## Get clusters summary

```
matlas --info "atlas://{pub_key}:{pri_key}"
```

## Pause a Cluster

```
matlas --pause "atlas://{pub_key}:{pri_key}@{group}/{cluster}"
```

## Resume a Cluster

```
matlas --resume "atlas://{pub_key}:{pri_key}@{group}/{cluster}"
```

## Download log files
By default, without specifying startDate nor endDate, it downloads today's logs.

```
matlas --loginfo "atlas://{pub_key}:{pri_key}@{group}/{cluster}[?startDate=yyyy-mm-dd&endDate=yyyy-mm-dd]"
```

## Create a Cluster

```
matlas --request POST "{pub_key}:{pri_key}@{group}" '
{
    "name": "Cluster0",
    "numShards": 1,
    "providerSettings": {
        "providerName": "AWS",
        "instanceSizeName": "M10",
        "regionName": "US_EAST_1"
    },
    "clusterType" : "REPLICASET",
    "backupEnabled": false
}'
```

## Terminate a Cluster

```
matlas --request DELETE "{pub_key}:{pri_key}@{group}/{cluster}"
```
