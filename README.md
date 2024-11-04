# Description

Dicthesaurus is a dictionary and thesaurus CLI program built only with Go standard packages. 

Currently, it makes use of the [Free Dictionary API](https://dictionaryapi.dev/).

## Motivation

It is my intent to replace my love of OSX's built-in dictionary lookup through the spotlight tool now that I'm solely using Linux, and I've found there's never a day or time where I don't have an active terminal ready for a quick search!

# Instructions

## Installation

TBD

## Usage

Via your terminal:

`$ dicthesaurus <word> [FLAGS]`

Flags can be used in either format: `-d, --d`

## Command Flags

| Flag | Description |
|--|--|
| `h, help` | help |

## Subcommand Flags

| Flag | Description |
|--|--|
| `e` | Display the word in a sentence |
| `t` | Include thesaurus results |
| `ud` | Also query Urban Dictionary for results |

# Use Cases

- [x] Defined command flags
- [x] Undefined command flags
- [x] Missing subcommand
- [ ] Defined subcommand flags
- [x] Undefined subcommand flags

# What I learned
