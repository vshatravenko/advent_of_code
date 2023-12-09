defmodule Solution do
  @input_path "input.txt"
  def read_input([]) do
    File.stream!(@input_path)
  end

  def read_input(args) do
    case args do
      [path] -> File.stream!(path)
      _ -> raise ArgumentError, "usage: solution_a.exs *input_path*"
    end
  end

  @doc """
    Algo:
    1. Read file line by line
    2. Filter out numbers only
    3. Select the first and last numbers from lists and save them in an enum
    4. Sum every enum
  """
  def solve() do
    read_input(System.argv())
    |> Stream.map(&String.to_charlist/1)
    |> Enum.reduce([], fn line, acc ->
      acc ++ [Enum.filter(line, fn c -> c >= ?0 && c <= ?9 end)]
    end)
    |> Enum.reduce(0, fn digits, acc ->
      acc + (Enum.at(digits, 0) - ?0) * 10 + (Enum.at(digits, -1) - ?0)
    end)
    |> IO.inspect()
  end
end

Solution.solve()
