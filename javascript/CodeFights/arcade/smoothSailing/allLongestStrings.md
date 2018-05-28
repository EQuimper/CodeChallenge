Given an array of strings, return another array containing all of its longest strings.

Example

For inputArray = ["aba", "aa", "ad", "vcd", "aba"], the output should be
allLongestStrings(inputArray) = ["aba", "vcd", "aba"].

```js
function allLongestStrings(inputArray) {
    let longest = 0;

    for (let i = 0; i < inputArray.length; i++) {
        if (inputArray[i].length > longest) {
            longest = inputArray[i].length
        }
    }

    return inputArray.filter(el => el.length === longest);
}
```

1. We need first the length number of the biggest number. I do this by looping each element and get the bigger one
2. We filter each element to get those who have same length as the biggest one
