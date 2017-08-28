# Repeat a string repeat a string

Repeat a given string (first argument) num times (second argument).
Return an empty string if num is not a positive number.

### Before

```javascript
function repeatStringNumTimes(str, num) {
  // repeat after me
  return str;
}

repeatStringNumTimes("abc", 3);
```

### Answers

```javascript
function repeatStringNumTimes(str, num) {

  var newStr = '';

  if (num < 0) return newStr;

  newStr = str.repeat(num);

  return newStr;

}

repeatStringNumTimes("abc", 3);
```

### Thinking

1. If the number is negative we return the var newStr who is a empty string
2. I use the `.repeat()` method for multiply the string with the given number

### Links

[.repeat()](https://developer.mozilla.org/en/docs/Web/JavaScript/Reference/Global_Objects/String/repeat)