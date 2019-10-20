# Use Higher-Order Functions map, filter, or reduce to Solve a Complex Problem

We have defined a function named squareList. You need to complete the code for the squareList function using any combination of map(), filter(), and reduce() so that it returns a new array containing only the square of only the positive integers (decimal numbers are not integers) when an array of real numbers is passed to it. An example of an array containing only real numbers is [-3, 4.8, 5, 3, -3.2].

Note: Your function should not use any kind of for or while loops or the forEach() function.


### Before

```js
const squareList = (arr) => {
  // only change code below this line
  return arr;
  // only change code above this line
};

// test your code
const squaredIntegers = squareList([-3, 4.8, 5, 3, -3.2]);
console.log(squaredIntegers);
```

### After

```js
const squareList = (arr) => {
  // only change code below this line
  return arr.filter(el => el > 0 && el % 1 === 0).map(el => el * el);
  // only change code above this line
};

// test your code
const squaredIntegers = squareList([-3, 4.8, 5, 3, -3.2]);
console.log(squaredIntegers);
```

### Thinking

1. First we need to remove the non-positive number so filter element who are smaller than 0
2. We also need to make sure those number are not decimal. The way to deal with that is checking with `%` and 1 to see if that equal 0. The reminder of a real
integer will always be 0 so the others are decimals number and should be remove.
3. We map over the rest and square those numbers.
