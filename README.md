# go_2048

## description

it's a cli 2048 game in golang. i don't know what else to say. it doesn't work
right now

## long term plan

this is a stepping stone to build an algorithm to solve 2048 games for me. i
plan to write it in go because the algorithm would be faster than using
something like python. i want to be able to read a 2048 board from the screen
with python, and pump that into my algorithm in go. then i feed the output from
go into the python script and send an input to the game, thus creating a genius
at 2048 (hopefully).

## short term plan

 - make the tiles actually move 😭
 - make sure the tiles actually merge with each other
 - add comments to everything
 - split things up into smaller functions, specifically in a way that will help
   me build the algorithm later
 - put things into multiple files so i don't have one large main.go file

## build instructions

make sure you have go installed, clone the repo and run `go build` inside the
directory. after that you can put the binary or executable wherever you want
