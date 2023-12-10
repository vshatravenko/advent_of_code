defmodule Solution do
  @limits %{"red" => 12, "green" => 13, "blue" => 14}
  @input_path "input.txt"
  def read_input([]) do
    File.stream!(@input_path)
  end

  def read_input(args) do
    case args do
      [path] -> File.stream!(path)
      _ -> raise ArgumentError, "usage: p1.exs *input_path*"
    end
  end

  def check_limits(set) do
    Enum.filter(set, fn {key, _val} -> set[key] > @limits[key] end) |> length() == 0
  end

  def parse_game(game) do
    [raw_id | sets] = String.trim(game) |> String.split(":")

    is_valid =
      String.split(hd(sets), ";")
      |> Enum.map(fn set ->
        Regex.scan(~r/(\d*)\s(red|green|blue)/, set, capture: :all_but_first)
        |> Enum.reduce(%{}, fn [num | color], acc ->
          Map.put(acc, hd(color), String.to_integer(num))
        end)
      end)
      |> Enum.reduce(true, fn set, acc -> acc && check_limits(set) end)

    if is_valid do
      [match | _rest] = Regex.run(~r/(\d*)$/, raw_id, capture: :all_but_first)
      String.to_integer(match)
    else
      0
    end
  end

  def solve() do
    read_input(System.argv()) |> Enum.map(&parse_game/1) |> Enum.reduce(&+/2) |> IO.inspect()
  end
end

Solution.solve()
