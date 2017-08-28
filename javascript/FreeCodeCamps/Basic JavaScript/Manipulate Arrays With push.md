# Manipulate Arrays With push

An easy way to append data to the end of an array is via
the `push()` function.

`.push()` takes one or more `parameters` and
"pushes" them onto the end of the array.

## Example

```javascript
var arr = [1,2,3];
arr.push(4);
// arr is now [1,2,3,4]
```

## Instructions
 - Push `["dog", 3]` onto the end of the `myArray` variable.

### Before

```javascript
// Example
var ourArray = ["Stimpson", "J", "cat"];
ourArray.push(["happy", "joy"]);
// ourArray now equals ["Stimpson", "J", "cat", ["happy", "joy"]]

// Setup
var myArray = [["John", 23], ["cat", 2]];

// Only change code below this line.
```

### Answers

```javascript
// Example
var ourArray = ["Stimpson", "J", "cat"];
ourArray.push(["happy", "joy"]);
// ourArray now equals ["Stimpson", "J", "cat", ["happy", "joy"]]

// Setup
var myArray = [["John", 23], ["cat", 2]];

// Only change code below this line.
myArray.push(['dog', 3]);
```