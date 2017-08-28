# Profile Lookup

We have an `array` of `objects` representing different people in our
contacts lists.

A `lookUpProfile` function that takes `firstName` and a `property` (prop) as
arguments has been pre-written for you.

The function should check if `firstName` is an actual
contact's `firstName` and the given property (`prop`) is a property of
that contact.

If both are true, then return the "value" of that property.

If `firstName` does not correspond to any contacts then return
"No such contact"

If prop does not correspond to any valid properties then return
"No such property"

### Before

```javascript

//Setup
var contacts = [
    {
        "firstName": "Akira",
        "lastName": "Laine",
        "number": "0543236543",
        "likes": ["Pizza", "Coding", "Brownie Points"]
    },
    {
        "firstName": "Harry",
        "lastName": "Potter",
        "number": "0994372684",
        "likes": ["Hogwarts", "Magic", "Hagrid"]
    },
    {
        "firstName": "Sherlock",
        "lastName": "Holmes",
        "number": "0487345643",
        "likes": ["Intriguing Cases", "Violin"]
    },
    {
        "firstName": "Kristian",
        "lastName": "Vos",
        "number": "unknown",
        "likes": ["Javascript", "Gaming", "Foxes"]
    }
];


function lookUpProfile(firstName, prop){
// Only change code below this line

// Only change code above this line
}

// Change these values to test your function
lookUpProfile("Akira", "likes");
```

### Answers

```javascript

//Setup
var contacts = [
    {
        "firstName": "Akira",
        "lastName": "Laine",
        "number": "0543236543",
        "likes": ["Pizza", "Coding", "Brownie Points"]
    },
    {
        "firstName": "Harry",
        "lastName": "Potter",
        "number": "0994372684",
        "likes": ["Hogwarts", "Magic", "Hagrid"]
    },
    {
        "firstName": "Sherlock",
        "lastName": "Holmes",
        "number": "0487345643",
        "likes": ["Intriguing Cases", "Violin"]
    },
    {
        "firstName": "Kristian",
        "lastName": "Vos",
        "number": "unknown",
        "likes": ["Javascript", "Gaming", "Foxes"]
    }
];


function lookUpProfile(firstName, prop){

  var msg = 'No such contact';

  for (var p in contacts) {
    if (contacts[p].firstName === firstName && contacts[p].hasOwnProperty(prop)) {
      msg = contacts[p][prop];
    } else if (!contacts[p].hasOwnProperty(prop)) {
      msg = 'No such property';
    }
  }

  return msg;
// Only change code above this line
}

// Change these values to test your function
lookUpProfile("Harry", "likes");
```

### Thinking

1. We setup a var msg and give in the last value we want to check
2. We loop the property inside contacts
3. We check if the firstName is one of the one inside the arr and if yes we need to
have the prop
4. Here we have the one who the firstName is true but they didn't pass
cause they don't have the prop so here we return the 'No such property'