# Advent of Code 2023

A bunch of code I've create to compete in the 2023 Advent of Code challenge.

My approach for this year:

* Learn how to write in Go
* Use TDD

![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/hellboy1975/aoc2023/.github%2Fworkflows%2Fgo.yml)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/hellboy1975/aoc2023)


## Usage

Clone this repo to your file system directory of choice, and make sure you have Go 1.21.4 installed

Build the project with:

```
go build -o ./bin
```

Run a specific day with a command like:

```
./bin/aoc2023 -day=1 -part=1
```

For limited help:

```
./bin/aoc2023 -h
Usage of ./aoc2023:
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

**Part 1** was relatively straight forward.

**Part 2** was significantly harder.  Initial approach of just a string replacement of the number words failed, as some of the data contained overlapping numbers.
For example: `1twoxyzoneeight`

I came up with a second approach that used a fancy regex using a look ahead:

```
(?=(one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9|0))
```

I'm pretty confident this would have worked, however the various regex libraries available for Go don't support this syntax.  As I'm keen to move on to Day 2,
I have decided refactoring this (again) isn't worth the effort.

### Day 2

Mostly just an exercise in parsing.  In a language I'd be more comfortable with I'd have done this a little more quickly.  Spent more time working on the tests than actual solutions.

With **Part 2** I was lucky that I'd developed a solution that was complimentary to the answer, rather than having to do a reworking.

### Day 3

For Day 2 the **Part 1** solution wasn't especially hard to design, but took me a while writing the appropriate code in Go. 
I spent more time with test writing this time around than in the other days.

**Part 2** was relatively easy to understand, but challenging to write an algorithm for that doesn't just scan and rescan data.  I couldn't think of a clever approach, so that's basically what I did.  I tried to minimise the amount of scanning by rendering the chunks as a 2D array of ints.

### Day 4

**Part 1** is probably the easiest challenge thus far - really just intersect a couple of arrays and multiply stuff.  Relatively easy to test for as well.

I'm used to using a built in array intersect function like PHP has.  Go doesn't have such a thing, but there's a handy package I borrowed here: https://github.com/juliangruber/go-intersect/tree/master

**Part 2** was easy in my head, as I was going to use a form of recursion to solve the issue.  I had problems nailing this, so took a plan B approach which resembled the walkthrough of the example on the AOC2023 website.

## References

Below is a list of sites/resources that have been useful in these challenges:

* https://regex101.com/
* https://gobyexample.com
* https://quii.gitbook.io/learn-go-with-tests/
