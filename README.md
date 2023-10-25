# Sorting Program

This Go program demonstrates sorting an array of integers using both in-memory and external sorting methods.

## Prerequisites

Before running this program, make sure you have Go installed on your system. If not, you can download and install it from the official website: [Go Downloads](https://golang.org/dl/)

## Installation

```
go mod tidy  
```


## Usage

To run the sorting program, use the following command:
```
go run cmd/lab_1/main.go [options]
```

## Options

- -input: Input file name (default: "input.txt")
- -output: Output file name (default: "output.txt")
- -external: Use external sort (default: false-in-memory sort)

## Testing
```
go test ./...
```
## Generating

There was were created two scripts for generating large files and deleting them.

- Use for delete prev file and create new one
```
go generate ./...
```

- Also there is posible to execute specified script
```
go run ./pkg/scripts/gen_rand_num_file/main.go
```
There is also possible to use command line params

## Note for laboratory work
```
.internal/domain/lab_1/ip-z21_stykhun_marian_lr1_2023.docx
```