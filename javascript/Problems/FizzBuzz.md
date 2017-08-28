# FizzBuzz

1 to 100, print Fizz when you can divided by 3. Print Buzz when you can divided by 5
and FizzBuzz when you can do this with both 3 and 5.

### Answers

```javascript
for (var i = 1; i <= 100; i++) {
  if (i % 3 === 0 && i % 5 === 0) {
    console.log('FizzBuzz');
  } else if (i % 3 === 0) {
    console.log('Fizz');
  } else if (i % 5 === 0) {
    console.log('Buzz');
  } else {
    console.log(i);
  }
}
```

### Thinking

1. Loop to all number
2. See when divide by number if whe have 0
3. We need to start whit FizzBuzz if we want it to work