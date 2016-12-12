If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Finish the solution so that it returns the sum of all the multiples of 3 or 5 below the number passed in.

> Note: If the number is a multiple of both 3 and 5, only count it once.

```js
function solution(number){
  let arr = [];
  let multiple = [];
  
  // push all num between 0 and the number to arr
  for (let i = 0; i < number; i++) {
    arr.push(i);
  }
  
  // if number in arr is a multiple or 3 or 5 push it to multiple
  for (let i = 0; i < arr.length; i++) {
    if (i % 3 === 0 || i % 5 === 0) {
      multiple.push(i);
    }
  }
  
  // sum the total of the multiple array
  const sumTotal = multiple.reduce((sum, num) => sum + num, 0);
  
  return sumTotal;
}
```
