# Access MultiDimensional Arrays With Indexes

One way to think of a `multi-dimensional` array, is as an array of
arrays. When you use brackets to access your array, the first set of
bracket refers to the entries in the outer-most (the first level) array,
and each additional pair of brackets refers to the next level of entries
inside.

## Example

```javascript
var arr = [
    [1,2,3],
    [4,5,6],
    [7,8,9],
    [[10,11,12], 13, 14]
];
arr[3]; // equals [[10,11,12], 13, 14]
arr[3][0]; // equals [10,11,12]
arr[3][0][1]; // equals 11
```

## Instructions
 - Using bracket notation select an element from `myArray` such that
 `myData` is equal to 8.

### Before

```javascript
// Setup
var myArray = [[1,2,3], [4,5,6], [7,8,9], [[10,11,12], 13, 14]];

// Only change code below this line.
var myData = myArray[0][0];
```

### Answers

```javascript
// Setup
var myArray = [[1,2,3], [4,5,6], [7,8,9], [[10,11,12], 13, 14]];

// Only change code below this line.
var myData = myArray[2][1];
```

### Thinking

Pretty straight-forward we just need to think about the first bracket
notation is the outer array and the second one is inside. Don't forget we
start counting at 0.