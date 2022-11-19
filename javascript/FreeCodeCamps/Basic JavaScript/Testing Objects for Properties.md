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

The first ```hasOwnProperty``` returns ```true```, while the second returns ```false```.

## Instructions
 - Modify the function `checkObj` to test if an object passed to the function (`obj`) contains a specific property (`checkProp`). 
 If the property is found, return that property's value. If not, return `"Not Found"`.

### Before

```javascript
function checkObj(obj, checkProp) {
  // Only change code below this line
  return "Change Me!";
  // Only change code above this line
}
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
