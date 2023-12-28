defmodule Solution do
  @input_path "../input.txt"
  def read_input([]) do
    read_input([@input_path])
  end

  def read_input(args) do
    case args do
      [path] -> File.read!(path) |> String.trim_trailing() |> String.split("\n")
      _ -> raise ArgumentError, "usage: p1.exs *input_path*"
    end
  end

  def calc_diff(layers) do
    latest = hd(layers)

    if Enum.count(latest, &(&1 == 0)) == Enum.count(latest) do
      layers
    else
      diff =
        Enum.chunk_every(latest, 2, 1, :discard)
        # because the list is reversed
        |> Enum.map(fn [next | [prev | _]] -> next - prev end)

      calc_diff([diff | layers])
    end
  end

  def calc_edge([], prev), do: prev

  def calc_edge([cur | rest], prev) do
    calc_edge(rest, prev + hd(cur))
  end

  @doc """
    Algorithm:
      1. Parse input line by line with every line resulting in a numbers list
      2. For every list:
      3. Create a list of lists and put the initial list there
      4. Calculate the difference between every element and save it into a new sublist
      5. Repeat until every element of the resulting sequence is zero
      6. Calculate the new element of the base list by recursively summing its adjacent elements
  """
  def solve() do
    read_input(System.argv())
    |> Enum.map(fn line ->
      # reverse initially because we'll use the last elem of the linked list later
      # and getting its head is much faster
      String.split(line, " ") |> Enum.reverse() |> Enum.map(&String.to_integer/1)
    end)
    |> Enum.map(fn history -> calc_diff([history]) end)
    |> Enum.map(fn layers -> calc_edge(layers, 0) end)
    |> Enum.reduce(0, &+/2)
    |> IO.inspect()
  end
end

Solution.solve()
