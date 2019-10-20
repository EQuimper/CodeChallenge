# Split a String into an Array Using the split Method

Use the split method inside the splitify function to split str into an array of words. The function should return the array. Note that the words are not always separated by spaces, and the array should not contain punctuation.

### Before

```js
function splitify(str) {
  // Add your code below this line


  // Add your code above this line
}
splitify("Hello World,I-am code");
```

### Answer

```js
function splitify(str) {
  // Add your code below this line

  return str.split(/\W/);
  // Add your code above this line
}
splitify("Hello World,I-am code");
```

### Thinking

1. We want to split on non letters
