# Use Bracket Notation to Find the Last Character in a String

In order to get the last letter of a string, you can subtract one
from the string's length.

For example, if `var firstName = "Charles",` you can get the value of the
last letter of the string by using `firstName[firstName.length - 1]`.

## Instructions
 - Use `bracket notation` to find the last character in the `lastName` variable.

## Hint
 - Try looking at the `lastLetterOfFirstName` variable declaration if you get
 stuck.

### Before

```javascript
// Example
var firstName = "Ada";
var lastLetterOfFirstName = firstName[firstName.length - 1];

// Setup
var lastName = "Lovelace";

// Only change code below this line.
var lastLetterOfLastName = lastName;
```

### Answers

```javascript
// Example
var firstName = "Ada";
var lastLetterOfFirstName = firstName[firstName.length - 1];

// Setup
var lastName = "Lovelace";

// Only change code below this line.
var lastLetterOfLastName = lastName[lastName.length - 1];
```