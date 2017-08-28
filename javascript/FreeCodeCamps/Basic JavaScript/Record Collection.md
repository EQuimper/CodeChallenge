# Record Collection

You are given a `JSON` object representing (a small part of) your record collection.
Each album is identified by a unique `id` number (`its key`) and has
several properties. Not all albums have complete information.

Write a function which takes an `id`, a property (prop), and a value.

For the given `id` in collection:

If prop does not contain the key `tracks`, then update or set
the value for that incomplete prop.

If prop does not contain the key `tracks` before you update
it, create an empty array before pushing a track to it.

If prop does contain the key `tracks` and its value is non-blank,
then push the value onto the end of its existing tracks array.

If value is blank, delete that prop.

Always return the entire collection object.

## Hints
 - Use `bracket notation` when accessing object properties with variables.

`Push` is an array method you can read about on Mozilla Developer Network.

You may refer back to Manipulating Complex ObjectsIntroducing JavaScript Object Notation (JSON) for a refresher.

### Before

```javascript
// Setup
var collection = {
    "2548": {
      "album": "Slippery When Wet",
      "artist": "Bon Jovi",
      "tracks": [
        "Let It Rock",
        "You Give Love a Bad Name"
      ]
    },
    "2468": {
      "album": "1999",
      "artist": "Prince",
      "tracks": [
        "1999",
        "Little Red Corvette"
      ]
    },
    "1245": {
      "artist": "Robert Palmer",
      "tracks": [ ]
    },
    "5439": {
      "album": "ABBA Gold"
    }
};
// Keep a copy of the collection for tests
var collectionCopy = JSON.parse(JSON.stringify(collection));

// Only change code below this line
function updateRecords(id, prop, value) {


  return collection;
}

// Alter values below to test your code
updateRecords(5439, "artist", "ABBA");
```

### Answers

```javascript
// Setup
var collection = {
    "2548": {
      "album": "Slippery When Wet",
      "artist": "Bon Jovi",
      "tracks": [
        "Let It Rock",
        "You Give Love a Bad Name"
      ]
    },
    "2468": {
      "album": "1999",
      "artist": "Prince",
      "tracks": [
        "1999",
        "Little Red Corvette"
      ]
    },
    "1245": {
      "artist": "Robert Palmer",
      "tracks": [ ]
    },
    "5439": {
      "album": "ABBA Gold"
    }
};
// Keep a copy of the collection for tests
var collectionCopy = JSON.parse(JSON.stringify(collection));

// Only change code below this line
function updateRecords(id, prop, value) {

  if (value === '') {
    delete collection[id][prop]; // If the value is empty remove the prop
  } else if (prop !== "tracks") {
    collection[id][prop] = value;
  } else {
    if (!collection[id].hasOwnProperty('tracks')) {
      collection[id].tracks = [];
      collection[id].tracks.push(value);
    } else {
      collection[id].tracks.push(value);
    }
  }
  return collection;
}

// Alter values below to test your code
updateRecords(2468, "tracks", "Free");
```

### Thinking

1. The more easy is to delete the prop if no value
2. If the prop is not tracks it's easy to add the value to this prop
3. Here we need to find if the object have a property tracks if not we create
this prop with a empty array and we push it the value. I use `.hasOwnProperty` for
doing this.