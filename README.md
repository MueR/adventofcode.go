# 🎄 Advent of Code 🎄

Solutions to [Advent of Code](https://adventofcode.com/) puzzles.

## Setup
Set your AoC session cookie in the `AOC_SESSION_TOKEN` env var.

## Running the solutions
```shell
# Create skeleton for a new day. Defaults to current day/year
make skeleton
# Run tests for the current day
make test
# Run current day
make run
```
All commands support `DAY` and `YEAR` variables to run solutions for a specific day and year. For example:
```shell
DAY=01 YEAR=2024 make run 
```
