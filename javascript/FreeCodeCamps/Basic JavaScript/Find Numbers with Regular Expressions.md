# Find Numbers with Regular Expressions

We can use special selectors in `Regular Expressions` to select a particular
type of value.

One such selector is the digit selector `\d` which is used to
retrieve one digit (e.g. numbers 0 to 9) in a string.

In JavaScript, it is used like this: `/\d/g`.

Appending a plus sign (`+`) after the selector, e.g. `/\d+/g`, allows
this regular expression to match one or more digits.

The trailing `g` is short for '`global`', which allows this regular
expression to find all matches rather than stop at the first match.

## Instructions
 - Use the `\d` selector to select the number of numbers
 in the string, allowing for the possibility of one or more digit.

### Before

```javascript
// Setup
var testString = "There are 3 cats but 4 dogs.";

// Only change code below this line.

var expression = /.+/g;  // Change this line

// Only change code above this line

// This code counts the matches of expression in testString
var digitCount = testString.match(expression).length;
```

### Answers

```javascript
// Setup
var testString = "There are 3 cats but 4 dogs.";

// Only change code below this line.

var expression = /\d+/g;  // Change this line

// Only change code above this line

// This code counts the matches of expression in testString
var digitCount = testString.match(expression).length;
```