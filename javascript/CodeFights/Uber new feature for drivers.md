# Uber new feature for drivers

 Uber is introducing a new feature for drivers that tells them if they'll
 need gas within the next hour. It analyzes all of the drive distances
 (in miles) that the driver has traveled in the past
 12 hours to make the recommendation.

 - Given the current gas level (in gallons), the drive data, and the
 average mileage of the vehicle, calculate the average amount of gas
 spent per hour and return true if the driver is likely to need a refill
 in the next hour, false otherwise.

 > Assume that the driver will need more gas if the average gas consumption per hour is greater than the amount of gas they have at the given moment. The average mileage is measured in miles per gallon and it shows how many miles a vehicle can travel on one gallon of gas.

 ## Example

```javascript
 For driveDistances = [12, 6, 17, 5, 20], currentGasLevel = 0.25 and avgMileage = 25, the output should be
 gasPrediction(driveDistances, currentGasLevel, avgMileage) = false.

 Since for driveDistances = [12, 6, 17, 5, 20], currentGasLevel = 0.25g and avgMileage = 25mpg, the total distance traveled is equal to 60 and it can be shown that on average the driver spends 0.2g of gas per hour, which is less than 0.25g.
```

### Answers

```javascript
function gasPrediction(driveDistances, currentGasLevel, avgMileage) {

	var totalDistance = 0;
	for (var i = 0; i < driveDistances.length; i++) {
		totalDistance += driveDistances[i];
	}
	var gasConsumed = totalDistance / avgMileage;

	return currentGasLevel < gasConsumed / 12;
}
```

