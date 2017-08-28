You are given a two-digit integer n. Return the sum of its digits.

Example

For n = 29, the output should be
addTwoDigits(n) = 11.

Input/Output

[time limit] 4000ms (js)
[input] integer n

A positive two-digit integer.

Constraints:
10 ≤ n ≤ 99.

[output] integer

The sum of the first and second digits of the input number.

## My Answer

```js
'use strict'

function addTwoDigits(n) {
  const arr = (n).toString(10).split("").map(Number);
  
  return arr[0] + arr[1];
}

addTwoDigits(25);
```

## My Thinking

- I make a array but before I need to make it a string cause we don't want them to add but more concat.
- They told in the task we have only 2 digits so I take 2 index.
