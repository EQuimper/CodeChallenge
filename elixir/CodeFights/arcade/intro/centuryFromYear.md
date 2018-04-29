Given a year, return the century it is in. The first century spans from the year 1 up to and including the year 100, the second - from the year 101 up to and including the year 200, etc.

```elixir
def centuryFromYear(year) do
  calculCentury(year, 0)
end

def calculCentury(year, centuryCount) when year <= 0, do: centuryCount

def calculCentury(year, centuryCount) do
  calculCentury(year - 100, centuryCount + 1)
end
```
