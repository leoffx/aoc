import gleam/bool
import gleam/int
import gleam/io
import gleam/list
import gleam/string
import utils

pub fn run() {
  let res = part1()
  // part2()
  io.print("Res: " <> int.to_string(res))
  Nil
}

pub fn process_file(file: String) {
  let assert Ok(content) = utils.read_file(file)

  let assert Ok(numbers) =
    content
    |> string.trim()
    |> string.split("\n")
    |> list.try_map(fn(line) {
      string.split(line, on: " ")
      |> list.filter(fn(s) { string.is_empty(s) == False })
      |> list.try_map(fn(s) {
        case int.parse(s) {
          Ok(num) -> Ok(num)
          Error(_) -> Error("Failed to parse number: " <> s)
        }
      })
    })
  numbers
}

pub fn check_line_safe(line: List(Int)) -> Bool {
  let order = case list.take(line, 2) {
    [a, b] if a > b -> Ok("desc")
    [a, b] if a < b -> Ok("asc")
    _ -> Error("Invalid order")
  }
  case order {
    Error(_) -> False
    Ok(order) ->
      list.fold_until(line, 0, fn(acc, num) {
        case acc {
          0 -> list.Continue(num)
          _ if order == "asc" && num > acc && num - acc >= 1 && num - acc <= 3 -> {
            list.Continue(num)
          }
          _ if order == "desc" && num < acc && acc - num >= 1 && acc - num <= 3 ->
            list.Continue(num)
          _ -> list.Stop(0)
        }
      })
      != 0
  }
}

pub fn part1() -> Int {
  let file = "./src/days/day02/input.txt"
  let lines = process_file(file)
  list.fold(lines, 0, fn(acc, line) {
    case check_line_safe(line) {
      True -> acc + 1
      False -> acc
    }
  })
}

fn part2() {
  let _file = "./src/days/day02/input.txt"
  // let #(col1, col2) = process_file(file)
  Nil
}
