# Reverse a String

Reverse the provided string.

You may need to turn the string into an array before you can reverse it.

Your result must be a string.

### Before

```javascript
function reverseString(str) {
  return str;
}

reverseString("hello");
```

### Answers

```javascript
function reverseString(str) {
  return str.split('').reverse().join('');
}

reverseString("hello");
```

### Thinking

1. We take the str and we make it a array
2. We reverse it
3. And finally we join it for be a string back