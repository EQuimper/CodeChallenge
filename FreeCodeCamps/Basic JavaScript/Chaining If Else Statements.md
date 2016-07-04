# Chaining If Else Statements

`if/else` statements can be chained together for complex logic.
Here is `pseudocode` of multiple chained `if / else if` statements:

```javascript
if (condition1) {
  statement1
} else if (condition2) {
  statement2
} else if (condition3) {
  statement3
. . .
} else {
  statementN
}
```

## Instructions
 - Write chained `if/else if` statements to fulfill the following conditions:

```javascript
num < 5 - return "Tiny"
num < 10 - return "Small"
num < 15 - return "Medium"
num < 20 - return "Large"
num >= 20 - return "Huge"
```

### Before

```javascript
function testSize(num) {
  // Only change code below this line


  return "Change Me";
  // Only change code above this line
}

// Change this value to test
testSize(7);
```

### Answers

```javascript
function testSize(num) {
  // Only change code below this line

  if (num < 5) return "Tiny";

  else if (num < 10) return "Small";

  else if (num < 15) return "Medium";

  else if (num < 20) return "Large";

  else return "Huge";

  // Only change code above this line
}

// Change this value to test
testSize(7);
```