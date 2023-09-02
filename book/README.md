# book

This is https://github.com/googleapis/gapic-generator-go/issues/1372

```
protoc -Iproto --go_opt=module=github.com/julieqiu/scratch/book --go_out=. --go-grpc_opt=module=github.com/julieqiu/scratch/book --go-grpc_out=. proto/book.proto
```
