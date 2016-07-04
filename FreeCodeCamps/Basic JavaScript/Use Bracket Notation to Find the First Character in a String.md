# Use Bracket Notation to Find the First Character in a String

`Bracket notation` is a way to get a character at a specific index
within a string.

Most modern programming languages, like JavaScript, **don't start counting
at 1 like humans do. They start at 0.** This is referred to as Zero-based
indexing.

For example, the character at index 0 in the word "Charles" is "C". So if `var
firstName = "Charles"`, you can get the value of the first letter
of the string by using `firstName[0]`.

## Instructions
 - Use bracket notation to find the first character in the `lastName`
 variable and assign it to `firstLetterOfLastName`.

## Hint
 - Try looking at the `firstLetterOfFirstName` variable declaration if you
 get stuck.

### Before

```javascript
// Example
var firstLetterOfFirstName = "";
var firstName = "Ada";

firstLetterOfFirstName = firstName[0];

// Setup
var firstLetterOfLastName = "";
var lastName = "Lovelace";

// Only change code below this line
firstLetterOfLastName = lastName;
```

### Answers

```javascript
// Example
var firstLetterOfFirstName = "";
var firstName = "Ada";

firstLetterOfFirstName = firstName[0];

// Setup
var firstLetterOfLastName = "";
var lastName = "Lovelace";

// Only change code below this line
firstLetterOfLastName = lastName[0];
```