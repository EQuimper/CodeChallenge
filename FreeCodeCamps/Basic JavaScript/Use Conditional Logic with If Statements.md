# Use Conditional Logic with If Statements

`If` statements are used to make decisions in code. The keyword `if`
tells JavaScript to execute the code in the curly braces
under certain conditions, defined in the parentheses. These conditions
are known as `Boolean` conditions because they may *only be true or false*.

When the condition evaluates to `true`, the program *executes
the statement inside the curly braces*. When the `Boolean` condition
evaluates to `false`, the *statement inside the curly braces will not execute*.

## PseudoCode

```javascript
if (condition is true) {
  statement is executed
}
```

## Example

```javascript
function test (myCondition) {
  if (myCondition) {
     return "It was true";
  }
  return "It was false";
}
test(true);  // returns "It was true"
test(false); // returns "It was false"
```

When test is called with a value of `true`, the `if` statement evaluates
`myCondition` to see if it is `true` or not. Since it is `true`,
the function returns `"It was true"`. When we call test with a value
of `false`, `myCondition` is not `true` and the statement in the curly braces
is not executed and the function returns `"It was false"`.

## Instructions
- Create an `if` statement inside the function to return `"Yes, that was
true"` if the parameter `wasThatTrue` is `true` and return `"No, that was
false"` `otherwise`.

### Before

```javascript
// Example
function ourFunction(isItTrue) {
  if (isItTrue) {
    return "Yes, it's true";
  }
  return "No, it's false";
}

// Setup
function trueOrFalse(wasThatTrue) {

  // Only change code below this line.



  // Only change code above this line.

}

// Change this value to test
trueOrFalse(true);
```

### Answers

```javascript
// Example
function ourFunction(isItTrue) {
  if (isItTrue) {
    return "Yes, it's true";
  }
  return "No, it's false";
}

// Setup
function trueOrFalse(wasThatTrue) {

  // Only change code below this line.

  if(wasThatTrue) return "Yes, that was true";

  return "No, that was false";

  // Only change code above this line.

}

// Change this value to test
trueOrFalse(true);
```

### Thinking

Here I use some shortcut.