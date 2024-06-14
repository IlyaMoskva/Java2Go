# Java2Go
[![Go](https://github.com/IlyaMoskva/Java2Go/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/IlyaMoskva/Java2Go/actions/workflows/go.yml)

Go Transition practice for Java developers. [Link](https://www.linkedin.com/learning/transition-from-java-to-go)

[Certificate](https://www.linkedin.com/learning/certificates/fb05cfed0c20b5e38cf7db116da3c774ed0ec4c56a05d8c6740a0b775af72234) for this course.

# Main Application
## Vacation Planner
CLI with parameters to plna vacation. 
Accepts parameters (OR-based)
```sh
-beach:   beach available            boolean
-ski:     mountains available        boolean
-month N: month of vacation [1..12]  int, required
-name:    resort contains name       string
```
To run:
```sh
go run ./hello.go -beach -month 8 -name York
```

## Goroutines concurrency demo
```sh
go run ./conc/main.go
```

## Generics demo
```sh
go run ./generics/main.go
```

## Run tests
```sh
go test ./... -v
```
