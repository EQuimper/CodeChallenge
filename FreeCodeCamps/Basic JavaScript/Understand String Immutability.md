# Understand String Immutability

In JavaScript, `String` values are `immutable`, which means that
they cannot be altered once created.

For example, the following code:

```javascript
var myStr = "Bob";
myStr[0] = "J";
```
cannot change the value of `myStr` to "Job", because the contents
of `myStr` cannot be altered. Note that this does not mean that `myStr`
cannot be changed, just that the individual characters of a string
literal cannot be changed. The only way to change myStr would be to assign
it with a new string, like this:

```javascript
var myStr = "Bob";
myStr = "Job";
```

## Instructions
 - Correct the assignment to `myStr` to achieve the desired effect.

### Before

```javascript
// Setup
var myStr = "Jello World";

// Only change code below this line

myStr[0] = "H"; // Fix Me
```

### Answers

```javascript
// Setup
var myStr = "Jello World";

// Only change code below this line

var newStr = myStr.slice(1);

myStr = 'H' + newStr;
```

### Thinking

Here I got a problem with the string so I try to figured out a answer.

1. I want to remove the first letter of `myStr` so I make a new var and use `.slice`
2. I push a new letter to `myStr`