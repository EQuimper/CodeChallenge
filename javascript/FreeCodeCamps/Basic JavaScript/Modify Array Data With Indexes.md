# Modify Array Data With Indexes

Unlike strings, the entries of arrays are
 `mutable` and can be changed freely.

## Example

```javascript
var ourArray = [3,2,1];
ourArray[0] = 1; // equals [1,2,1]
```

## Instructions
 - Modify the data stored at index `0` of `myArray` to a value of `3`.

### Before

```javascript
// Example
var ourArray = [1,2,3];
ourArray[1] = 3; // ourArray now equals [1,3,3].

// Setup
var myArray = [1,2,3];

// Only change code below this line.
```

### Answers

```javascript
// Example
var ourArray = [1,2,3];
ourArray[1] = 3; // ourArray now equals [1,3,3].

// Setup
var myArray = [1,2,3];

// Only change code below this line.

myArray[0] = 3;
```