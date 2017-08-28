# Generate Random Whole Numbers with JavaScript

It's great that we can generate random decimal numbers,
but it's even more useful if we use it to generate random whole numbers.

1. Use `Math.random()` to generate a random decimal.
2. Multiply that random decimal by 20.
3. Use another function, `Math.floor()` to round the number
down to its nearest whole number.

Remember that `Math.random()` can never quite return a 1 and,
because we're rounding down, it's impossible to actually get 20.
This technique will give us a whole number between 0 and 19.

Putting everything together, this is what our code looks like:

`Math.floor(Math.random() * 20);`

We are calling `Math.random()`, multiplying the result by 20, then
passing the value to `Math.floor()` function to round the value down to
the nearest whole number.

## Instructions
 - Use this technique to generate and return a random whole number between
 0 and 9.

### Before

```javascript
var randomNumberBetween0and19 = Math.floor(Math.random() * 20);

function randomWholeNum() {

  // Only change code below this line.

  return Math.random();
}
```

### Answers

```javascript
var randomNumberBetween0and19 = Math.floor(Math.random() * 20);

function randomWholeNum() {

  // Only change code below this line.

  return Math.floor(Math.random() * 10);
}
```