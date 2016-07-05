# Iterate Through an Array with a For Loop

A common task in JavaScript is to iterate through the contents
of an `array`. One way to do that is with a `for loop`. This code will
output each element of the `array` arr to the console:

```javascript
var arr = [10,9,8,7,6];
for (var i=0; i < arr.length; i++) {
   console.log(arr[i]);
}
```

Remember that `Arrays` have `zero-based` numbering, which means
the **last index of the array is length - 1**. Our condition for
this loop is `i < arr.length`, which stops when `i is at length - 1`.

## Instructions
 - Declare and initialize a variable total to 0. Use a `for loop` to
 add the value of each element of the `myArr` array to total.

### Before

```javascript
// Example
var ourArr = [ 9, 10, 11, 12];
var ourTotal = 0;

for (var i = 0; i < ourArr.length; i++) {
  ourTotal += ourArr[i];
}

// Setup
var myArr = [ 2, 3, 4, 5, 6];

// Only change code below this line
```

### Answers

```javascript
// Example
var ourArr = [ 9, 10, 11, 12];
var ourTotal = 0;

for (var i = 0; i < ourArr.length; i++) {
  ourTotal += ourArr[i];
}

// Setup
var myArr = [ 2, 3, 4, 5, 6];
var total = 0;

for (var i = 0; i < myArr.length; i++) {
  total += myArr[i];
}
// Only change code below this line
```