# Constructing Strings with Variables

Sometimes you will need to build a string, [Mad Libs](https://en.wikipedia.org/wiki/Mad_Libs) style.
By using the concatenation operator (`+`), you can insert one or more
variables into a string you're building.

## Instructions
 - Set myName to a string equal to your name and build myStr with myName
 between the strings `"My name is "` and `" and I am swell!"`

### Before

```javascript
// Example
var ourName = "Free Code Camp";
var ourStr = "Hello, our name is " + ourName + ", how are you?";

// Only change code below this line
var myName;
var myStr;
```

### Answers

```javascript
// Example
var ourName = "Free Code Camp";
var ourStr = "Hello, our name is " + ourName + ", how are you?";

// Only change code below this line
var myName = 'Emanuel';
var myStr = 'My name is ' + myName + ' and I am swell!';
```