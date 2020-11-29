# Testing Objects for Properties

Sometimes it is useful to check if the property of a given `object` exists
or not. We can use the `.hasOwnProperty(propname)` method of `objects` to
determine if that object has the given property name. `.hasOwnProperty()`
returns `true` or `false` `if` the property is found or not.

## Example

```javascript
var myObj = {
  top: "hat",
  bottom: "pants"
};
myObj.hasOwnProperty("top");    // true
myObj.hasOwnProperty("middle"); // false
```

## Instructions
 - Modify the function `checkObj` to test `myObj` for `checkProp`. If
 the property is found, return that property's value. If not, return "Not Found".

### Before

```javascript
// Setup
var myObj = {
  gift: "pony",
  pet: "kitten",
  bed: "sleigh"
};

function checkObj(checkProp) {
  // Your Code Here

  return "Change Me!";
}

// Test your code by modifying these values
checkObj("gift");
```

### Answers

```javascript
function checkObj(obj, checkProp) {
  // Only change code below this line
    if (obj.hasOwnProperty(checkProp)) {
   return obj[checkProp];
 } else {
   return "Not Found";
 }
  // Only change code above this line
}

```

### Thinking

Here because the args pass is a string, that mean he have quote. We need
to use bracket notation.
