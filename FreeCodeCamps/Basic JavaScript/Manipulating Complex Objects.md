# Manipulating Complex Objects

JavaScript objects are flexible because they allow
for `Data Structures` with arbitrary combinations of `strings`, `numbers`,
`booleans`, `arrays`, `functions`, and `objects`.

Here is an example of a complex data structure:

```javascript
var ourMusic = [
  {
    "artist": "Daft Punk",
    "title": "Homework",
    "release_year": 1997,
    "formats": [
      "CD",
      "Cassette",
      "LP" ],
    "gold": true
  }
];
```

This is an `array` of `objects` and the object has various
pieces of metadata about an album. It also has a
nested formats `array`. Additional album records
could be added to the top level array.

## Note
 *You will need a comma in between objects with more than one object in the array.*

**JavaScript Object Notation** or `JSON` is a data interchange format
used to store data (source: json.org).

A property is the part of an object that associates a key
(either a `String` value or a `Symbol` value) and a value
(source: ecma-international.org/ecma-262/6.0/#sec-property). So, a
property consists of a `key` - value pair. (source: spacetelescope.github.io/understanding-json-schema/reference/object.html#properties).
Property keys (also known as names) should be in quotation marks.

Like `JavaScript Objects`, `JSON` is flexible because it is heterogeneous, meaning it permits `Data Structures` with arbitrary combinations of `strings`, `booleans`, `numbers`, `arrays`, and `objects`.

Here is an example of a `JSON` object:

```javascript
var ourMusic = [
  {
    "artist": "Daft Punk",
    "title": "Homework",
    "release_year": 1997,
    "formats": [
      "CD",
      "Cassette",
      "LP" ],
    "gold": true
  }
];
```

This is an `array` of `objects` and the `object` has various pieces of metadata about an album. It also has a nested formats array. Additional album records could be added to the top level array.

## Note
 *You will need to place a comma in between objects in JSON unless there is only one object in the array or containing object.*

## Instructions
 - Add a new album to the `myMusic` object. Add `artist` and `title` `strings`, `release_year` `number`, and a formats `array` of `strings`.

### Before

```javascript
var myMusic = [
  {
    "artist": "Billy Joel",
    "title": "Piano Man",
    "release_year": 1973,
    "formats": [
      "CS",
      "8T",
      "LP" ],
    "gold": true
  }
  // Add record here
];
```

### Answers

```javascript
var myMusic = [
  {
    "artist": "Billy Joel",
    "title": "Piano Man",
    "release_year": 1973,
    "formats": [
      "CS",
      "8T",
      "LP" ],
    "gold": true
  },
  // Add record here
  {
    "artist": "Bob",
    "title": "Hello",
    "release_year": 1950,
    "formats": [
      "CS",
      "8T"
    ]
  }
];
```