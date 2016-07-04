# Escape Sequences in Strings

Quotes are not the only characters that can be escaped inside a string.
Here is a table of common `escape` sequences:

| Code | Output|
-------|--------
\'  | single quote
\"	| double quote
\\	| backslash
\n	| new line
\r	| carriage return
\t	| tab
\b	| backspace
\f	| form feed

*Note that the backslash itself must be escaped in order
to display as a backslash.*

## Instructions
 - Assign the following two lines of text into the single variable myStr using escape sequences.

`Here is a backslash: \.
        Here is a new line with two tabs.`

You will need to use escape sequences to insert special
characters correctly. You will also need to follow the spacing
as it looks above with no additional spaces between each escape sequence.

Here is the text with the escape sequences written out.

`Here is a backslash: backslash.newline tab tab Here
is a new line with two tabs.`

### Before

```javascript
var myStr; // Change this line
```

### Answers

```javascript
var myStr = 'Here is a backslash: \\.\n\t\tHere is a new line with two tabs.';
```