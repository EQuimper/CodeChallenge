You get an array of numbers, return the sum of all of the positives ones.

Example [1,-4,7,12] => 1 + 7 + 12 = 20

Note: array may be empty, in this case return 0.

``` js
function positiveSum(arr) {
  return arr.reduce((prev, current) => {
    if (current > 0) {
      return prev + current;
    }

    return prev
  }, 0)
}
```

1. Sum using reduce who start a 0
2. if the current is bigger than 0 than mean is position so add the value to the prev one
3. if not return the prev
