# Join Strings with join

We can use the `join` method to `join` each element of
an `array` into a `string` separated by whatever
delimiter you provide as an argument.

The following is an example of using `join` to `join` all of the
elements of an array into a `string` with all the elements
separated by word and:

```javascript
var veggies = ["Celery", "Radish", "Carrot", "Potato"];
var salad = veggies.join(" and ");
console.log(salad); // "Celery and Radish and Carrot and Potato"
```

Use the `join` method to create a `string` from `joinMe` with `spaces`
in between each element and assign it to `joinedString`.

### Before

```javascript
var joinMe = ["Split","me","into","an","array"];
var joinedString = '';

// Only change code below this line.

joinedString = joinMe;
```

### Answers

```javascript
var joinMe = ["Split","me","into","an","array"];
var joinedString = '';

// Only change code below this line.

joinedString = joinMe.join(" ");
```