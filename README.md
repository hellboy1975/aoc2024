# Advent of Code 2024

A bunch of code I've create to compete in the 2024 Advent of Code challenge.

My approach for this year:

* Continue to improve my Golang craft
* Use TDD - possibly less of a focus on this - we'll see!

![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/hellboy1975/aoc2024/.github%2Fworkflows%2Fgo.yml)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/hellboy1975/aoc2024)


## Usage

Clone this repo to your file system directory of choice.  Best to have a version of Go that is at least 1.22

To add data, create a folder called `data` and then the respective day files are `day_<day>_<part>.txt`.  ie. `day_1_2.txt` for day 1 part 2.

Build the project with:

```
go build -o ./bin
```

Run a specific day with a command like:

```
./bin/aoc2024 -day=1 -part=1
```

For limited help:

```
./bin/aoc2024 -h
Usage of ./aoc2024:
  -day int
        which day to run (default 1)
  -part int
        which part to run (default 1)
```

Run all of the tests with:

```
go test ./...
```

## Day notes

### Day 1

**Part 1**

**Part 2**

## References

Below is a list of sites/resources that have been useful in these challenges:

* https://regex101.com/
* https://gobyexample.com
* https://quii.gitbook.io/learn-go-with-tests/
