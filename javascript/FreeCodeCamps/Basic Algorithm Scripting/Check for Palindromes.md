# Check for Palindromes

Return true if the given string is a palindrome. Otherwise, return false.

A palindrome is a word or sentence that's spelled the same
way both forward and backward, ignoring punctuation, case, and spacing.

## Note
 *You'll need to remove all non-alphanumeric characters (punctuation, spaces and symbols) and turn everything lower case in order to check for palindromes.*

We'll pass strings with varying formats, such as "racecar", "RaceCar", and "race CAR" among others.

We'll also pass strings with special symbols, such as "2A3*3a2", "2A3 3a2", and "2_A3*3#A2".

### Before

```javascript
function palindrome(str) {
  // Good luck!
  return true;
}

palindrome("eye");
```

### Answers

```javascript
function palindrome(str) {
  // Good luck!

  var newStr = str.replace(/[\W_]/g, '').toLowerCase();

  var strReverse = newStr.split('').reverse().join('');

  return newStr === strReverse;
}

palindrome("race car");
```

### Thinking

1. Create a new var who is the str cleaner and lowercase
2. Create a new var who take the newStr and transform it into a array for
reverse it and transform back into string
3. Check if the two var are equal