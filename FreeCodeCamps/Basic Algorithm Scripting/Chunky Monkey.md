# Chunky Monkey

Write a function that splits an array (first argument) into groups
the length of size (second argument) and returns them as a
two-dimensional array.

### Before

```javascript
function chunkArrayInGroups(arr, size) {
  // Break it up.
  return arr;
}

chunkArrayInGroups(["a", "b", "c", "d"], 2);
```

### Answers

```javascript
function chunkArrayInGroups(arr, size) {
  var masterArr = [];

  var item = 0;

  while(item < arr.length) {
    masterArr.push(arr.slice(item, item += size));
  }

  return masterArr;

}

chunkArrayInGroups(["a", "b", "c", "d"], 2);
```

### Thinking

1. Setup a empty arr and and var 0
2. I use a while loop who gonna stop when the item size gonna be bigger
than the arr size
3. I push the slice to the arr and I use the item for be the start index inside
my slice and I sum it with the size.

Because of the for loop the item size growth with the size args so I can do the
loop until nothing left