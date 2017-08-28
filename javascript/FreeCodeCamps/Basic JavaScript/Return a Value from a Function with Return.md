# Return a Value from a Function with Return

We can pass values into a function with arguments. You can use a
`return` statement to send a value back out of a `function`.

## Example

```javascript
function plusThree(num) {
  return num + 3;
}
var answer = plusThree(5); // 8
```

`plusThree` takes an argument for `num` and returns a value equal to `num + 3`.

## Instructions
 - Create a function `timesFive` that accepts one argument, multiplies it by `5`,
 and returns the new value.

### Before

```javascript
// Example
function minusSeven(num) {
  return num - 7;
}

// Only change code below this line
```

### Answers

```javascript
// Example
function minusSeven(num) {
  return num - 7;
}

// Only change code below this line
function timesFive(num) {
  return num * 5;
}

var answer = timesFive(5);
```