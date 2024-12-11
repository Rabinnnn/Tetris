# TETRIS-OPTIMIZER

## Description
This is a program that receives only one argument, a path to a text file which contains a list of tetrominoes and assemble them in order to create the smallest square possible.

The program must:
* Compile successfully
* Assemble all of the tetrominoes in order to create the smallest square possible.
* Identify each tetromino in the solution by printing them with uppercase latin letters (A for the first one, B for the second, etc)
* Expect at least one tetromino in the text file
* In case of bad format on the tetrominoes or bad file format it should print ERROR

## Example of a text file

```bash
$ cat -e data.txt
...#$
...#$
...#$
...#$
$
....$
....$
....$
####$
$
.###$
...#$
....$
....$
$
....$
..##$
.##.$
....$
$
....$
.##.$
.##.$
....$
$
....$
....$
##..$
.##.$
$
##..$
.#..$
.#..$
....$
$
....$
###.$
.#..$
....$
```
```bash
go run . data.txt | cat -e
ABBBB.$
ACCCEE$
AFFCEE$
A.FFGG$
HHHDDG$
.HDD.G$
$

```

## Usage
1. Clone the repo
```bash
$ git clone https://learn.zone01kisumu.ke/git/rotieno/tetris-optimizer.git
```
2. Navigate to the project directory
```bash
$ cd tetris-optimizer
```
3. Execute
```bash
$ go run main.go data.txt
```

## Tests
* To run the tests:
```bash
$ cd tests
$ go test
```

## Author
* [Rabin Otieno](https://learn.zone01kisumu.ke/git/rotieno/tetris-optimizer.git)