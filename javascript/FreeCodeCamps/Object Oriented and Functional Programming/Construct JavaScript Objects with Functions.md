# Construct JavaScript Objects with Functions

We are also able to create objects using `constructor` functions.

A `constructor` function is given a capitalized name
to make it clear that it is a constructor.

Here's an example of a constructor function:

```javascript
var Car = function() {
  this.wheels = 4;
  this.engines = 1;
  this.seats = 1;
};
```

In a constructor the this variable refers to the new
object being created by the constructor. So when we write,

  `this.wheels = 4;`

inside of the `constructor` we are giving the new object it
creates a property called wheels with a value of 4.

You can think of a `constructor` as a description for the object
it will create.

Have your MotorBike `constructor` describe an object with wheels,
engines and seats properties and set them to numbers.

### Before

```javascript
var Car = function() {
  this.wheels = 4;
  this.engines = 1;
  this.seats = 1;
};

// Only change code below this line.

var MotorBike = function() {

};
```

### Answers

```javascript
var Car = function() {
  this.wheels = 4;
  this.engines = 1;
  this.seats = 1;
};

// Only change code below this line.

var MotorBike = function() {
  this.wheels = 2;
  this.engines = 1;
  this.seats = 1;
};
```