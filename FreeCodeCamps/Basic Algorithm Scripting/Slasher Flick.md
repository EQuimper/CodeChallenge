# Slasher Flick

Return the remaining elements of an array after chopping
off `n` elements from the head.

The head means the beginning of the array, or the zeroth index.

### Before

```javascript
function slasher(arr, howMany) {
  // it doesn't always pay to be first
  return arr;
}

slasher([1, 2, 3], 2);
```

### Answers

```javascript
function slasher(arr, howMany) {
  var chop = [];

  chop = arr.splice(0, howMany);

  return arr;
}

slasher([1, 2, 3], 2);
```

### Thinking

1. Setup empty array
2. Use this array for get the item we remove
3. Return the arr who don't have anymore this item