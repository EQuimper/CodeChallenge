# Factorial

## Example

```javascript
 For n = 5, // the output should be
 factorial(n) = 120.

 /*Here*/ 1 * 2 * 3 * 4 * 5 = 120.
```

### Answers

```javascript
function factorial(n) {
	// If number is less than 0 reject.
	if (n < 0) {
		return -1;
	} else if (n == 0) { // If number 0, return factorial 1.
		return 1;
	} else {
		return (n * factorial(n-1));
	}
}

factorial(5);

```