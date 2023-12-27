defmodule Solution do
  @input_path "../input.txt"
  def read_input([]) do
    read_input([@input_path])
  end

  def read_input(args) do
    case args do
      [path] -> File.read!(path) |> String.split("\n")
      _ -> raise ArgumentError, "usage: p2.exs *input_path*"
    end
  end

  def parse_node(network, line) do
    [id | [targets | _]] = String.split(line, " = ")
    [left | [right | _]] = Regex.scan(~r/\w+/, targets) |> Enum.map(&hd/1)

    Map.put(network, id, %{?L => left, ?R => right})
  end

  @start "A"
  @terminus "Z"
  def traverse_network(_network, @terminus, _path, _initial, len), do: len

  def traverse_network(network, id, [], initial, len) do
    traverse_network(network, id, initial, initial, len)
  end

  def traverse_network(network, id, [cur_dir | remaining], initial, len) do
    if String.ends_with?(id, @terminus) do
      len
    else
      traverse_network(network, network[id][cur_dir], remaining, initial, len + 1)
    end
  end

  @doc """
    Instead of trying to calculate the route simultaneously,
    we find the result for each starting node separately,
    and then calculate the total count by multiplying
    and dividing them by their greatest common divisor
  """
  def solve() do
    [path | nodes] =
      read_input(System.argv()) |> Enum.filter(fn line -> line != "" end)

    path = String.to_charlist(path)
    network = Enum.reduce(nodes, %{}, fn node, acc -> parse_node(acc, node) end)

    Enum.filter(network, &String.ends_with?(elem(&1, 0), @start))
    |> Enum.reduce([], fn {key, _val}, acc -> [key | acc] end)
    |> Enum.map(&traverse_network(network, &1, [], path, 0))
    |> Enum.reduce(fn x, acc -> div(x * acc, Integer.gcd(x, acc)) end)
    |> IO.puts()
  end
end

Solution.solve()
