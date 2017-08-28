# Global Scope and Functions

In JavaScript, `scope` refers to the visibility of variables.
Variables which are defined outside of a function block have
Global scope. **This means, they can be seen everywhere in your
JavaScript code.**

Variables which are **used without the var keyword are automatically
created in the global scope**. This **can create unintended consequences
elsewhere** in your code or when running a function again. **You should
always declare your variables with var.**

## Instructions
 - Declare a global variable `myGlobal` outside of any function.
 Initialize it to have a value of `10`

 - Inside function `fun1`, assign `5` to `oopsGlobal` without using the var keyword.

### Before

```javascript
// Declare your variable here


function fun1() {
  // Assign 5 to oopsGlobal Here

}

// Only change code above this line
function fun2() {
  var output = "";
  if (typeof myGlobal != "undefined") {
    output += "myGlobal: " + myGlobal;
  }
  if (typeof oopsGlobal != "undefined") {
    output += " oopsGlobal: " + oopsGlobal;
  }
  console.log(output);
}
```

### Answers

```javascript
// Declare your variable here
var myGlobal = 10;

function fun1() {
  // Assign 5 to oopsGlobal Here
  oopsGlobal = 5;
}

// Only change code above this line
function fun2() {
  var output = "";
  if (typeof myGlobal != "undefined") {
    output += "myGlobal: " + myGlobal;
  }
  if (typeof oopsGlobal != "undefined") {
    output += " oopsGlobal: " + oopsGlobal;
  }
  console.log(output);

```