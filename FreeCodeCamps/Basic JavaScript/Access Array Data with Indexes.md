# Access Array Data with Indexes

We can access the data inside arrays using `indexes`.

Array indexes are written in the same bracket notation that strings use,
except that instead of specifying a character, they are specifying an
entry in the array. Like strings, arrays use `zero-based` indexing, so
the first element in an array is element 0.

## Example

```javascript
var array = [1,2,3];
array[0]; // equals 1
var data = array[1];  // equals 2
```

## Instructions
 - Create a variable called `myData` and set it to equal the first
 value of `myArray` using `bracket notation`.

### Before

```javascript
// Example
var ourArray = [1,2,3];
var ourData = ourArray[0]; // equals 1

// Setup
var myArray = [1,2,3];

// Only change code below this line.
```

### Answers

```javascript
// Example
var ourArray = [1,2,3];
var ourData = ourArray[0]; // equals 1

// Setup
var myArray = [1,2,3];

// Only change code below this line.
var myData = myArray[0];
```