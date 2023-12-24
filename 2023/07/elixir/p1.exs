defmodule Solution do
  defmodule Hand do
    defstruct [:cards, :bid, combo: nil]
  end

  @combos %{
    nil => 0,
    :high_card => 1,
    :one_pair => 2,
    :two_pair => 3,
    :three_kind => 4,
    :full_house => 5,
    :four_kind => 6,
    :five_kind => 7
  }
  @card_order ~c"23456789TJQKA"
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

  def find_combo(cards) do
    freqs =
      Enum.frequencies(String.to_charlist(cards))
      |> Enum.flat_map(fn {_k, v} -> [v] end)
      |> Enum.frequencies()

    key =
      cond do
        Map.get(freqs, 5, 0) > 0 -> :five_kind
        Map.get(freqs, 4, 0) > 0 -> :four_kind
        Map.get(freqs, 3, 0) > 0 and Map.get(freqs, 2, 0) > 0 -> :full_house
        Map.get(freqs, 3, 0) > 0 -> :three_kind
        Map.get(freqs, 2, 0) == 2 -> :two_pair
        Map.get(freqs, 2, 0) == 1 -> :one_pair
        Map.get(freqs, 1, 0) == 5 -> :high_card
        true -> nil
      end

    @combos[key]
  end

  # true is returned when first is <= second
  def compare_hands(first, second) do
    cond do
      first.combo == 0 and second.combo != 0 ->
        true

      first.combo != 0 and second.combo == 0 ->
        false

      first.combo != 0 and second.combo != 0 and first.combo < second.combo ->
        true

      first.combo != 0 and second.combo != 0 and first.combo > second.combo ->
        false

      # both combos equal
      true ->
        Enum.zip(String.to_charlist(first.cards), String.to_charlist(second.cards))
        |> Enum.reduce_while(false, fn {a, b}, _acc ->
          first_idx = Enum.find_index(@card_order, fn e -> e == a end)
          second_idx = Enum.find_index(@card_order, fn e -> e == b end)

          cond do
            first_idx > second_idx -> {:halt, false}
            first_idx < second_idx -> {:halt, true}
            first_idx == second_idx -> {:cont, true}
          end
        end)
    end
  end

  def parse_hand(line) do
    [cards | [bid | _]] = String.split(line)

    %Hand{cards: cards, bid: String.to_integer(bid), combo: find_combo(cards)}
  end

  def solve() do
    read_input(System.argv())
    |> Enum.map(&parse_hand/1)
    |> Enum.sort(&compare_hands/2)
    |> Enum.with_index()
    |> Enum.reduce(0, fn {hand, rank}, acc -> acc + hand.bid * (rank + 1) end)
    |> IO.inspect(limit: :infinity)
  end
end

Solution.solve()
