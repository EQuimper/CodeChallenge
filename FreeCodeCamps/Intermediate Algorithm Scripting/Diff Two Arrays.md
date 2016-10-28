# Diff Two Arrays

Compare two `arrays` and return a new `array` with any items only found in one of the two given `arrays`, but not both. In other words, return the symmetric difference of the two `arrays`.

### Before

```js
function diffArray(arr1, arr2) {
  var newArr = [];
  // Same, same; but different.
  return newArr;
}

diffArray([1, 2, 3, 5], [1, 2, 3, 4, 5]);
```

### Answers

```js
function diffArray(arr1, arr2) {
  var newArr = [];
  
  newArr = arr1.concat(arr2); 
  
  function checkNum(num) {
    if (arr1.indexOf(num) === -1 || arr2.indexOf(num) === -1) {
      return num;
    }
  }
  
  return newArr.filter(checkNum);
  
}

diffArray([1, 2, 3, 5], [1, 2, 3, 4, 5]);
```

### Thinking

1. Create an empty array and passed it the value of the both array concat. That's gonna help alot if the value is in the same array.
2. Just create a function who check in a num is the indexOf and if the answer is no that give us -1.
3. Pass the action in the `filter()` method.
