Given the string, check if it is a palindrome.

```elixir
def checkPalindrome(inputString) do
  inputString == String.reverse(inputString)
end
```
