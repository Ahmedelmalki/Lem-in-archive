# lem-in

## Description
`lem-in` is a program that simulates ants moving through a colony of rooms and tunnels. The objective is to calculate and display the quickest way to get a specified number of ants from the starting room (`##start`) to the ending room (`##end`) with as few moves as possible. 

## How it works
- The colony consists of rooms connected by tunnels.
- All ants start in the `##start` room, and the goal is to move them to the `##end` room.
- The program must find and display the fastest way to achieve this.
- The shortest path is not always the most optimal due to the potential for traffic jams.

## Input
- The program reads from a file passed as an argument, which contains:
  - The number of ants.
  - Descriptions of rooms and their coordinates.
  - Tunnels connecting the rooms.
- Example of file format:
```
$ go run . test0.txt
3
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5

L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0
$
```


## Error Handling
- The program must handle a variety of invalid inputs, such as:
- No start or end room.
- Invalid number of ants.
- Invalid room or tunnel descriptions.
- When an error is encountered, an appropriate message such as `ERROR: invalid data format` should be returned.

## Features
- Multiple paths can be taken by the ants at the same time.
- A tunnel can only be used by one ant per move.
- Each room can only hold one ant at a time, except for the `##start` and `##end` rooms, which can hold multiple ants.

## Usage
To run the program, use the following command:
```bash
$ go run . file.txt
