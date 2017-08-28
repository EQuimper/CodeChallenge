# Find Whitespace with Regular Expressions

We can also use `regular expression` selectors like `\s` to find
whitespace in a string.

The whitespace characters are `" "` (space), `\r` (the carriage return),
`\n` (newline), `\t` (tab), and `\f` (the form feed).

The whitespace regular expression looks like this:

`/\s+/g`

## Instructions
 - Use `\s` to select all the whitespace characters in the sentence string.

### Before

```javascript
// Setup
var testString = "How many spaces are there in this sentence?";

// Only change code below this line.

var expression = /.+/g;  // Change this line

// Only change code above this line

// This code counts the matches of expression in testString
var spaceCount = testString.match(expression).length;
```

### Answers

```javascript
// Setup
var testString = "How many spaces are there in this sentence?";

// Only change code below this line.

var expression = /\s+/g;  // Change this line

// Only change code above this line

// This code counts the matches of expression in testString
var spaceCount = testString.match(expression).length;
```