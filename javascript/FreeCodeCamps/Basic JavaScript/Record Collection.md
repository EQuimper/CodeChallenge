# Record Collection

You are given an object literal representing a part of your musical album collection. Each album has a unique id number as its key and several other properties. Not all albums have complete information. <br>
<br>
You start with an `updateRecords` function that takes an object literal, `records`, containing the musical album collection, an `id`, a `prop` (like `artist` or `tracks`), and a `value`. Complete the function using the rules below to modify the object passed to the function.<br>
<br>
Your function must always return the entire record collection object.<br>
- If `prop` isn't `tracks` and value isn't an empty string, update or set that album's `prop` to `value`.
- If `prop` is` tracks` but the album doesn't have a `tracks` property, create an empty array and add `value` to it.
- If `prop` is `tracks` and `value` isn't an empty string, add `value` to the end of the album's existing `tracks` array.
- If `value` is an empty string, delete the given `prop` property from the album.
Note: A copy of the `recordCollection` object is used for the tests.<br>
<br>
## Hints
 - Use `bracket notation` when accessing object properties with variables.

`Push` is an array method you can read about on Mozilla Developer Network.

You may refer back to Manipulating Complex ObjectsIntroducing JavaScript Object Notation (JSON) for a refresher.

### Before

```javascript
// Setup
const recordCollection = {
  2548: {
    albumTitle: 'Slippery When Wet',
    artist: 'Bon Jovi',
    tracks: ['Let It Rock', 'You Give Love a Bad Name']
  },
  2468: {
    albumTitle: '1999',
    artist: 'Prince',
    tracks: ['1999', 'Little Red Corvette']
  },
  1245: {
    artist: 'Robert Palmer',
    tracks: []
  },
  5439: {
    albumTitle: 'ABBA Gold'
  }
};

// Only change code below this line
function updateRecords(records, id, prop, value) {
  return records;
}

updateRecords(recordCollection, 5439, 'artist', 'ABBA');
```

### Answers

```javascript
// Setup
const recordCollection = {
  2548: {
    albumTitle: 'Slippery When Wet',
    artist: 'Bon Jovi',
    tracks: ['Let It Rock', 'You Give Love a Bad Name']
  },
  2468: {
    albumTitle: '1999',
    artist: 'Prince',
    tracks: ['1999', 'Little Red Corvette']
  },
  1245: {
    artist: 'Robert Palmer',
    tracks: []
  },
  5439: {
    albumTitle: 'ABBA Gold'
  }
};

// Only change code below this line
function updateRecords(records, id, prop, value) {

 if (value === '') {
    delete records[id][prop];
  } else if (prop === 'tracks') {
    records[id][prop] = records[id][prop] || []; // this is called shortcircuit evaluation, see below for explanation
    records[id][prop].push(value);
  } else {
    records[id][prop] = value;
  }
  return records;
}
updateRecords(recordCollection, 5439, 'artist', 'ABBA');
```

### Thinking

1. The more easy is to delete the prop if no value
2. If the prop is not tracks it's easy to add the value to this prop
3. Here we need to find if the object have a property tracks if not we create
this prop with a empty array and we push it the value. I use `.hasOwnProperty` for
doing this.
