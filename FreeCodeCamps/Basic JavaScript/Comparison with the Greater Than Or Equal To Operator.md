# Comparison with the Greater Than Or Equal To Operator

The greater than or equal to operator (`>=`) compares the values of
two numbers. If the number to the left is greater than or equal to
the number to the right, it returns true. Otherwise, it returns false.

Like the equality operator, greater than or equal to operator
will **convert data types while comparing.**

## Examples

```javascript
 6  >=  6  // true
 7  >= '3' // true
 2  >=  3  // false
'7' >=  9  // false
```

## Instructions
 - Add the greater than or equal to operator to the indicated lines
 so that the return statements make sense.

### Before

```javascript
function testGreaterOrEqual(val) {
  if (val) {  // Change this line
    return "20 or Over";
  }

  if (val) {  // Change this line
    return "10 or Over";
  }

  return "9 or Under";
}

// Change this value to test
testGreaterOrEqual(10);
```

### Answers

```javascript
function testGreaterOrEqual(val) {
  if (val >= 20) {  // Change this line
    return "20 or Over";
  }

  if (val >= 10) {  // Change this line
    return "10 or Over";
  }

  return "9 or Under";
}

// Change this value to test
testGreaterOrEqual(10);
```