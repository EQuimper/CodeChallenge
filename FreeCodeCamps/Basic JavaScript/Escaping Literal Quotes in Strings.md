# Escaping Literal Quotes in Strings

When you are defining a string you must start and end with a single
or double quote. What happens when you need a literal quote: `"` or `'` inside
of your string?

In JavaScript, you can escape a quote from considering it as an end
of string quote by placing a backslash (\) in front of the quote.

`var sampleStr = "Alan said, \"Peter is learning JavaScript\".";`

This signals to JavaScript that the following quote is not the end of the
string, but should instead appear inside the string. So if you were to print
this to the console, you would get:

`Alan said, "Peter is learning JavaScript".`

## Instructions
 - Use backslashes to assign a string to the myStr variable so that if you were to print it to the console, you would see:

`I am a "double quoted" string inside "double quotes".`

### Before

```javascript
var myStr; // Change this line
```

### Answers

```javascript
var myStr = "I am a \"double quoted\" string inside \"double quotes\".";
```