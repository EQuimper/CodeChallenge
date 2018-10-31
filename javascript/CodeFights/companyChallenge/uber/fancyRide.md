Being a new Uber user, you have $20 off your first ride. You want to make the most out of it and drive in the fanciest car you can afford, without spending any out-of-pocket money. There are 5 options, from the least to the most expensive: "UberX", "UberXL", "UberPlus", "UberBlack" and "UberSUV".

You know the length l of your ride in miles and how much one mile costs for each car. Find the best car you can afford.

Example

For l = 30 and fares = [0.3, 0.5, 0.7, 1, 1.3], the output should be
fancyRide(l, fares) = "UberXL".

The cost for the ride in this car would be $15, which you can afford, but "UberPlus" would cost $21, which is too much for you.

## Result

```js
function fancyRide(l, fares) {
    const currentAmount = 20;

    const vehicles = ["UberX", "UberXL", "UberPlus", "UberBlack", "UberSUV"];

    let vehicle = null;

    const faresWithAmount = fares.map(el => el * l);

    for (let i = fares.length - 1; i >= 0; i--) {
        if (faresWithAmount[i] <= currentAmount) {
            vehicle = vehicles[i]
            break;
        }
    }

    return vehicle
}
```

First we need to put all vehicles to an array and the current amount to 20.

We initialize a vehicle who is null by default. This will be the result we return

We create an array of the fares * by the miles. This will know the full cost of each vehicles.

After we do a reverse loop because the faresWithAmount will be an array of [lower, higher] value. By doing this we can then check if the amount is smaller or equal to the current amount. IF this is yes we give the vehicle value to this. Dont forget to break cause that will always give the smaller one if you dont
