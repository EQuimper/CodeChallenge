After they became famous, the CodeBots all decided to move to a new building and live together. The building is represented by a rectangular matrix of rooms. Each cell in the matrix contains an integer that represents the price of the room. Some rooms are free (their cost is 0), but that's probably because they are haunted, so all the bots are afraid of them. That is why any room that is free or is located anywhere below a free room in the same column is not considered suitable for the bots to live in.

Help the bots calculate the total price of all the rooms that are suitable for them.

Example

For
```
matrix = [[0, 1, 1, 2],
          [0, 5, 0, 0],
          [2, 0, 3, 3]]
```
the output should be
matrixElementsSum(matrix) = 9.


```js
function matrixElementsSum(matrix) {
  for(let i = 0; i < matrix.length; i++) {
    for(let x = 0; x < matrix[0].length; x++){
      if (matrix[i][x] === 0 && i < matrix.length - 1) {
        matrix[i + 1][x] = 0;
      }
    }
  }

  return matrix.reduce((prev, current) => {
    current.forEach(el => {
      prev += el
    })

    return prev
  }, 0)
}
```

1. We need to loop over each array first.
2. If element === 0 and the second array exist we put this element as 0
3. We sum all the value to a single one
