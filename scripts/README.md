# Download FTDC from Atlas
## Required Info
To download FTDC files, the following are required:

- Altas Credential in **{pub_key}:{pri_key}** format
- Group ID
- Resource Name

### Atlas Credential
Obtain an API key with **Project Read Only** permission. Export **ATLAS_AUTH** variable.

```
export ATLAS_AUTH="{pub_key}:{pri_key}"
```

### Group ID
Locate your Group ID.  You can use [keyhole](https://github.com/simagix/keyhole) to obtain it.  For example:

```
keyhole --info "atlas://${ATLAS_AUTH}"

- Group: {GROUP_ID}
  - cluster name: Cluster0
    - 4.2.2, REPLICASET, mongodb+srv://Cluster0-jgtm2.mongodb.net
    - Hosts:
      - Cluster0-shard-00-00-jgtm2.mongodb.net (REPLICA_SECONDARY)
      - Cluster0-shard-00-01-jgtm2.mongodb.net (REPLICA_SECONDARY)
      - Cluster0-shard-00-02-jgtm2.mongodb.net (REPLICA_PRIMARY)
```

Followed by the export **ATLAS_GROUP** command.

```
export ATLAS_GROUP="{GROUP_ID}"
```

## Download FTDC
### Replica Set
Resource ID for a replica set is the set name.  If not sure, use keyhole to find the set name if needed.

```
keyhole --info "mongodb+srv://{user}:{password}@{hostname}/" | jq '.repl.setName'
```

For example:

```
ATLAS_RESOURCE=Cluster0-shard-0 get_ftdc.sh
```

### Sharded Cluster
Resource ID for a sharded cluster is the cluster name.  For example:

```
ATLAS_RESOURCE=Cluster0 get_ftdc.sh shard
```
