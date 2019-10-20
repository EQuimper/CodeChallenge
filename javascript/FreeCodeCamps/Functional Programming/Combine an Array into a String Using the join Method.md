# Combine an Array into a String Using the join Method

Use the join method (among others) inside the sentensify function to make a sentence from the words in the string str. The function should return a string. For example, "I-like-Star-Wars" would be converted to "I like Star Wars". For this challenge, do not use the replace method.

### Before

```js
function sentensify(str) {
  // Add your code below this line


  // Add your code above this line
}
sentensify("May-the-force-be-with-you");
```

### Answer

```js
function sentensify(str) {
  // Add your code below this line

  return str.split(/\W/).join(' ');
  // Add your code above this line
}
sentensify("May-the-force-be-with-you");
```

### Thinking

1. We need to split first on non-letters and join right after with a space.
