# emil

Writing some go code to learn chess programming.

## First goal:

The computer should win a game with a king and a rock against a king (KRK).
So we need:
- a chess board
- king and rock pieces
- movements for king and rock
- delete illegal moves
- create an endgame database for KRK
	- Step 1: Generating all possible positions
	- Step 2: Evaluating positions using retrograde analysis
	- Step 3: Verification (todo)
- play for checkmate


### retrograde analysis:
- generating all 249.984 (64 x 63 x 62) positions for KRK take 10s
- filter 26.040 illegal positions, where the kings are to close
- remaining 223.944 positions for analysis 
- found 216 checkmates and 68 patt positions in < 1s, analysing only 13.144 positions where
	- the black king is on a border square and
	- the distance between the kings is 2


### numbers:

	db.FindMatesIn 0:     216 boards
	db.FindMatesIn 1:   1.512 boards
	db.FindMatesIn 2:     624 boards
	db.FindMatesIn 3:   4.676 boards
	db.FindMatesIn 4:   1.948 boards
	db.FindMatesIn 5:   3.852 boards
	db.FindMatesIn 6:     648 boards
	db.FindMatesIn 7:   1.900 boards
	db.FindMatesIn 8:   1.584 boards
	db.FindMatesIn 9:   4.848 boards
	db.FindMatesIn 10:  3.768 boards
	db.FindMatesIn 11:  8.708 boards
	db.FindMatesIn 12:  4.728 boards
	db.FindMatesIn 13: 11.320 boards
	db.FindMatesIn 14:  5.444 boards
	db.FindMatesIn 15: 17.172 boards
	db.FindMatesIn 16: 11.448 boards
	db.FindMatesIn 17: 20.088 boards
	db.FindMatesIn 18: 13.672 boards
	db.FindMatesIn 19: 19.016 boards
	db.FindMatesIn 20: 15.872 boards
	db.FindMatesIn 21: 20.476 boards
	db.FindMatesIn 22: 22.788 boards
	db.FindMatesIn 23: 21.480 boards
	db.FindMatesIn 24: 28.732 boards
	db.FindMatesIn 25: 17.824 boards
	db.FindMatesIn 26: 33.516 boards
	db.FindMatesIn 27: 16.136 boards
	db.FindMatesIn 28: 36.372 boards
	db.FindMatesIn 29:  5.244 boards
	db.FindMatesIn 30: 17.284 boards
	db.FindMatesIn 31:    916 boards
	db.FindMatesIn 32:  3.056 boards
	db.FindMatesIn 33:      0 boards

	duration 16m0.7s
