Convert the characters &, <, >, " (double quote), and ' (apostrophe), in a string to their corresponding HTML entities.

```js
function convertHTML(str) {
  // make an array wth the string
  // much more easy to work with an
  // array than a string
  var arr = str.split('');
  // loop over the array
  for (var i = 0; i < arr.length; i++) {
    // use a switch statement for check every index
    switch (arr[i]) {
      case '&':
        arr[i] = '&amp;';
        break;
      case '<':
        arr[i] = '&lt;';
        break;
      case '>':
        arr[i] = '&gt;';
        break;
      case '"':
        arr[i] = '&quot;';
        break;
      case "'":
        arr[i] = '&apos;';
        break;
    }
  }
  // return as a string back
  return arr.join('');
}

convertHTML("Dolce & Gabbana");
```
