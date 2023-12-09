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
    Checks whether a given string starts with a literal digit
    Returns the matching digit and length of the matched word
  """
  def match_digit_word(line) do
    cond do
      String.starts_with?(line, "one") -> {:ok, 1}
      String.starts_with?(line, "two") -> {:ok, 2}
      String.starts_with?(line, "three") -> {:ok, 3}
      String.starts_with?(line, "four") -> {:ok, 4}
      String.starts_with?(line, "five") -> {:ok, 5}
      String.starts_with?(line, "six") -> {:ok, 6}
      String.starts_with?(line, "seven") -> {:ok, 7}
      String.starts_with?(line, "eight") -> {:ok, 8}
      String.starts_with?(line, "nine") -> {:ok, 9}
      true -> {:none, 0}
    end
  end

  def process_line("", acc), do: acc

  def process_line(line, acc) do
    case match_digit_word(line) do
      {:ok, digit} ->
        process_line(String.slice(line, 1..-1), [digit | acc])

      _ ->
        case String.at(line, 0) |> Integer.parse() do
          {digit, _rest} -> process_line(String.slice(line, 1..-1), [digit | acc])
          _ -> process_line(String.slice(line, 1..-1), acc)
        end
    end
  end

  @doc """
    Algo:
    1. Read file line by line
    2. Process each line by keeping digits and replacing digit words
    3. Select the first and last numbers from lists and save them in an enum
    4. Sum every enum
  """
  def solve() do
    read_input(System.argv())
    |> Enum.reduce([], fn line, acc -> acc ++ [Enum.reverse(process_line(line, []))] end)
    |> Enum.reduce(0, fn digits, acc ->
      acc + Enum.at(digits, 0) * 10 + Enum.at(digits, -1)
    end)
    |> IO.inspect()
  end
end

Solution.solve()
