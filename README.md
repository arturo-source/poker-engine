# Poker Engine

Poker engine is a library written in Golang. It can be used to build:

- A game REST API using the Go HTTP standard library.
- An AI that learns using Reinforcement Learning.
- A poker-odds application ([CLI example](https://github.com/arturo-source/poker-odds)).

Ensure the library is installed:

```bash
go get -u github.com/arturo-source/poker-engine/
```

## Example

NewCard reads a string and returns the card, only if string has valid number and suit.

- Valid numbers [**A K Q J T 9 8 7 6 5 4 3 2**].
- Valid suits [**s c h d**].

```go
package main

import (
    "fmt"

    "github.com/arturo-source/poker-engine"
)

func main(t *testing.T) {
 c := poker.NewCard
 cards := poker.JoinCards(c("7s"), c("Kc"), c("7c"), c("As"), c("Ad"))

 winningCards, found := poker.TwoPair(cards)
 if found {
    fmt.Println("Winning cards:", winningCards)
 } else {
    fmt.Println("Two pair was not found")
 }
}
```

I will add an example for a full game, when the TODO is finished.

## TODO

I have to add more funcs to `Game`, to be able to control the game with just a game:

- Dealing cards should shuffle the deck first.
- Add players.
- Add default entry coins per player.
- A func to proceed to the next BoardState.
- Subtract coins from the players when they bet.
- Add coins when a player wins.
- Calculate how much money if to each player if multiple of them win.
