# Seek and Destroy

You will be provided with an initial array (the first argument in
the destroyer function), followed by one or more arguments.
Remove all elements from the initial array that are of the
same value as these arguments.

### Before

```javascript
function destroyer(arr) {
  // Remove all the values
  return arr;
}

destroyer([1, 2, 3, 1, 2, 3], 2, 3);
```

### Answers

```javascript
function destroyer(arr) {
  var check = [];

  for (var i = 1; i < arguments.length; i++) {
    check.push(arguments[i]);
  }

  var itemsNotDestroy = arguments[0].filter(function(item) {
    var itemPass = true;
    for (var i = 0; i < check.length; i++) {
      if (item === check[i]) {
        itemPass = false;
      }
    }
    return itemPass;
  });

  return itemsNotDestroy;
}

destroyer([1, 2, 3, 5, 1, 2, 3], 2, 3);
```

### Thinking

1. We create a arr you take the arguments we remove
2. We filter the item in the arguments 0 and filter it
3. We just return the one who pass the filter