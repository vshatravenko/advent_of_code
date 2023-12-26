defmodule Solution do
  @input_path "../input.txt"
  def read_input([]) do
    read_input([@input_path])
  end

  def read_input(args) do
    case args do
      [path] -> File.read!(path) |> String.split("\n")
      _ -> raise ArgumentError, "usage: p1.exs *input_path*"
    end
  end

  def parse_node(network, line) do
    [id | [targets | _]] = String.split(line, " = ")
    [left | [right | _]] = Regex.scan(~r/\w+/, targets) |> Enum.map(&hd/1)

    Map.put(network, id, %{?L => left, ?R => right})
  end

  @start "AAA"
  @terminus "ZZZ"
  def traverse_network(_network, @terminus, _path, _initial, len), do: len

  def traverse_network(network, id, [last_dir], initial, len) do
    traverse_network(network, network[id][last_dir], initial, initial, len + 1)
  end

  def traverse_network(network, id, [cur_dir | remaining], initial, len) do
    traverse_network(network, network[id][cur_dir], remaining, initial, len + 1)
  end

  def solve() do
    [path | nodes] =
      read_input(System.argv()) |> Enum.filter(fn line -> line != "" end)

    path = String.to_charlist(path)
    network = Enum.reduce(nodes, %{}, fn node, acc -> parse_node(acc, node) end)

    IO.puts(traverse_network(network, @start, path, path, 0))
  end
end

Solution.solve()
