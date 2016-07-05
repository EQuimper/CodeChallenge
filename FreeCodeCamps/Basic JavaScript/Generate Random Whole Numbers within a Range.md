# Generate Random Whole Numbers within a Range

Instead of generating a random number between zero and a
given number like we did before, we can generate a random number
that falls within a range of two specific numbers.

To do this, we'll define a minimum number min and a maximum number max.

Here's the formula we'll use. Take a moment to read it
and try to understand what this code is doing:

`Math.floor(Math.random() * (max - min + 1)) + min`

## Instructions
 - Create a function called `randomRange` that takes a range `myMin` and `myMax`
 and returns a random number that's greater than or equal to `myMin`, and is
 less than or equal to `myMax`, inclusive.

### Before

```javascript
// Example
function ourFunction(ourMin, ourMax) {

  return Math.floor(Math.random() * (ourMax - ourMin + 1)) + ourMin;
}

ourFunction(1, 9);

// Only change code below this line.

function randomRange(myMin, myMax) {

  return 0; // Change this line

}

// Change these values to test your function
var myRandom = randomRange(5, 15);
```

### Answers

```javascript
// Example
function ourFunction(ourMin, ourMax) {

  return Math.floor(Math.random() * (ourMax - ourMin + 1)) + ourMin;
}

ourFunction(1, 9);

// Only change code below this line.

function randomRange(myMin, myMax) {

  return Math.floor(Math.random() * (myMax - myMin + 1) + myMin);

}

// Change these values to test your function
var myRandom = randomRange(5, 15);
```