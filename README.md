# Intel 8080 CPU emulator in Go

This is my Intel 8080 CPU emulator, written in Go.  I built it because computers are awesome, and it was a great excuse to also write [an Intel 8080 CPU assembler](https://github.com/lukepeterson/go8080assembler).

[![Tests](https://github.com/lukepeterson/go8080cpu/actions/workflows/go.yml/badge.svg)](https://github.com/lukepeterson/go8080cpu/actions/workflows/go.yml)
![Go Report Card](https://goreportcard.com/badge/github.com/lukepeterson/go8080cpu)
![GitHub release](https://img.shields.io/github/v/release/lukepeterson/go8080cpu)

![Running some INR and DCR](./images/running.gif)

## Features
- :white_check_mark: Registers
- :white_check_mark: Memory
- :white_check_mark: Fetch/decode/execute cycle
- :white_check_mark: [Assembler support](https://github.com/lukepeterson/go8080assembler)

## Instructions supported
- :white_check_mark: Move, load and store (84 instructions)
- :white_check_mark: Stack operations (13 instructions)
- :white_check_mark: Jump (10 instructions)
- :white_check_mark: Call (9 instructions)
- :white_check_mark: Return (9 instructions)
- :white_check_mark: Restart (8 instructions)
- :white_check_mark: Increment and decrement (22 instructions)
- :white_check_mark: Add (22 instructions)
- :white_check_mark: Subtract (18 instructions)
- :white_check_mark: Logical (36 instructions)
- :white_check_mark: Rotate (4 instructions)
- :white_check_mark: Specials (4 instructions)
- :white_check_mark: Input/output (2 instructions)
- :white_check_mark: Control (4 instructions)

## Future enhancements
- Replace the memory locations in tests with labels once the assembler supports them.

# Running tests
Run `go test ./...`.
