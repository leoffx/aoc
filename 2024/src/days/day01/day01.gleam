import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/string
import utils

pub fn run() {
  // part1()
  part2()
}

fn process_file(file: String) {
  let assert Ok(content) = utils.read_file(file)
  content
  |> string.trim()
  |> string.split("\n")
  |> list.filter_map(fn(line) {
    case
      string.split(line, on: " ")
      |> list.filter(fn(s) { string.is_empty(s) == False })
    {
      [a, b] -> {
        case int.parse(a), int.parse(b) {
          Ok(num1), Ok(num2) -> Ok(#(num1, num2))
          _, _ -> Error(Nil)
        }
      }
      _ -> Error(Nil)
    }
  })
  |> list.unzip()
}

pub fn part1() {
  let file = "./src/days/day01/input.txt"
  let #(col1, col2) = process_file(file)
  let sorted_col1 = list.sort(col1, int.compare)
  let sorted_col2 = list.sort(col2, int.compare)

  let result =
    list.zip(sorted_col1, sorted_col2)
    |> list.map(fn(pair) {
      let #(a, b) = pair
      int.absolute_value(a - b)
    })
    |> list.fold(0, fn(acc, x) { acc + x })
  io.debug(result)
  Nil
}

pub fn count_numbers(numbers: List(Int)) {
  list.fold(numbers, dict.new(), fn(counts, number) {
    case dict.get(counts, number) {
      Ok(current_count) -> dict.insert(counts, number, current_count + 1)
      Error(Nil) -> dict.insert(counts, number, 1)
    }
  })
}

fn part2() {
  let file = "./src/days/day01/input.txt"
  let #(col1, col2) = process_file(file)
  let count_map = count_numbers(col2)
  let result =
    list.fold(col1, 0, fn(acc, number) {
      case dict.get(count_map, number) {
        Ok(count) -> acc + number * count
        Error(Nil) -> acc
      }
    })
  io.debug(result)
  Nil
}
