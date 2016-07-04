# Understanding Uninitialized Variables

When JavaScript variables are declared, they have an initial value of undefined. If you do a mathematical
operation on an undefined variable your result will be `NaN` which means "Not a Number".
If you concatenate a string with an undefined variable, you will get a literal string of "undefined".

## Instructions
 - Initialize the three variables a, b, and c with 5, 10, and "I am a" respectively so that they will not be undefined.

### Answers

```javascript
// Initialize these three variables
var a = 5;
var b = 10;
var c = 'I am a';

// Do not change code below this line

a = a + 1;
b = b + 5;
c = c + " String!";

```

