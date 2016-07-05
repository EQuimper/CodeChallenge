# Accessing Objects Properties with the Dot Operator

There are two ways to access the properties of an object: the
dot operator (`.`) and `bracket notation` (`[]`), similar to an `array`.

The dot operator is what you use when you
know the name of the property you're trying to access ahead of time.

Here is a sample of using the dot operator (`.`) to read an object property:

```javascript
var myObj = {
  prop1: "val1",
  prop2: "val2"
};
var prop1val = myObj.prop1; // val1
var prop2val = myObj.prop2; // val2
```

## Instructions
 - Read in the property values of `testObj` using `dot` notation. Set the
 variable `hatValue` equal to the `object` property hat and set the
 variable `shirtValue` equal to the `object` property `shirt`.

### Before

```javascript
// Setup
var testObj = {
  "hat": "ballcap",
  "shirt": "jersey",
  "shoes": "cleats"
};

// Only change code below this line

var hatValue = testObj;      // Change this line
var shirtValue = testObj;    // Change this line
```

### After

```javascript
// Setup
var testObj = {
  "hat": "ballcap",
  "shirt": "jersey",
  "shoes": "cleats"
};

// Only change code below this line

var hatValue = testObj.hat;      // Change this line
var shirtValue = testObj.shirt;    // Change this line
```