In the world of birding there are four-letter codes for the common names of birds. These codes are created by some simple rules:

If the bird's name has only one word, the code takes the first four letters of that word.
If the name is made up of two words, the code takes the first two letters of each word.
If the name is made up of three words, the code is created by taking the first letter from the first two words and the first two letters from the third word.
If the name is four words long, the code uses the first letter from all the words.
There are other ways that codes are created. But, this Kata will only use the four rules listed above. The input arrays will have a list of common birds from North America.

In this Kata you will write a function that takes an array of strings of common bird names and create the codes for those names based on the rules above. The function will return an array of the codes in the same order in which the input names were presented.

Additional considertations:
The four-letter codes in the returned array should be in UPPER CASE.
If a common name has a hyphen/dash, it should be considered a space.
Example:
If the input array is: ["Black-Capped Chickadee", "Common Tern"]

The return array would be: ["BCCH", "COTE"]

``` js
function birdCode(arr){
  return arr.map(el => buildCode(el))
}

function buildCode(str) {
  const arr = str.split(/[ -]+/);
  let newStr = '';

  switch (arr.length) {
    case 1:
      newStr = arr[0].slice(0, 4);
      break;
    case 2: {
      newStr = arr[0].slice(0, 2);
      newStr += arr[1].slice(0, 2);
      break;
    }
    case 3: {
     newStr = arr[0].slice(0, 1);
      newStr += arr[1].slice(0, 1);

      newStr += arr[2].slice(0, 2);
      break;
    }
    case 4:
      newStr = arr.map(el => el.slice(0, 1)).join('')
      break;
  }

  return newStr.toUpperCase();
}
```

1. I create a function who take a str and return the new str
2. we split this string in an array if space or dash
3. create a newStr variables who gonna get the new string in
4. we do the logic for each statement
5. make sure to return this string in uppercase
6. We map over each string with the function we create. With map this is gonna return an new array
