defmodule CenturyFromYear do
  def centuryFromYear(year) do
    calculCentury(year, 0)
  end

  def calculCentury(year, centuryCount) when year <= 0, do: centuryCount

  def calculCentury(year, centuryCount) do
    calculCentury(year - 100, centuryCount + 1)
  end
end
