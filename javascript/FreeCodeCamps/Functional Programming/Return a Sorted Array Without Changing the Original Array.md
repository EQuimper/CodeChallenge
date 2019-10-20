# Return a Sorted Array Without Changing the Original Array

Use the sort method in the nonMutatingSort function to sort the elements of an array in ascending order. The function should return a new array, and not mutate the globalArray variable.

### Before

```js
var globalArray = [5, 6, 3, 2, 9];
function nonMutatingSort(arr) {
  // Add your code below this line


  // Add your code above this line
}
nonMutatingSort(globalArray);
```

### Answer

```js
var globalArray = [5, 6, 3, 2, 9];
function nonMutatingSort(arr) {
  // Add your code below this line

  return [...arr].sort()
  // Add your code above this line
}
nonMutatingSort(globalArray);
```

### Thinking

1. We need a copy of the array before sorting, cause we do not want to mutated it. So by using `[...arr]` we make a copy and now we can sort.
