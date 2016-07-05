# Make Instances of Objects with a Constructor Function

Now let's put that great `constructor` function we made in
the last lesson to use!

To use a `constructor` function we call it with the new
keyword in front of it like:

`var myCar = new Car();`

`myCar` is now an instance of the `Car` constructor that looks like
the object it described:

```
{
  wheels: 4,
  engines: 1,
  seats: 1
}
```
Note that it is important to use the new keyword
when calling a constructor. This is how Javascript knows
to create a new object and that all the references to this
inside the constructor should be referring to this new object.

Now, once the `myCar` instance is created it can be used like
any other object and can have its properties accessed and
modified the same way you would usually. For example:

`myCar.turboType = "twin";`

Our myCar variable now has a property `turboType` with a value of "twin".

In the editor, use the Car constructor to create a new instance and
assign it to `myCar`.

Then give myCar a nickname property with a string value.

### Before

```javascript
var Car = function() {
  this.wheels = 4;
  this.engines = 1;
  this.seats = 1;
};

// Only change code below this line.

var myCar;
```

### Answers

```javascript
var Car = function() {
  this.wheels = 4;
  this.engines = 1;
  this.seats = 1;
};

// Only change code below this line.

var myCar = new Car();

myCar.nickname = "Lambo";
```