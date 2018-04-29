Given an array of integers, find the pair of adjacent elements that has the largest product and return that product.

```js
function adjacentElementsProduct(inputArray) {
  let maxAdjacent = inputArray[0] * inputArray[1];

  for (let i = 1; i < inputArray.length - 1; i++) {
    if (getAdjacent(inputArray, i) > maxAdjacent) {
      maxAdjacent = getAdjacent(inputArray, i);
    }
  }

  return maxAdjacent;
}

function getAdjacent(inputArray, index) {
  return inputArray[index] * inputArray[index + 1];
}
```
