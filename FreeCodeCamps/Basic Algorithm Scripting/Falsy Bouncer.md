# Falsy Bouncer

Remove all `falsy` values from an array.

`Falsy` values in JavaScript are `false`, `null`, `0`, `""`, `undefined`, and `NaN`.

### Before

```javascript
function bouncer(arr) {
  // Don't show a false ID to this bouncer.
  return arr;
}

bouncer([7, "ate", "", false, 9]);
```

### Answer

```javascript
function bouncer(arr) {

  var check = arr.filter(function(i) {
    return Boolean(i);
  });

  return check;
}

bouncer([7, "ate", "", false, 9]);
```

### Thinking

1. Make a var you gonna get the filter item who is not falsy
2. Return this var