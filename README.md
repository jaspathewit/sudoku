# sudoku 2020
Sudoku solver for the "Sudoku Heaven" problem from the 2020 Belgian Defensie Puzzle.
[https://beldefnews.mil.be/cyber/2020_opgave_NL.pdf](https://beldefnews.mil.be/cyber/2020_opgave_NL.pdf)

Puzzle presented several Sudoku puzzles to be solved. Traditional rules applied for placement of the digits.
However, the "topology" of the puzzles did not follow the traditional 3x3 blocks.
The traditional 9 columns and 9 rows are warped onto each other.
The Golang program solves first three puzzles in a pretty straightforward way.
However, it takes the definition of the topology of the puzzle into account. 

## Build and test
In the unlikely event that this code will interest anyone.

It should be sufficient to clone the repository and go build the program.
To change from one topology (puzzle) to another the main method should be adjusted.

## Contribute
I would be amazed if anyone was interested in contributing... but if anyone were to feel so inclined I'm open to pull requests.

