# Intel 8080 CPU emulator in Go

An Intel 8080 CPU emulator, written in Go.  This project also uses [my Intel 8080 CPU assembler](https://github.com/lukepeterson/go8080assembler), especially for running tests.

## Progress

- :white_check_mark: Base functionality: 
  - :white_check_mark: Registers
  - :white_check_mark: Memory
  - :white_check_mark: Fetch/decode/execute cycle
  - :warning: Test framework
  - :x: Input/output
  - :x: Interrupts
- :x: Instruction set groups:
  - :x: Move, load and store
  - :x: Stack operations
  - :x: Jump
  - :x: Call
  - :x: Return
  - :x: Restart
  - :warning: Increment and decrement
  - :x: Add
  - :x: Subtract
  - :x: Logical
  - :x: Rotate
  - :x: Specials
  - :x: Input/output
  - :x: Control
- :warning: Write an assembler
- :warning: Group functions a little better
- :warning: Fix inteface warnings on read/write memory functions

# Running tests

Run `go test ./...`.