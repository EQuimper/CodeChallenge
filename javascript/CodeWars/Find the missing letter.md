Write a method that takes an array of consecutive (increasing) letters as input and that returns the missing letter in the array.

You will always get an valid array. And it will be always exactly one letter be missing. The length of the array will always be at least 2.
The array will always contain letters in only one case.

Example:

```
['a','b','c','d','f'] -> 'e'
['O','Q','R','S'] -> 'P'
```

## Thinking

1. We need to be able to get the position in the chartCode from each letter. So I create the function getLetterPosition who get a single letter and return the position.
2. Javascript don't have a range function like in elixir. So I create a little function numberRanger who take a start and a end and give you a array of the number between those numbers.
3. getMissingNumber sort the array, this way we make sure we have a real sequence. After we init a variable who gonna be the chartCode of the missing letter. We loop over the range, if the number don't exist inside the arraySort we know this is the one missing.
4. Finally I use String.fromCharCode for returning the letter missing.


## Solution

```js
function numberRange (start, end) {
  return new Array(end - start).fill().map((d, i) => i + start);
}

function getLetterPosition(letter) {
  return letter.charCodeAt(0);
}

function getMissingNumber(array) {
  const arraySort = array.sort((a, b) => a - b);

  let numberMissing;

  const range = numberRange(arraySort[0], arraySort[arraySort.length - 1])

  for (let i = 0; i < range.length; i++) {
    const num = range[i]
    if (arraySort.indexOf(num) < 0) {
      numberMissing = num;
      break;
    }
  }

  return numberMissing;
}

function findMissingLetter(array) {
  const letterPosition = array.map(el => getLetterPosition(el));

  const numberMissing = getMissingNumber(letterPosition)

  return String.fromCharCode(numberMissing)
}
```
