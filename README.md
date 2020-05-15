# Phi Skills Generic gRPC API Server for Go

| **Homepage** | [https://phiskills.com][0]        |
| ------------ | --------------------------------- | 
| **GitHub**   | [https://github.com/phiskills][1] |

## Overview

This project contains the Go module to create a generic **gRPC API Server**.  

## Installation

```bash
go get github.com/phiskills/grpc-api.go
```

## Creating the server

```go
package main
import "github.com/phiskills/grpc-api.go"

api := grpc.New('My API')
xxx.RegisterXxxServer(api.Server(), &xxxServer{})
api.Start()
```
For more details, see [gRPC Basics - Go: Creating the server][10].

[0]: https://phiskills.com
[1]: https://github.com/phiskills
[10]: https://www.grpc.io/docs/tutorials/basic/go/#server
