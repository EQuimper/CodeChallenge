Given two strings, find the number of common characters between them.

Example

For s1 = "aabcc" and s2 = "adcaa", the output should be
commonCharacterCount(s1, s2) = 3.

Strings have 3 common characters - 2 "a"s and 1 "c".


```js
function commonCharacterCount(s1, s2) {
    let amount = 0;

    const arr1 = s1.split('');
    const arr2 = s2.split('');

    for (let i = 0; i < arr1.length; i++) {
        const index = arr2.indexOf(arr1[i]);
        if (index !== -1) {
            arr2.splice(index, 1);
            amount += 1
        }
    }

    return amount
}

```

1. First split both string in array.
2. Loop over the first one and get the index of it if this match
3. if match add 1 to the amount variable and remove the index this way we dont count in double
