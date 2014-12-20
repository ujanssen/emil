#A chess board

The board has 64 squares arranged in an eight-by-eight grid.
The rows called ranks and denoted with numbers 1 to 8.
The columns called files and denoted with letters a to h:

<pre>
a8 b8 c8 d8 e8 f8 g8 h8 
a7 b7 c7 d7 e7 f7 g7 h7 
a6 b6 c6 d6 e6 f6 g6 h6 
a5 b5 c5 d5 e5 f5 g5 h5 
a4 b4 c4 d4 e4 f4 g4 h4 
a3 b3 c3 d3 e3 f3 g3 h3 
a2 b2 c2 d2 e2 f2 g2 h2 
a1 b1 c1 d1 e1 f1 g1 h1 
</pre>

The board can be represented as an array of size 64.
The square a1 has index 0, b1 has index 1, c1 has index 2, ..., a2 has index 8, ..., and h8 has index 63:

<pre>
   a  b  c  d  e  f  g  h  
8 56 57 58 59 60 61 62 63 8
7 48 49 50 51 52 53 54 55 7
6 40 41 42 43 44 45 46 47 6
5 32 33 34 35 36 37 38 39 5
4 24 25 26 27 28 29 30 31 4
3 16 17 18 19 20 21 22 23 3
2  8  9 10 11 12 13 14 15 2
1  0  1  2  3  4  5  6  7 1
   a  b  c  d  e  f  g  h  
</pre>

