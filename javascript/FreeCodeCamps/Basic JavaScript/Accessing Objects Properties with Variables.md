# Accessing Objects Properties with Variables

Another use of `bracket notation` on `objects` is to use a `variable` to
access a property. This can be very useful for `iterating through` lists of
the object properties or for doing the lookup.

Here is an example of using a variable to access a property:

```javascript
var someProp = "propName";
var myObj = {
  propName: "Some Value"
}
myObj[someProp]; // "Some Value"
```

Here is one more:

```javascript
var myDog = "Hunter";
var dogs = {
  Fido: "Mutt",
  Hunter: "Doberman",
  Snoopie: "Beagle"
}
var breed = dogs[myDog];
console.log(breed);// "Doberman"
```

Note that we do not use `quotes` around the `variable` name when using
it to access the property because we are using the `value` of the `variable`,
not the name

## Instructions

 - Use the `playerNumber` variable to lookup player 16 in `testObj` using
` bracket notation`.

### Before

```javascript
// Setup
var testObj = {
  12: "Namath",
  16: "Montana",
  19: "Unitas"
};

// Only change code below this line;

var playerNumber;       // Change this Line
var player = testObj;   // Change this Line
```

### Answers

```javascript
// Setup
var testObj = {
  12: "Namath",
  16: "Montana",
  19: "Unitas"
};

// Only change code below this line;

var playerNumber = 16;       // Change this Line
var player = testObj[playerNumber];   // Change this Line
```