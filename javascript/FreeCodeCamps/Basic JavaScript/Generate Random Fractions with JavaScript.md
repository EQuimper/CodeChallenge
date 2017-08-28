# Generate Random Fractions with JavaScript

Random numbers are useful for creating random behavior.

JavaScript has a `Math.random()` function that
generates a random decimal number between 0 (inclusive) and
not quite up to 1 (exclusive). Thus `Math.random()` can return a 0 but
never quite return a 1

## Note
 *Like Storing Values with the Equal Operator, all function calls will be
 resolved before the return executes,
 so we can simply return the value of the Math.random() function.*

## Instructions
 - Change `randomFraction` to return a random number instead of returning 0.

### Before

```javascript
function randomFraction() {

  // Only change code below this line.

  return 0;

  // Only change code above this line.
}
```

### Answers

```javascript
function randomFraction() {

  // Only change code below this line.

  return Math.random();

  // Only change code above this line.
}
```