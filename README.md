# emil

Writing some go code to learn chess programming.

## First goal:

The computer should win a game with a king and a rock against a king (KRK).
So we need:
- a chess board (done)
- king and rock pieces (done)
- movements for king and rock (done)
- delete illegal moves (done)
- create an endgame database for KRK
	- Step 1: Generating all possible positions (done)
	- Step 2: Evaluating positions using retrograde analysis
	- Step 3: Verification
- play for checkmate


### retrograde analysis:
- generating all 249.984 (64 x 63 x 62) positions for KRK take 10s
- filter 26.040 illegal positions, where the kings are to close
- remaining 223.944 positions for analysis 
- found 216 checkmates and 68 patt positions in < 1s, analysing only 13.144 positions where
	- the black king is on a border square and
	- the distance between the kings is 2


### profiling:

	$ go test
	PASS
	ok  	github.com/ujanssen/emil	171.641s
	ok  	github.com/ujanssen/emil	168.514s
