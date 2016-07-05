# Add New Properties to a JavaScript Object

You can add `new` properties to existing JavaScript `objects` the same
way you would modify them.

Here's how we would add a `bark` property to `ourDog`:

```javascript
ourDog.bark = "bow-wow";
```

or

```javascript
ourDog["bark"] = "bow-wow";
```

Now when we evaluate `ourDog.bark`, we'll get his bark, `bow-wow`.

## Instructions
 - Add a `bark` property to `myDog` and set it to a dog `sound`, such as
 "woof". You may use either `dot` or `bracket notation`.

### Before

```javascript
// Example
var ourDog = {
  "name": "Camper",
  "legs": 4,
  "tails": 1,
  "friends": ["everything!"]
};

ourDog.bark = "bow-wow";

// Setup
var myDog = {
  "name": "Happy Coder",
  "legs": 4,
  "tails": 1,
  "friends": ["Free Code Camp Campers"]
};

// Only change code below this line.
```

### Answers

```javascript
// Example
var ourDog = {
  "name": "Camper",
  "legs": 4,
  "tails": 1,
  "friends": ["everything!"]
};

ourDog.bark = "bow-wow";

// Setup
var myDog = {
  "name": "Happy Coder",
  "legs": 4,
  "tails": 1,
  "friends": ["Free Code Camp Campers"]
};

// Only change code below this line.

myDog.bark = "woof"; // Dot notation
myDog["bark"] = "woof"; // Bracket notation
```