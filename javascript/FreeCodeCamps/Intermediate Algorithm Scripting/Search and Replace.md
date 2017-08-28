# Search and Replace

Perform a search and replace on the sentence using the arguments provided and return the new sentence.

First argument is the sentence to perform the search and replace on.

Second argument is the word that you will be replacing (before).

Third argument is what you will be replacing the second argument with (after).

NOTE: Preserve the case of the original word when you are replacing it. For example if you mean to replace the word "Book" with the word "dog", it should be replaced as "Dog"

## Solutions

```js
function myReplace(str, before, after) {
  // create a variable who gonna contains the after var if need to be change
  var nAfter;
  // create a function who return a string with the first letter uppercase
  function makeFirstUpperCase(s) {
    return s.charAt(0).toUpperCase() + s.slice(1);
  }
  // check if the first letter of "before" is uppercase
  if (before.charAt(0) === before.charAt(0).toUpperCase()) {
    // if it's uppercase make the after variable first letter uppercase with the function we create.
    nAfter = makeFirstUpperCase(after);
  } else {
    // if not just return the variable
    nAfter = after;
  }
  // split the str with the word we want to replace and join that with the after variable
  return str.split(before).join(nAfter);
}

myReplace("A quick brown fox jumped over the lazy dog", "jumped", "leaped");
```
