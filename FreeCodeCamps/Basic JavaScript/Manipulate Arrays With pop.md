# Manipulate Arrays With pop

Another way to change the data in an array is with the `.pop()` function.

.`pop()` is used to "pop" a value off of the end of an array. We can
store this "popped off" value by assigning it to a variable.

Any type of entry can be "popped" off of an
`array - numbers, strings`, even `nested arrays`.

## Example

```javascript
var oneDown = [1, 4, 6].pop();
```

the variable oneDown now holds the value 6 and the array becomes `[1, 4]`.

## Instructions
 - Use the `.pop()` function to remove the last item from `myArray`,
 assigning the "popped off" value to `removedFromMyArray`.

### Before

```javascript
// Example
var ourArray = [1,2,3];
var removedFromOurArray = ourArray.pop();
// removedFromOurArray now equals 3, and ourArray now equals [1,2]

// Setup
var myArray = [["John", 23], ["cat", 2]];

// Only change code below this line.
var removedFromMyArray;
```

### Answers

```javascript
// Example
var ourArray = [1,2,3];
var removedFromOurArray = ourArray.pop();
// removedFromOurArray now equals 3, and ourArray now equals [1,2]

// Setup
var myArray = [["John", 23], ["cat", 2]];

// Only change code below this line.
var removedFromMyArray = myArray.pop();
```