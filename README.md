# Intel 8080 CPU emulator in Go

An Intel 8080 CPU emulator, written in Go.  This project also uses [my Intel 8080 CPU assembler](https://github.com/lukepeterson/go8080assembler), especially for running tests.

[![Tests](https://github.com/lukepeterson/go8080cpu/actions/workflows/go.yml/badge.svg)](https://github.com/lukepeterson/go8080cpu/actions/workflows/go.yml)
![Go Report Card](https://goreportcard.com/badge/github.com/lukepeterson/go8080cpu)
![GitHub release](https://img.shields.io/github/v/release/lukepeterson/go8080cpu)

![Running some INR and DCR](./images/running.gif)

## Features
- :white_check_mark: Registers
- :white_check_mark: Memory
- :white_check_mark: Fetch/decode/execute cycle
- :white_check_mark: [Write an assembler](https://github.com/lukepeterson/go8080assembler)

## In progress :warning:
- Increment and decrement
- Test framework

## TODO
- :x: Move, load and store
- :x: Stack operations
- :x: Jump
- :x: Call
- :x: Return
- :x: Restart
- :x: Increment/decrement
- :x: Add
- :x: Subtract
- :x: Logical
- :x: Rotate
- :x: Specials
- :x: Input/output
- :x: Control
- :x: Interrupts
- :x: Fix inteface warnings on read/write memory functions
- :x: Decide what I want to do with the ALU

# Running tests

Run `go test ./...`.