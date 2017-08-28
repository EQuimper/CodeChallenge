# Compound Assignment With Augmented Addition 

In programming, it is common to use assignments to modify the contents of a variable.
Remember that everything to the right of the equals sign is evaluated first, so we can say:

```javascript
myVar = myVar + 5;
```

to add 5 to myVar. Since this is such a common pattern, there are operators which do both a mathematical operation and assignment in one step.

One such operator is the += operator.

`myVar += 5;` will add 5 to myVar.

## Instructions
 - Convert the assignments for a, b, and c to use the `+=` operator.

### Before 

```javascript
var a = 3;
var b = 17;
var c = 12;

// Only modify code below this line

a = a + 12;
b = 9 + b;
c = c + 7;
```

### Answers

```javascript
var a = 3;
var b = 17;
var c = 12;

// Only modify code below this line

a += 12;
b += 9;
c += 7;
```
