Write a function that takes two or more arrays and returns a new array of unique values in the order of the original provided arrays.

In other words, all values present from all arrays should be included in their original order, but with no duplicates in the final array.

The unique numbers should be sorted by their original order, but the final array should not be sorted in numerical order.

```js
function uniteUnique(arr) {
  // make an array with the arguments object
  var args = Array.from(arguments);
  // concat each array in args to make one
  var arrConcat = args.reduce(function(a, b) {
    return a.concat(b);
  }, []);
  // function for return unique element
  var uniqueArr = function(a) {
    // reduce the array give
    return a.reduce(function(b, c) {
      // find if this is unique
      if (b.indexOf(c) < 0) {
        // if unique push it in the array
        b.push(c);
      }
      // return the array
      return b;
    }, []);
  };

  return uniqueArr(arrConcat);
}

uniteUnique([1, 3, 2], [5, 2, 1, 4], [2, 1]);
```
