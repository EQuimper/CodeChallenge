# Iterate with JavaScript While Loops

You can run the same code multiple times by using a loop.

Another type of JavaScript loop is called a `while loop`, because
it runs `while` a specified condition is true and **stops once
that condition is no longer true**.

```javascript
var ourArray = [];
var i = 0;
while(i < 5) {
  ourArray.push(i);
  i++;
}
```
In the code example above, the `while` loop will execute 5 times and append the numbers 0 through 4 to `ourArray`.


Let's try getting a while loop to work by pushing values to an array.

## Instructions
 - Add the numbers 5 through 0 (inclusive) in descending order to `myArray` using a `while` loop.

### Before

```javascript
// Setup
var myArray = [];

// Only change code below this line.
```

### Answers

```javascript
// Setup
const myArray = [];

// Only change code below this line
let i = 0;

while (i < 6) {
  myArray.push(i);
  i++;
}
myArray.reverse(myArray);
console.log(myArray);
```

