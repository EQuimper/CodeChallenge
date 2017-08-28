# Sift through Text with Regular Expressions

Regular expressions are used to find certain
words or patterns inside of strings.

For example, if we wanted to find the word the in the string
`The dog chased the cat`, we could use the following regular
expression: `/the/gi`

Let's break this down a bit:

`/` is the start of the regular expression.

the is the pattern we want to match.

`/` is the end of the regular expression.

`g` means **global**, which causes the pattern to return
all matches in the string, not just the first one.

`i` means that we want to ignore the case (`uppercase or lowercase`) when
searching for the pattern.

## Instructions
 - Select all the occurrences of the word and in `testString`.

You can do this by replacing the `.` part of the regular
expression with the word `and`.

### Before

```javascript
// Setup
var testString = "Ada Lovelace and Charles Babbage designed the first computer and the software that would have run on it.";

// Example
var expressionToGetSoftware = /software/gi;
var softwareCount = testString.match(expressionToGetSoftware).length;


// Only change code below this line.

var expression = /./gi;  // Change this Line

// Only change code above this line

// This code counts the matches of expression in testString
var andCount = testString.match(expression).length;
```

### Answers

```javascript
// Setup
var testString = "Ada Lovelace and Charles Babbage designed the first computer and the software that would have run on it.";

// Example
var expressionToGetSoftware = /software/gi;
var softwareCount = testString.match(expressionToGetSoftware).length;


// Only change code below this line.

var expression = /and/gi;  // Change this Line

// Only change code above this line

// This code counts the matches of expression in testString
var andCount = testString.match(expression).length;
```