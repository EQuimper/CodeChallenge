# Where do I belong

Return the lowest index at which a value (second argument) should be
inserted into an array (first argument) once it has been sorted.
The returned value should be a number.

For example, getIndexToIns(`[1,2,3,4], 1.5`) should return 1 because it
is greater than 1 (index 0), but less than 2 (index 1).

Likewise, getIndexToIns(`[20,3,5], 19`) should return 2 because once the
array has been sorted it will look like [3,5,20] and 19 is less than
20 (index 2) and greater than 5 (index 1).

### Before

```javascript
function getIndexToIns(arr, num) {
  // Find my place in this sorted array.
  return num;
}

getIndexToIns([40, 60], 50);
```

### Answers

```javascript
function getIndexToIns(arr, num) {

  if(arr.indexOf(num) === -1) {
    arr.push(num);
  }

  var newArr = arr;

  function sortNumber(a,b) {
    return a - b;
  }

  newArr.sort(sortNumber);

  return newArr.indexOf(num);

}

getIndexToIns([3, 10, 5], 3);
```

### Thinking

1. Check if the args num is inside the arr. If yes we don't add it in the arr
2. We create a function who gonna be a helpers to your sort cause of problem
with double digits number
3. We sort the array
4. Return the `indexOf` of the args num