# Split Strings with split

You can use the `split` method to `split` a `string` into an `array`.

`split` uses the argument you pass in as a delimiter to
determine which points the `string` should be `split` at.

Here is an example of `split` being used to `split` a `string` at every
`s` character:

```javascript
var array = string.split('s');
```

Use split to create an `array` of words from `string` and assign it to `array`.

### Before

```javascript
var string = "Split me into an array";
var array = [];

// Only change code below this line.

array = string;
```

### Answers

```javascript
var string = "Split me into an array";
var array = [];

// Only change code below this line.

array = string.split(/\s/g);
```

### Thinking

Here I use the regex `\s` for find the space