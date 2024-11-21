# Overview

![dt_example_720](https://github.com/user-attachments/assets/98da69fb-1f69-44fe-b613-f4f4f6dda1d6)

Dicthesaurus is a dictionary and thesaurus CLI program built only with Go standard packages. 

Currently, it makes use of the [Merriam Webster's API](https://dictionaryapi.dev/).

## Motivation

It is my intent to replace my love of OSX's built-in dictionary lookup through the spotlight tool now that I'm solely using Linux, and I've found there's never a day or time where I don't have an active terminal ready for a quick search!

## Instructions

### Installation

#### Bash script
This script will check to see if Go is installed and updated on your system, download this package, build it with Go, and store it in /usr/local/bin.

TBD

#### Building via Go
0. Create an account and get dictionary and thesaurus keys at [Merriam Webster](https://dictionaryapi.com/)
1. Clone and navigate to the root of this package
2. Create a `.env` file with both ENV variables:
    ```
    MERRIAM_WEBSTER_DICTIONARY_API_KEY="..."
    MERRIAM_WEBSTER_THESAURUS_API_KEY="..."
    ```
3. Run `$ go build -o [whatever you want the binary to be - dt]`
4. Move the newly created binary from the root of your project to a bin in your $PATH
    1. I believe MacOS recommends `/usr/local/bin`, which is for user-installed binaries
5. Restart/source your consoles
6. `$ [binary] taco` should return definitions for "taco"

### Usage

Via your terminal:

`$ dt <word> [FLAGS]`

Flags can be used in either format: `-d, --d`

### Command Flags

| Flag | Description |
|--|--|
| `h, help` | help |

### Subcommand Flags

| Flag | Description |
|--|--|
| `e` | Display the word in a sentence |
| `t` | Include thesaurus results |
| `ud` | Also query Urban Dictionary for results |

## Use Cases

- [x] Defined command flags
- [x] Undefined command flags
- [x] Missing subcommand
- [ ] Defined subcommand flags
- [x] Undefined subcommand flags

## Version 1 Review

- The [Free Dictionary API](https://dictionaryapi.dev) does not provide adequate responses:
    - Many words or "psuedo" words are have not returned as having a definition, such as "api" "look"
    - Response can contain an array of root level JSON objects while each JSON object can also contain an array of definitions for an array of parts of speech, which is too much confusion to work through to organize if the result quality isn't what I'm looking for
    - The antonym and synonym fields are blank 99% of the time, and there isn't a "free thesaurus API" to match
- Thesaurus flag does not work given the antonym and synonym fields are always empty in the API responses

## Version 2 Review
- Becuase it was a personal project, I was a little too relaxed with making small commits, and sequentially performing really small reviews and PRs, which would have taken less time, mental energy, and taught me the SOLID principles more quickly, which I hope would indicate all subsequent PRs would require less effort to review because I was writing better code...
    - Lesson: Small tasks, commit often, PR often, review often (AKA small pieces)
- While reviewing similarity of the output formatting between the dictionary and thesaurus I realized I should have came up with a presentation layout beforehand; however, this is what I happily came up with based on a few basic text-based design principles (white space, weight, case)
    ```text
    \n
    "Searching [dictionary/thesaurus] for [subcommand] ..."
    \n
    \n
    \n
    [Part of Speech]
    \n
    \n
    [Value]
    \n
    [Value]
    \n
    \n
    \n
    [Part of Speech]
    ...repeat values and parts of speech until end
    --- end of output ---
    \n
    ```
- I'm incredibly happy with the difference in my code from V1 to V2 - the SOLID principles and refactoring presentation and input to be almost stateless was a great decision

## What I learned

- What Go considers a package
    - Basically, this is a directory with go file(s)
    - Each .go file in a directory, outside of root, should have package "[directoryName]" to say, "I'm a part of this package and when you import said package, you're also importing me"
- How Exported names are available or made available throughout a package containing other packages
    - Within the same package, they're available without a namespace
    - Within packages they're imported, they're available under their package name (directory name) or a given namespace before the import path
- Standard "flag" package
    - Great at parsing flags, but requires custom coverage for implementating subcommands
    - Usage errors (flag parsing) are very easy to overwrite with custom functions
    - Rather shallow package in terms of only working with flags and not commands
    - Subcommands will need to be parsed separately
    - Cobra seems to be the standard external library for creating CLI apps with Go, including the GitHub CLI, which was my inspiration for Dicthesaurus
- It's very easy to parse CLI input via `os.Args`:
    - The first argument is always going to be your command
    - The second argument is either a subcommand or a command flag
        - `strings.HasPrefix(os.Arg[1], "-")` == flag
        - `!strings.HasPrefix(os.Arg[1], "-")` == subcommand
- Handling JSON is a bit odd in Go because Go data structures aren't built so similar, such as a JavaScript object.
    - "encoding/json" provides two functions for encoding and decoding JSON: `Marshal` and `Unmarshal` (WTF is up with those names)
    - `Unmarshal` parses the given []byte into a pointer.
        - If the pointer is an empty interface, `interface{}`, the `Unmarshal` will fill the interface with a structure matching the JSON's structure
        - If the pointer is a struct, `Unmarshal` will match the JSON properties with the struct's properties
            - By providing a typed structure to the pointer for `Unmarshal`, you can predefine the JSON's structure into a more friendly structure for you application
            - You can even provide "tags" to your struct properties which will map specific JSON properties to specific struct properties, even if the names don't match, without writing complicated mapping logic
            - By typing the pointer structure, you can resuse those types throughout your application to make clear what data is being passed around and in what form
- `defer` tells a statement/function to defer execution until the end of the parent closure, so if when you defer closing a response body (`defer response.Body.Close()`) immediately after handling the error for a GET request, the response (io.Reader) is not actually closed until right before the closing brace of the function calling the defer.
- "fmt"
    - `Fprintf` provides a third (the first) argument for where the output should go to (such as stdout or stderr)
    - `Sprintf` allows you to use `Printf`'s formatting to create a string that can be returned or stored but not immediately printed to stdout or stderr
- `os.Exit(code int)` exits the current program to exit with the given status code (Exist status X in console)
    - 0 == success
    - Non-zero == error
    - defer'd functions are not run
- When building a CLI, some things require knowledge of which kernel you're building for:
    - Unix-based systems, such as Mac OSX and Linux
    - Non-unix based systems, such as Windows
    - Text can be formatted in terminals, such as bold, background color, and color
        - ANSI escape codes uses escaping to create said styling, but depending on the system, you have to use different codes to escape:
- Text can be formatted in the Linux and macOS terminal via ANSI escape codes, such as:
    - `\033[` prefaces the escape of all text modifications
    - `\033[0m` resets the escaped text modifications (or ends an ongoing one)
    - Check out "ECMA-48 Select Graphic Rendition (SGR)" codes > [Linux console codes](https://man7.org/linux/man-pages/man4/console_codes.4.html)
- []string can be combined into a string with `strings.Join`
- When making an empty slice, `make([]string, length)`, creates a slice with zero values until the "length", which means if you're only appending to a slice and printing out each slice element, the first "length" elements will be empty
- Implementing `.env` files with `joho/godotenv` package, which also was my first time using an external, non-standard library package
    - Within your project RWD, `go get "github.com/joho/godotenv"`
    - `go mod tidy`
    - `go.mod` should contain `require github.com/joho/godotenv v1.x.x`
    - `go.sum` should contain two lines pointing to "github.com/joho/godotenv"
- Generics enable a function to infer parameter types based on the caller of the function
- By using pure functions you can eliminate state variables (checkout input.go), which then enforces the practice of using functions, not state, to retrieve values from different places in your appending 
    - My general assumption is accessing state directly is a risky thing to do becuase of side effects vs getting it with a function that can easily be tested
    - pure functions just make your application far easier and quicker to write tests for, which means better coverage, faster iterations, and more reliable deploys
- SRP and DI are probably the most impactful principles of SOLID, especially if you're using composition over inheritance
- I've heard the critiques, but small functions, specifically where the logic is only 1-10 lines long, is an incredible boon to development... a function can be longer than that, but it shouldn't be the logic (i.e. variables, etc.)
- Map out the layers of an API before decoding the JSON into a defined type

