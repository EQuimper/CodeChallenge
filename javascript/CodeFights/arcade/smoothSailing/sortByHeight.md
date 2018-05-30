Some people are standing in a row in a park. There are trees between them which cannot be moved. Your task is to rearrange the people by their heights in a non-descending order without moving the trees.

Example

For a = [-1, 150, 190, 170, -1, -1, 160, 180], the output should be
sortByHeight(a) = [-1, 150, 160, 170, -1, -1, 180, 190]

```js
function sortByHeight(a) {
  const nums = a.filter(el => el !== -1).sort((a, b) => a - b);
  const result = [];

  for (let i = 0; i < a.length; i++) {
    if (a[i] === -1) {
      result.push(-1);
    } else {
      result.push(nums[0]);
      nums.splice(0, 1);
    }
  }

  return result;
}
```

1. We create an array who have each number who is not -1 plus also this one get sort already
2. We create an array who gonna receive the order we need
3. we loop over the argument a and we check if this is === -1.
4. if this is equal we push -1 right away
5. if not we push the first element in the nums array who is the filter one and after we remove this element from this one. So we dont have in double.
