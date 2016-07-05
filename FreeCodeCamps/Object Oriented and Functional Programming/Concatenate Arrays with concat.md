# Concatenate Arrays with concat

`concat` can be used to merge the contents of two `arrays` into one.

`concat` takes an `array` as an argument and returns a `new array` with
the elements of this `array` concatenated onto the **end**.

Here is an example of concat being used to concatenate
`otherArray` onto the end of `oldArray`:

`newArray = oldArray.concat(otherArray);`

Use `.concat()` to concatenate `concatMe` onto the end of `oldArray` and
assign it to newArray.

### Before

```javascript
var oldArray = [1,2,3];
var newArray = [];

var concatMe = [4,5,6];

// Only change code below this line.

newArray = oldArray;
```

### Answers

```javascript
var oldArray = [1,2,3];
var newArray = [];

var concatMe = [4,5,6];

// Only change code below this line.

newArray = oldArray.concat(concatMe);
```