### Hexago(lang)nal Architecture

![](image/arch.webp)

- hexagonal architecture? architectural pattern, just like microservice, clean-arch (the great uncle bob), event-driven, layered, ddd, mvc(bloat)

<pre>
└── Go Apps
   ├── cmd
   │   └── main.go
   ├── go.mod
   ├── go.sum
   └── internal
       ├── adapters
       │   ├── handler
       │   │   └── http.go
       │   └── repository
       │       ├── postgres.go
       │       └── redis.go
       └── core
           ├── domain
           │   └── model.go
           ├── ports
           │   └── ports.go
           └── services
               └── services.go
</pre>
