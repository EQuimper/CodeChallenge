# Iterate over Arrays with map

The `map` method is a convenient way to iterate through arrays.
Here's an example usage:

```javascript
var timesFour = oldArray.map(function(val){
  return val * 4;
});
```

The map method will iterate through every element of the array,
creating a new array with values that have been modified by the
callback function, and return it. Note that it does not modify
the original array.

In our example the callback only uses the value of the array
element (the val argument) but your callback can also include
arguments for the index and array being acted on.

Use the `map` function to add 3 to every value in the variable `oldArray`,
and save the results into variable newArray. oldArray should not change.

### Before

```javascript
var oldArray = [1,2,3,4,5];

// Only change code below this line.

var newArray = oldArray;
```

### Answers

```javascript
var oldArray = [1,2,3,4,5];

// Only change code below this line.
var addThree = oldArray.map(function(val) {
  return val + 3;
});

var newArray = addThree;
```