Given two arrays of strings a1 and a2 return a sorted array r in lexicographical order of the strings of a1 which are substrings of strings of a2.

#Example 1: a1 = ["arp", "live", "strong"]

a2 = ["lively", "alive", "harp", "sharp", "armstrong"]

returns ["arp", "live", "strong"]

#Example 2: a1 = ["tarp", "mice", "bull"]

a2 = ["lively", "alive", "harp", "sharp", "armstrong"]

returns []

Notes:
Arrays are written in "general" notation. See "Your Test Cases" for examples in your language.

In Shell bash a1 and a2 are strings. The return is a string where words are separated by commas.

Beware: r must be without duplicates.

### Thinking

1. I create a array store who receive the match
2. Loop over both array and push inside the match array those who array2 include the value of the array1 element
3. I filter each element to make sure they are unique by using indexOf. This way the only one who pass is the first value.
4. I finish with the sort function. I could have make it one line but I want to make sure all of the value is lowerCase.

### Solution

```js
function inArray(array1,array2){
  const _match = []

  for (let i = 0; i < array2.length; i++) {
    for (let x = 0; x < array1.length; x++) {
      const a1El = array1[x];
      const a2El = array2[i];
      if (a2El.includes(a1El)) {
        _match.push(a1El)
      }
    }
  }

  return _match.filter((el, i, a) => a.indexOf(el) === i).sort((a, b) => {
    const _a = a.toLowerCase();
    const _b = b.toLowerCase();

    if (_a < _b) {
      return -1;
    }

    return 1;
  })
}
```
