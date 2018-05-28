Ticket numbers usually consist of an even number of digits. A ticket number is considered lucky if the sum of the first half of the digits is equal to the sum of the second half.

Given a ticket number n, determine if it's lucky or not.

Example

For n = 1230, the output should be
isLucky(n) = true;
For n = 239017, the output should be
isLucky(n) = false.

```js
function isLucky(n) {
    const arr = String(n).split('');

    const firstHalf = sum(arr.slice(0, arr.length / 2));
    const secondHalf = sum(arr.slice(arr.length / 2));

    return firstHalf === secondHalf;
}

function sum(arr) {
  return arr.reduce((prev, current) => {
    const _current = parseInt(current, 0);

    return prev += _current
  }, 0)
}
```

1. Split the n to an array and make sure this is a String first.
2. Create a function sum who take and array and use reduce to get a sum, make all element a number first before addition them.
3. return true or false if both value equal
