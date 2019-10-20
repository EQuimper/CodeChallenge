# Apply Functional Programming to Convert Strings to URL Slugs

Fill in the urlSlug function so it converts a string title and returns the hyphenated version for the URL. You can use any of the methods covered in this section, and don't use replace. Here are the requirements:

The input is a string with spaces and title-cased words

The output is a string with the spaces between words replaced by a hyphen (-)

The output should be all lower-cased letters

The output should not have any spaces

### Before

```js
// the global variable
var globalTitle = "Winter Is Coming";

// Add your code below this line
function urlSlug(title) {


}
// Add your code above this line

var winterComing = urlSlug(globalTitle); // Should be "winter-is-coming"
```

### Answer

```js
// the global variable
var globalTitle = "Winter Is Coming";

// Add your code below this line
function urlSlug(title) {
    const parseTitle = title.toLowerCase().trim();

    return parseTitle.split(/\W/).filter(el => el.length > 0).join('-');
}
// Add your code above this line

var winterComing = urlSlug(globalTitle); // Should be "winter-is-coming"
```

### Thinking

I'm so bad in regex so I filter those empty space :)
