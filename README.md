# etcd Parser

A Go package that converts etcd-style key-value pairs into nested map structures.

## Usage

```go
import parser "github.com/gokpm/go-etcd-parser"

// Convert etcd KeyValue pairs to nested map
kvPairs := []*mvccpb.KeyValue{
    {Key: []byte("app/database/host"), Value: []byte("localhost")},
    {Key: []byte("app/database/port"), Value: []byte("5432")},
    {Key: []byte("app/name"), Value: []byte("myapp")},
}

result := parser.Parse(kvPairs)
// Returns:
// {
//   "app": {
//     "database": {
//       "host": "localhost",
//       "port": "5432"
//     },
//     "name": "myapp"
//   }
// }
```

## Functions

- `Parse(kvPairs []*mvccpb.KeyValue) map[string]interface{}` - Converts flat etcd key-value pairs into a nested map structure using "/" as the delimiter.

## Dependencies

- `go.etcd.io/etcd/api/v3/mvccpb`