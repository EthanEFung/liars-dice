# About This Game
Liars dice is a two or more player game

# Game Rules
Amoung two or more players is evenly distributed a set of dice. Each player has their own hand that is held privately. Typically each hand is 5 dice. A round begins by all players simulaneously rolling their dice, and the first player must make a bid. Each player will take turns choosing the following options:

- Make a bid
  1. up the bid same face
    3 dice of 5 -> 4 dice of 5
  2. up the bid change face
    3 dice of 5 -> 4 dice of 2
  3. same bid higher face
    4 dice of 2 -> 4 dice of 3
- Call the previous bid

Challenging the previous bid prompts the table to reveal their dice and count to see whether the bidder or challenger is correct. The loser must remove one of their dice, and begins the next round.

## Variations
- Ones are considered wild

# Product Requirements
- When a client visits this site, they will either be able to.
  1. Make a new game
  2. Join a game that has not started
  3. Observe an existing game
- In order to make or join a existing game requires a username
- By making a new game, the client becomes the host
- As the host, you will have the ability to start the game
- All other players of a game must wait on the host to begin playing
- All clients that join after a game has started will not be able to play, but should be able to spectate.
- Each player must have an alotted timeframe they are allowed to make their move on.
- If a player does not make a move in the alotted timeframe, they immediately lose.
- Once a player has won the game ends. Users should be able to join the lobby.

## Game Mechanics
Every Game Needs to have the following states:
1. initial: the state in which users are still allowed to join
2. playing: the state in which users are not allowed to join, and players are taking turns
3. ended: the state in which one of the users have won

Every Game needs the following objects:
1. Dice -> Each player that joins should be given 5
2. First Player Reference: initialized on the host
3. Winner Reference: initialized as `nil`
4. Reference to the table (doubly linked list)

Every Round needs the following mechanics:
1. All players dice are rolled
2. first player begins with a bid
3. reference to the current bid
4. reference to the loser

# Technical Requirements
1. each http connection should be upgraded to a websocket connection
2. needing a channel that manages all current websocket connections
3. some sort of dispatcher that sends messages to all current connections
4. redis storage to store session data
