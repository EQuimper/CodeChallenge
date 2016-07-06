# Title Case a Sentence

Return the provided `string` with the first letter of each
word capitalized. Make sure the rest of the word is in lower case.

For the purpose of this exercise, you should also
capitalize connecting words like "the" and "of".

### Before

```javascript
function titleCase(str) {
  return str;
}

titleCase("I'm a little tea pot");
```

### Answers

```javascript
function titleCase(str) {

  var arr = str.split(' ');

  var newStr = '';

  for (var i = 0; i < arr.length; i++) {
    var lower = arr[i].toLowerCase();
    newStr += lower.charAt(0).toUpperCase() + lower.slice(1) + ' ';
  }

  return newStr.trim();
}

titleCase("HERE IS MY HANDLE HERE IS MY SPOUT");
```

### Thinking

1. We transform the str into a arr;
2. We setup a empty var;
3. We loop inside the arr
4. We put all word to lower case
5. We change only the first letter to upper case
6. We need to `trim()` the last whitespace