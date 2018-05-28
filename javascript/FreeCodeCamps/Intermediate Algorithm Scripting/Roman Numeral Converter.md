# Roman Numeral Converter

Convert the given number into a roman numeral.

All roman numerals answers should be provided in upper-case.

### Before

```javascript
function convertToRoman(num) {
 return num;
}

convertToRoman(36);
```

### Answers

```javascript
function convertToRoman(num) {
 var romans = ["I", "V", "X", "L", "C", "D", "M"],
     ints = [],
     romanNumber = [],
     numeral = "";

  while (num) {
    ints.push(num % 10);
    num = Math.floor(num/10);
  }

  for (var i = 0; i < ints.length; i++){
    units(ints[i]);
  }

  function units() {
    numeral = romans[i*2];
    switch(ints[i]) {
      case 1:
        romanNumber.push(numeral);
        break;
      case 2:
        romanNumber.push(numeral.concat(numeral));
        break;
      case 3:
        romanNumber.push(numeral.concat(numeral).concat(numeral));
        break;
      case 4:
        romanNumber.push(numeral.concat(romans[(i*2)+1]));
        break;
      case 5:
        romanNumber.push(romans[(i*2)+1]);
        break;
      case 6:
        romanNumber.push(romans[(i*2)+1].concat(numeral));
        break;
      case 7:
        romanNumber.push(romans[(i*2)+1].concat(numeral).concat(numeral));
        break;
      case 8:
        romanNumber.push(romans[(i*2)+1].concat(numeral).concat(numeral).concat(numeral));
        break;
      case 9:
        romanNumber.push(romans[i*2].concat(romans[(i*2)+2]));
      }
    }
  return romanNumber.reverse().join("").toString();
}

convertToRoman(36);
```

New thinking using recursive

```js
var normalFormat = [1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000];

var romanFormat = ['I', 'IV', 'V', 'IX', 'X', 'XL', 'L', 'XC', 'C', 'CD', 'D', 'CM', 'M'];

function convertToRoman(num) {
  if (num === 0) return '';

  var value = '';

  for (var i = 0; i < normalFormat.length; i++) {
    if (num >= normalFormat[i]){
      value = '';

      value = romanFormat[i];
    }
  }

  var index = romanFormat.indexOf(value);

  var subValue = normalFormat[index];

  return value + convertToRoman(num - subValue);
}
```
