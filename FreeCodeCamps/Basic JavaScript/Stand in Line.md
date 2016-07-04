# Stand in Line

In Computer Science a `queue` is an abstract `Data Structure`
where items are kept in order. New items can be added at the back
of the `queue` and old items are taken off from the front of the `queue`.

## Instructions
 - Write a function `nextInLine` which takes an `array`
(`arr`) and a `number` (`item`) as `arguments`.
 - Add the `number` to the `end of the array`.
 - Then remove the `first element of array`. The `nextInLine` function
 should then return the element that was removed.

### Before

```javascript
function nextInLine(arr, item) {
  // Your code here

  return item;  // Change this line
}

// Test Setup
var testArr = [1,2,3,4,5];

// Display Code
console.log("Before: " + JSON.stringify(testArr));
console.log(nextInLine(testArr, 6)); // Modify this line to test
console.log("After: " + JSON.stringify(testArr));
```

### Answers

```javascript
function nextInLine(arr, item) {
  // Your code here

  var queue = arr.push(item);

  var removeItem = arr.shift();

  return removeItem;  // Change this line
}

// Test Setup
var testArr = [1,2,3,4,5];

// Display Code
console.log("Before: " + JSON.stringify(testArr));
console.log(nextInLine(testArr, 6)); // Modify this line to test
console.log("After: " + JSON.stringify(testArr));

```

### Thinking

Here I just follow what they ask
1. Add number at the end of the array. So I use `.push()`
2. Remove the first element. `.shift()`
3. Return the element was removed. `return removeItem` After I setup the var
with the `.shift()` method.