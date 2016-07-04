# Appending Variables to Strings

Just as we can build a string over multiple lines out of string
literals, we can also append variables to a string using the plus
equals (`+=`) operator.

## Instructions
 - Set someAdjective and append it to myStr using the `+=` operator.

### Before

```javascript
// Example
var anAdjective = "awesome!";
var ourStr = "Free Code Camp is ";
ourStr += anAdjective;

// Only change code below this line

var someAdjective;
var myStr = "Learning to code is ";
```

### Answers

```javascript
// Example
var anAdjective = "awesome!";
var ourStr = "Free Code Camp is ";
ourStr += anAdjective;

// Only change code below this line

var someAdjective = 'so nice';
var myStr = "Learning to code is ";
myStr += someAdjective;
```