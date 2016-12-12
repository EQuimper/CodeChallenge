Find the missing letter in the passed letter range and return it.

If all letters are present in the range, return undefined.

Answer

```js

function fearNotLetter(str) {
  // Declare variable
  var arrToCheck = [];
  var notHere;

  // Loop over the str and push the charCode number
  for (var i = 0; i < str.length; i++) {
    arrToCheck.push(str.charCodeAt(i));
  }

  // loop over the arrToCheck and check which number is missing
  for (var x = 1; x < arrToCheck.length; x++) {
    if (arrToCheck[x] - arrToCheck[x - 1] !== 1) {
      // arrToCheck is the number + 1 so that why we -1
      notHere = arrToCheck[x] - 1;
    }
  }

  // Return undefined if no miss
  if (!notHere) {
    notHere = undefined;
  } else {
    // Return the string from the char code
    notHere = String.fromCharCode(notHere);
  }

  return notHere;
}

fearNotLetter("abce");
```
