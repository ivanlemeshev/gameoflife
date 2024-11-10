# Game of Life

![Build](https://github.com/ivanlemeshev/gameoflife/actions/workflows/build.yml/badge.svg)
![Tests](https://github.com/ivanlemeshev/gameoflife/actions/workflows/test.yml/badge.svg)
![Linter](https://github.com/ivanlemeshev/gameoflife/actions/workflows/lint.yml/badge.svg)
![Gosec](https://github.com/ivanlemeshev/gameoflife/actions/workflows/sec.yml/badge.svg)

<img src="./gameoflife.gif" width="600" alt="Game of Life"/>

I recently visited a workshop, [9 Steps for Better Object Oriented Design](https://www.meetup.com/tech-excellence-finland/events/304005147/),
in Helsinki, which gave me an idea for implementing this game. I did not follow
strictly the principles presented there. I just wanted to make a game that
works in the terminal.

It is a terminal-based implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)
with finite grid. (The cells outside the grid are considered to be dead.)

The Game of Life, also known simply as Life, is a cellular automaton devised by
the British mathematician John Horton Conway in 1970.

## Rules

1. Any live cell with fewer than two live neighbors dies, as if by underpopulation.
2. Any live cell with two or three live neighbors lives on to the next generation.
3. Any live cell with more than three live neighbors dies, as if by overpopulation.
4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

## Build and run

```bash
make build
./bin/gameoflife
```

or

```bash
make run
```
