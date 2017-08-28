# Caesars Cipher

One of the simplest and most widely known ciphers is a
Caesar cipher, also known as a shift cipher. In a shift cipher
the meanings of the letters are shifted by some set amount.

A common modern use is the ROT13 cipher, where the values of the
letters are shifted by 13 places. Thus 'A' ↔ 'N', 'B' ↔ 'O' and so on.

Write a function which takes a ROT13 encoded string as input
and returns a decoded string.

All letters will be uppercase. Do not transform any non-alphabetic
character (i.e. spaces, punctuation), but do pass them on.

### Before

```javascript
function rot13(str) { // LBH QVQ VG!

  return str;
}

// Change the inputs below to test
rot13("SERR PBQR PNZC");
```

### Answers

```javascript
function rot13(str) { // LBH QVQ VG!
  var arr = [];

  for (var i = 0; i < str.length; i++) {
    if (str.charCodeAt(i) < 65 || str.charCodeAt(i) > 90) {
      arr.push(str.charAt(i));
      console.log(arr);
    } else if (str.charCodeAt(i) > 77) {
      arr.push(String.fromCharCode(str.charCodeAt(i) - 13));
    } else {
      arr.push(String.fromCharCode(str.charCodeAt(i) + 13));
    }
  }

  return arr.join("");
}

// Change the inputs below to test
rot13("SERR CVMMN!");
```

### Thinking

1. Setup a global empty arr
2. We loop the args
3. We check for getting only the alphanumeric
4. We know the code need to be inside 65 - 90
5. We transform it to be a String from the char code
6. Because it's a array of letter we join it