n children have got m pieces of candy. They want to eat as much candy as they can, but each child must eat exactly the same amount of candy as any other child. Determine how many pieces of candy will be eaten by all the children together. Individual pieces of candy cannot be split.

Example

For n = 3 and m = 10, the output should be
candies(n, m) = 9.

Each child will eat 3 pieces. So the answer is 9.

Input/Output

[time limit] 4000ms (js)
[input] integer n

The number of children.

Constraints:
1 ≤ n ≤ 10.

[input] integer m

The number of pieces of candy.

Constraints:
2 ≤ m ≤ 100.

[output] integer

The total number of pieces of candy the children will eat provided they eat as much as they can and all children eat the same amount.

## My Answer

```js
function candies(n, m) {
    'use strict';
    
    if (m % n === 0) {
        return m / n;
    }
    
    const tempSub = m - (m % n);
    
    const numEach = tempSub / n;
    
    return numEach * n;
}

candies(3, 10);
```
