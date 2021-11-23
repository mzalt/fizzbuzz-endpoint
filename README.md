# fizzbuzz-api

Project of fizzbuzz endpoint. [doc swagger](https://app.swaggerhub.com/apis/mzalt/fizzbuzz/1.0.0-oas3#/leboncoin/myFizzBuzz)

## Prerequisites

* [Golang](https://golang.org/doc/)

## Configuration

| Variable Name    | Description                                              | Default value  |
|------------------|----------------------------------------------------------|----------------|
| HTTP_PORT        | Port of the http server                                  | 8001           |
| TOP_USED_REQUEST | Number of top used request values that we would monitor  | 5              |


## Run
    
run the http server: `go run cmd/server/main.go`

## API

this endpoint takes five parameters (int1, int2, limit of type integer) and (str1, str2 of type string) to return a string of numbers from 1 to limit, where all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2 and all multiples of int1 and int2 are replaced by str1str2.

Example: http://localhost:8001/mzalt/leboncoin/fizzbuzz/?int1=2&int2=4&limit=10&str1=fizz&str2=buzz

```
int1=2, int2=4, limit=10, str1=fizz, str2=buzz
output: "1, fizz, 3, fizzbuzz, 5, fizz, 7, fizzbuzz, 9, fizz"
```

## METRICS

Shows the top used requests and parameters.

http://localhost:8001/metrics 
