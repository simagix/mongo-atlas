# MongoDB Atlas API

Invoke MongoDB Atlas API using Golang.  Source codes were merged back to [Keyhole](https://github.com/simagix/keyhole).

## Supported Functions

- Clusters summary
- Create/Terminate a cluster
- Start/Pause a cluster
- Download Logs
- Display Alerts Settings
- Add Recommended Alerts

## Get clusters summary

```bash
matlas --info "atlas://{pub_key}:{pri_key}"
```

## Pause a Cluster

```bash
matlas --pause "atlas://{pub_key}:{pri_key}@{group}/{cluster}"
```

## Resume a Cluster

```bash
matlas --resume "atlas://{pub_key}:{pri_key}@{group}/{cluster}"
```

## Download log files

By default, without specifying startDate nor endDate, it downloads today's logs.

```bash
matlas --loginfo "atlas://{pub_key}:{pri_key}@{group}/{cluster}[?startDate=yyyy-mm-dd&endDate=yyyy-mm-dd]"
```

## Create a Cluster

```bash
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

```bash
matlas --request DELETE "{pub_key}:{pri_key}@{group}/{cluster}"
```

## Get All Alerts

```bash
matlas --alerts "{pub_key}:{pri_key}@{group}"
```

## Configure Recommended Alerts

```bash
matlas --addAlerts conf/alerts.json "{pub_key}:{pri_key}@{group}"
```

## Disclaimer

This software is not supported by MongoDB, Inc. under any of their commercial support subscriptions or otherwise. Any usage of keyhole is at your own risk. Bug reports, feature requests and questions can be posted in the Issues section on GitHub.
