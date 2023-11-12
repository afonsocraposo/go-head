# Head Challenge in Go

## Overview

This project implements the "Head" challenge from [Coding Challenges](https://codingchallenges.substack.com/p/coding-challenge-33-head), a task that focuses on creating a program to display the first N lines of a given text file. Written in Go, this project offers a lightweight and efficient solution for file reading and data processing.

## Installation

To install this project, ensure you have Go installed on your system. You can download it from [Go's official website](https://golang.org/dl/).

1. Clone the repository:
   ```
   git clone https://github.com/afonsocraposo/go-head
   ```
2. Navigate to the project directory:
   ```
   cd go-head
   ```

## Usage

To use the program, run the following command in your terminal:

```
go run head.go [file name] ...
```

Usage of `head.go`:
```
Usage of head.go:
  -c int
        display first c bytes
  -n int
        display first n lines (default 10)
```

## Example

### Print 10 lines of 1 file

```
go run head.go -n 10 test.txt
```

### Print 10 bytes of 1 file

```
go run head.go -c 10 test.txt
```

### Print 10 lines of each file

```
go run head.go -n 10 test.txt test2.txt
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
