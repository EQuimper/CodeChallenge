# Nesting For Loops

If you have a multi-dimensional array, you can use the same
logic as the prior waypoint to loop through both the array and
any sub-arrays. Here is an example:

```javascript
var arr = [
  [1,2], [3,4], [5,6]
];
for (var i=0; i < arr.length; i++) {
  for (var j=0; j < arr[i].length; j++) {
    console.log(arr[i][j]);
  }
}
```

This outputs each sub-element in `arr` one at a time. Note that for
the inner loop, we are checking the `.length` of `arr[i]`, since `arr[i]`
is itself an `array`.

## Instructions
 - Modify function `multiplyAll` so that it multiplies the product
 variable by each number in the sub-arrays of `arr`

### Before

```javascript
function multiplyAll(arr) {
  var product = 1;
  // Only change code below this line

  // Only change code above this line
  return product;
}

// Modify values below to test your code
multiplyAll([[1,2],[3,4],[5,6,7]]);
```

### Answers

```javascript
function multiplyAll(arr) {
  var product = 1;
  // Only change code below this line

  for (var i = 0; i < arr.length; i++) {
    for (var x = 0; x < arr[i].length; x++) {
      product *= arr[i][x];
    }
  }


  // Only change code above this line
  return product;
}

// Modify values below to test your code
multiplyAll([[1,2],[3,4],[5,6,7]]);
```

### Thinking

1. Need to loop the first array for find all the sub one
2. Loop inside the sub one for get all the values
3. Use *= for multiply product to the values of the sub arr