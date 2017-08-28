# Counting Cards

In the casino game Blackjack, a player can gain an advantage
over the house by keeping track of the relative number of high
and low cards remaining in the deck. This is called Card Counting.

Having more high cards remaining in the deck favors
the player. Each card is assigned a value according to the table
below. When the count is `positive`, the player should bet `high`.
When the count is zero or `negative`, the player should bet `low`.

Count Change | Cards
--- | ---
+1 | 2, 3, 4, 5, 6
0 |	7, 8, 9
-1 | 10, 'J', 'Q', 'K', 'A'

You will write a card counting function. It will receive a
`card parameter` and `increment` or `decrement` the `global count variable`
according to the card's value (see table). The function
will then return a string with the current count and the string `"Bet"` if
the count is `positive`, or `"Hold"` `if` the count is `zero` or `negative`.
The current count and the player's decision (`"Bet" or "Hold"`) should be
separated by a single space.

## Example
```javascript
"-3 Hold"
"5 Bet"
```

## Hint
 - Do NOT reset count to 0 when value is 7, 8, or 9.

### Before

```javascript
var count = 0;

function cc(card) {
  // Only change code below this line


  return "Change Me";
  // Only change code above this line
}

// Add/remove calls to test your function.
// Note: Only the last will display
cc(2); cc(3); cc(7); cc('K'); cc('A');
```

### Answers

```javascript
var count = 0;

function cc(card) {
  // Only change code below this line

  var msg = '';

  switch (card) {
    case 2:
    case 3:
    case 4:
    case 5:
    case 6:
      count++;
      break;
    case 7:
    case 8:
    case 9:
      count = count;
      break;
    case 10:
    case "J":
    case "Q":
    case "K":
    case "A":
      count--;
      break;
  }

  if (count === 5) msg = "5 Bet";
  else if (count === 0) msg = "0 Hold";
  else if (count === -5) msg = "-5 Hold";
  else if (count === -1) msg = "-1 Hold";
  else if (count === 1) msg = "1 Bet";


  return msg;
  // Only change code above this line
}

// Add/remove calls to test your function.
// Note: Only the last will display
cc(2); cc(3); cc(7); cc('K'); cc('A');
```

### Thinking

1. Here we need to setup a global var call `count`.
2. We use the switch statement to setup the count with the card who is
the args
3. I take the if/else if statement for practice the return msg