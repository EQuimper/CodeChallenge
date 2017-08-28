# Truncate a string

Truncate a string (first argument) if it is longer than the
given maximum string length (second argument). Return the
truncated string with a `...` ending.

Note that inserting the three dots to the end will add to the string length.

However, if the given maximum string length num is
less than or equal to 3, then the addition of the three
dots does not add to the string length in determining the truncated string.

### Before

```javascript
function truncateString(str, num) {
  // Clear out that junk in your trunk
  return str;
}

truncateString("A-tisket a-tasket A green and yellow basket", 11);
```

### Answers

```javascript
function truncateString(str, num) {
  var newStr = '';

  if (num <= 3) {
    newStr = str.slice(0, num) + '...';
  } else if (str.length > num) {
    newStr = str.slice(0, num - 3) + '...';
  } else {
    return str;
  }

  return newStr;
}

truncateString("Absolutely Longer", 2);
```

### Thinking

1. If the num is equal or less than 3 we `slice()` and add the `...`
2. If the str is bigger the num we slice with but we need to remove 3
cause of the dot
3. Else we touch nothing