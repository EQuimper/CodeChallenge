# Manipulate Arrays With unshift

Not only can you `shift` elements off of the beginning of an array,
you can also unshift elements to the beginning of an array i.e.
add elements in front of the array.

`.unshift()` works exactly like `.push()`, but instead of adding the
element at the end of the array, `unshift()` **adds the element at the
beginning of the array**.

## Instructions
 - Add `["Paul",35]` to the beginning of the `myArray` variable using `unshift()`.

### Before

```javascript
// Example
var ourArray = ["Stimpson", "J", "cat"];
ourArray.shift(); // ourArray now equals ["J", "cat"]
ourArray.unshift("Happy");
// ourArray now equals ["Happy", "J", "cat"]

// Setup
var myArray = [["John", 23], ["dog", 3]];
myArray.shift();

// Only change code below this line.
```

### Answers

```javascript
// Example
var ourArray = ["Stimpson", "J", "cat"];
ourArray.shift(); // ourArray now equals ["J", "cat"]
ourArray.unshift("Happy");
// ourArray now equals ["Happy", "J", "cat"]

// Setup
var myArray = [["John", 23], ["dog", 3]];
myArray.shift();

// Only change code below this line.
myArray.unshift(["Paul", 35]);
```