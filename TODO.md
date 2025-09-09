# Card Game Engine: War

## Cards -> class
- number value
- suit of cards (spades, Clubs, Hearts, Diamonds)

## Deck - class (holds the cards)

## Two players
- holds cards -> array/list
- take turn putting top card in the middle to compare -> if statement
  - if player 1 card is higher, player 1 gets card
  - else if player 2 card is higher player 2 gets card
  - else in the event of a tie, both players put down 4 cards on top of the already existing card, and the last card of the 4 that is higher wins
  - Cards that are won are bottom of the array (cards in hand array)

## Shuffle cards -> randomly sort the array
