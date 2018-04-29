A Narcissistic Number is a number which is the sum of its own digits, each raised to the power of the number of digits.

For example, take 153 (3 digits):
```
1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153
```
and 1634 (4 digits):
```
1^4 + 6^4 + 3^4 + 4^4 = 1 + 1296 + 81 + 256 = 1634
```
The Challenge:

Your code must return true or false depending upon whether the given number is a Narcissistic number.

Error checking for text strings or other invalid inputs is not required, only valid integers will be passed into the function.

### Thinking

1. I make an array with the value. But I need to transform it to string first to make it happen. I build an array so I can iterate over each element.
2. I create a variable `_result` who is a data store of the value.
3. Loop over each value and need to transform back the value to a number so I can make some Math with it.
4. I use Math.pow you give me the power of a number. The second arguments I use is the length of the value. So if the number is `123` the pow become `1^3 2^3 3^3`
5. I make sure to add the value to the _result data store.
6. We need to return only a boolean about if this is match or not.

### Solution

```js
function narcissistic(value) {
  const _value = String(value).split('');

  let _result = 0;

  for (ch of _value) {
    const num = parseInt(ch, 0)

    _result += Math.pow(num, _value.length);
  }

  return _result === value;
}
```
