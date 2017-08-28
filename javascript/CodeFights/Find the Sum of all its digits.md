# Find the Sum of all its digits

- Given an integer, find the sum of all its digits.

## Example
```javascript
 For n = 111, // the output should be
 digitSum(n) = 3.

 1 + 1 + 1 = 3.
 ```

### Answers

```javascript
function digitSum(n) {
	var myNum = n.toString().split(''); // Result [ '6', '6', '6' ]
	var sum = 0; // Create the var sum

	for (var i = 0; i < myNum.length; i++) { // i is my digits
		var x = Number(myNum[i]); // String to Number

		sum += x; // Addition the digits
	}
	return sum;
}

digitSum(666); // Result: 18
```

### Thinking

Here I need to addition every digits inside a number.

1. I put the number inside a string and after I can split it and he become an array.
2. I setup the var sum for use it later.
3. I do a for loop for my array I just create so now I can use the .length.
4. I need to convert the string to a number for make a addition
5. Finally I return the sum.

