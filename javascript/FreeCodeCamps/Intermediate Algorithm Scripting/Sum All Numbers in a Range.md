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
  let biggerNumber = Math.max(...arr);
  let smallerNumber = Math.min(...arr)

  let total = biggerNumber + smallerNumber

  for (let i = smallerNumber + 1; i < biggerNumber; i++) {
    total += i
  }

  return total
}

sumAll([1, 4]);
```

### Thinking

1. Get the biggest and the smallest number;
2. Create a var total who initialize with the biggestNumber sum to the smallest one
3. Loop between between the range and add this to the total
