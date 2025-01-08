# Pichess
###### You like Lichess? You will love Pichess ❤️ 

![](https://github.com/ahmed-debbech/go_chess_puzzle/actions/workflows/deploy.yml/badge.svg)

A Chess puzzle generator from Lichess.org match databases using Stockfish to calculate best moves.

### Demo


### How does it work?

##### Generation phase:
I built three simple programs:
**1 (Cutter)** - Parses the downloaded file from [Lichess database](https://database.lichess.org/) (a huge text file of several GBs). 
The file consists of thousands of PGN games and Cutter is responsible for splitting every game's PGN to a seperate file inside a specified output directory and naming them with unique IDS. 
**2 (Generator)** -  The main core that produces puzzle candidates.
- Generator selects and opens a PGN file with a given random id (say 15260) from the output directory. 
- Generator then reads the PGN and jumps to also a random position played already in the game (say 15).
- At this point Generator Calls stockfish chess engine and tells it "Hey, here is a chess game at the position (say 15) and white plays now, can you finish the game for me?".
- Note that Generator asks stockfish to play at depth 24 and depth 4 so that we can reach to a checkmate. 
- Eventually Generator points to the last 2 to 5 moves in the game that is finished by stockfish. 
- Now we have a new puzzle candidate

**3 (Storer)** - Stores puzzle candidates coming from generator to MongoDB
Storer accepts any puzzle candidate coming from Generator to save it directly into `puzzles` collection in MongoDB


##### Serving phase:
**The backend** 
The actual program that you need to deploy is `backend` that will serve all web pages content and do everything related to checking if client solved or seen a puzzle.
The rest three programs we talked about above are like a toolbox for generating/storing puzzles forbackend to be able to serve them.

### We depend on these amazing libraries...
* [notnil/chess](https://github.com/notnil/chess) library for easier manipulation of chess games in Golang.
* [chessboardjs](https://chessboardjs.com/) the library that draws chess the board you see in your browser.

### Want to contribute?
Please feel free to fork this repo and to open a new pull request with your own modifications.
Contact me at debbech.ahmed@gmail.com