Given a year, return the century it is in. The first century spans from the year 1 up to and including the year 100, the second - from the year 101 up to and including the year 200, etc.

```js
function centuryFromYear(year) {
  let centuryCount = 0;
  while (year > 0){
    year = year - 100;
    centuryCount = centuryCount + 1;
  }
  return centuryCount;
}
```
