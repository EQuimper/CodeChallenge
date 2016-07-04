# Storing Values With the assignment Operator

In JavaScript, you can store a value in a variable with the assignment operator.

`myVariable = 5;`

Assigns the Number value 5 to myVariable.

>Assignment always goes from right to left. 

Everything to the right of the = operator is resolved before the value is assigned to the variable to the left of the operator.

```javascript
myVar = 5;
myNum = myVar;

```

Assigns 5 to myVar and then resolves myVar to 5 again and assigns it to myNum.

## Instructions

- Assign the value 7 to variable a.
- Assign the contents of a to variable b.

### Answers

```javascript
// Setup
var a;
var b = 2;

// Only change code below this line
a = 7;
b = 7;
b = a; // Assign the contents of a to var b /> so we need to have b at left cause we assign the right

```
