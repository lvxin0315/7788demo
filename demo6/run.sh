#!/bin/bash

go run main.go -keyword="elasticsearch" &
go run main.go -keyword="flink" &
go run main.go -keyword="golang" &
go run main.go -keyword="spring boot"