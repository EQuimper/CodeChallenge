Given the string, check if it is a palindrome.

```js
function checkPalindrome(inputString) {
  let rvStr = inputString.split('').reverse().join('');

  let isPalindrome = false;

  if (rvStr === inputString) {
    isPalindrome = true;
  }

  return isPalindrome;
}
```
