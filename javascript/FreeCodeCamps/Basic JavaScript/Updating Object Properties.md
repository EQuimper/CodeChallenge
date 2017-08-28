# Updating Object Properties

After you've created a JavaScript `object`, you can update its
properties at any time just like you would update any other `variable`.
You can use either `dot` or `bracket notation` to update.

For example, let's look at `ourDog`:

```javascript
var ourDog = {
  "name": "Camper",
  "legs": 4,
  "tails": 1,
  "friends": ["everything!"]
};
```

Since he's a particularly happy dog, let's change his name to `Happy Camper`.
Here's how we update his object's name property:

```javascript
ourDog.name = "Happy Camper"; or

ourDog["name"] = "Happy Camper";
```

Now when we evaluate `ourDog.name`, instead of getting `Camper`,
we'll get his new name, `Happy Camper`.

## Instructions
 - Update the `myDog` object's name property. Let's change her name from
 `Coder` to `Happy Coder`. You can use either `dot` or `bracket notation`.

### Before

```javascript
// Example
var ourDog = {
  "name": "Camper",
  "legs": 4,
  "tails": 1,
  "friends": ["everything!"]
};

ourDog.name = "Happy Camper";

// Setup
var myDog = {
  "name": "Coder",
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

ourDog.name = "Happy Camper";

// Setup
var myDog = {
  "name": "Coder",
  "legs": 4,
  "tails": 1,
  "friends": ["Free Code Camp Campers"]
};

// Only change code below this line.

myDog.name = "Happy Coder"; // Dot notation
myDog["name"] = "Happy Coder"; // Bracket notation
```