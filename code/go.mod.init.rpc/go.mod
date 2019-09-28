module microservices

go 1.13

require handler v0.0.0

replace handler => ./handler

require pkg/rpc v0.0.0

replace pkg/rpc => ./pkg/rpc

require github.com/gorilla/mux v1.7.3 // indirect
