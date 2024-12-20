# Advent of Code 2024

A bunch of code I've create to compete in the [2024 Advent of Code](https://adventofcode.com/2024/) challenge.

My approach for this year:

* Continue to improve my Golang craft
* Use TDD - possibly less of a focus on this - we'll see!

![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/hellboy1975/aoc2024/.github%2Fworkflows%2Fgo.yml)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/hellboy1975/aoc2024)


## Usage

Clone this repo to your file system directory of choice.  Best to have a version of Go that is at least 1.22

To add data, create a folder called `data` and then the respective day files are `day_<day>_<part>.txt`.  ie. `day_1_2.txt` for day 1 part 2.  Test data can be run by passing the flag --test and having a `day_<day>_<part>_test.txt` file in the `data` directory

Build the project with:

```
go build -o ./bin/
```

Run a specific day with a command like:

```
./bin/aoc2024 -day=1 -part=1
```

You

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

* Day 3 Part 2 - solutions looks good (to me) but doesn't resolve in a correct answer.  Can't work out why not :(

## References

Below is a list of sites/resources that have been useful in these challenges:

* https://regex101.com/
* https://gobyexample.com
* https://quii.gitbook.io/learn-go-with-tests/
* https://pterm.sh/

## Other Stuff

For debugging in VSCode you can use a launch.json config such as:

```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "args": ["--day", "2", "--part", "1", "--test"]
        }
    ]
}
```
