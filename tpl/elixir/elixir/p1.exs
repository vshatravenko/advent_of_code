defmodule Solution do
  @input_path "../input.txt"
  def read_input([]) do
    File.stream!(@input_path)
  end

  def read_input(args) do
    case args do
      [path] -> File.stream!(path)
      _ -> raise ArgumentError, "usage: p1.exs *input_path*"
    end
  end

  def solve() do
    read_input(System.argv())
  end
end

Solution.solve()
