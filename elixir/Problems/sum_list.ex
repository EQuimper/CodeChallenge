defmodule SumList do
  # add initial value do the sum function
  def sum(list, acc \\ 0)
  # if no list anymore return the accumulator
  def sum([], acc), do: acc
  # cut the head and tail for let recurse over the with the tail as the list
  # and change the accumulator to be the addition of the past value and the head
  def sum([head | tail], acc), do: sum(tail, acc + head)
end
