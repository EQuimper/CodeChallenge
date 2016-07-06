# Return Largest Numbers in Arrays

Return an array consisting of the largest number
from each provided sub-array. For simplicity, the provided array
will contain exactly 4 sub-arrays.

Remember, you can iterate through an array with a simple for
loop, and access each member with array syntax `arr[i]`.

### Before

```javascript
function largestOfFour(arr) {
  // You can do this!
  return arr;
}

largestOfFour([[4, 5, 1, 3], [13, 27, 18, 26], [32, 35, 37, 39], [1000, 1001, 857, 1]]);
```

### Answers

```javascript
function largestOfFour(arr) {

  var arrMax = [];

  for (var i = 0; i < arr.length; i++) {

    var max = 0;

    for (var x = 0; x < arr[i].length; x++) {
      if (arr[i][x] > max) {
        max = arr[i][x];
      }
    }

    arrMax.push(max);
  }

  return arrMax;

}

largestOfFour([[4, 5, 1, 3], [13, 27, 18, 26], [32, 35, 37, 39], [1000, 1001, 857, 1]]);
```

### Thinking

1. We setup the empty arr who we gonna push the answer
2. We for loop the arr and we setup a var max for this one for get
the bigger number for his array
3. We for loop in the sub array and we check for the bigger one
4. We push all the bigger number inside your array