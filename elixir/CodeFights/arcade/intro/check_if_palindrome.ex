defmodule CheckIfPalindrome do
  def checkPalindrome(inputString) do
    reverse_string(inputString) == inputString
  end

  def reverse_string(str) do
    str
    |> String.split(~r{})
    |> Enum.reverse
    |> Enum.join
  end
end
