defmodule Solution do
  @input_path "input.txt"
  def read_input([]) do
    File.stream!(@input_path)
  end

  def read_input(args) do
    case args do
      [path] -> File.stream!(path)
      _ -> raise ArgumentError, "usage: p2.exs *input_path*"
    end
  end

  def find_min_cubes(set, acc) do
    Enum.reduce(set, acc, fn {key, val}, acc ->
      case Map.get(acc, key) do
        old when old != nil ->
          if val > old do
            %{acc | key => val}
          else
            acc
          end

        nil ->
          Map.put(acc, key, val)
      end
    end)
  end

  # It actually does a lot more than parsing so a better name would be nice
  def parse_game(game) do
    [_ | sets] = String.trim(game) |> String.split(":")

    String.split(hd(sets), ";")
    |> Enum.map(fn set ->
      Regex.scan(~r/(\d*)\s(red|green|blue)/, set, capture: :all_but_first)
      |> Enum.reduce(%{}, fn [num | color], acc ->
        Map.put(acc, hd(color), String.to_integer(num))
      end)
    end)
    |> Enum.reduce(%{}, fn set, acc -> find_min_cubes(set, acc) end)
    |> Enum.reduce(1, fn {_key, val}, acc -> acc * val end)
  end

  def solve() do
    read_input(System.argv()) |> Enum.map(&parse_game/1) |> Enum.reduce(&+/2) |> IO.inspect()
  end
end

Solution.solve()
