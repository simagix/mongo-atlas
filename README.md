# MongoDB Atlas API
Invoke MongoDB Atlas API using Golang.

## Get clusters summary
Set up environment variables

```
export ATLAS_GROUP=5b75e8da0bd66b7ea1abc123
export ATLAS_PUB=ABCDEFGH
export ATLAS_PRI=d936a93e-ba0a-4fd3-be73-d075f7321cba
export ATLAS_AUTH="${ATLAS_PUB}:${ATLAS_PRI}"
```

```
mongo-atlas --clusters ${ATLAS_AUTH}
```

## Download log files

```
mongo-atlas --clusters ${ATLAS_AUTH}@${ATLAS_GROUP}[/yyyy-mm-dd]
```
