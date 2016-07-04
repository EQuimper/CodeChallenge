# Find its minimal element

- Given an array of integers, find its minimal element.

## Example

```javascript
 For inputArray = [19, 32, 11, 23], the output should be
 arrayMinimum(inputArray) = 11.
```

```javascript
function arrayMinimum(inputArray) { // inputArray = [19, 32, 11, 23]
	var indexOfMinimum = 0; // Create the var indexOfMinimum.

	for (var i = 1; i < inputArray.length; i++) { // i is my number inside the array.
		if (inputArray[i] < inputArray[indexOfMinimum]) {
			indexOfMinimum = i; // That give me the index of the smaller one.
		}
	}

	return  inputArray[indexOfMinimum];
}

arrayMinimum([19, 32, 11, 23]); // Result: 11

```
### Thinking

Here I need to find which number is the smaller inside that array.

1. I create the var you gonna accept the index of the smaller.
2. I do a for loop for this given array.
3. If the number is smaller than the last one return the index of this one
4. Finally return the array with the index for give the number appropriate to this index
