# Confirm the Ending

Check if a string (first argument, str) ends with the given
target string (second argument, target).

This challenge can be solved with the `.endsWith()` method, which was
introduced in ES2015. But for the purpose of this challenge,
we would like you to use one of the JavaScript substring methods instead.

### Before

```javascript
  // "Never give up and good luck will find you."
  // -- Falcor
  return str;
}

confirmEnding("Bastian", "n");
```

### Answers

```javascript
function confirmEnding(str, target) {
  var newStr = '';

  newStr = str.substring(str.length - target.length);

  return newStr === target;

}

confirmEnding("Bastian", "n");
```

### Thinking

1. We setup a empty var for get the answer
2. We give to this var the value of the str and we get the last index of the string
with substring.
3. If the newStr !== the target with return false