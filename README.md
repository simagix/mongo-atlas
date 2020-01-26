# MongoDB Atlas API
Invoke MongoDB Atlas API using Golang.

## Supported Functions
- Cluster summary
- Start/Pause cluster
- Download Logs
- Download FTDC (TODO)
- Billing/Invoices (TODO)

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

```
matlas --loginfo "atlas://{pub_key}:{pri_key}@{group}/{cluster}/[/yyyy-mm-dd]"
```
