# Sum All Numbers in a Range

We'll pass you an array of two numbers. Return the sum of those two numbers and all numbers between them.

The lowest number will not always come first.

### Before

```javascript
function sumAll(arr) {
  return 1;
}

sumAll([1, 4]);
```

### Answer

```javascript
function sumAll(arr) {

  var newArr = [];

  var min = Math.min(arr[0], arr[1]);
  var max = Math.max(arr[0], arr[1]);

  for (var x = min; x <= max; x++) {
    newArr.push(x);
  }

  var sum = newArr.reduce(function(prev, curr) {
    return prev + curr;
  });

  return sum;

}

sumAll([1, 4]);
```

### Thinking

1. Create a empty who gonna receive the number between the min and max
2. Find with both argument who is the max and who is the min
3. Loop between both number and push it to the empty array
4. Use `reduce()` who make for us the job of sum all the number
