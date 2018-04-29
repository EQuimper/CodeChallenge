Make a program that filters a list of strings and returns a list with only your friends name in it.

If a name has exactly 4 letters in it, you can be sure that it has to be a friend of yours! Otherwise, you can be sure he's not...

Ex: Input = ["Ryan", "Kieran", "Jason", "Yous"], Output = ["Ryan", "Yous"]

Note: keep the original order of the names in the output.

### Thinking

Pretty simple stuff to make here. We just need to filter each name and returning only those who have the length === 4.

### Solution

```js
function friend(friends){
  return friends.filter(el => el.length === 4);
}
```
