import gleam/int
import gleam/io
import gleam/list
import gleam/string
import utils

pub fn run() {
  // let res = part1()
  let res = part2()
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
  let file = "./src/days/day2/input.txt"
  let lines = process_file(file)
  list.count(lines, check_line_safe)
}

pub fn remove_one_at(my_list, index) -> List(Int) {
  let before = list.take(my_list, index)
  let after = list.drop(my_list, index + 1)
  list.flatten([before, after])
}

pub fn create_sublists(my_list) {
  case my_list {
    [] -> []
    [_] -> []
    _ -> {
      let indices = list.range(0, list.length(my_list) - 1)
      list.flatten([
        [my_list],
        list.map(indices, fn(i) { remove_one_at(my_list, i) }),
      ])
    }
  }
}

fn part2() {
  let file = "./src/days/day2/input.txt"
  let lines = process_file(file)
  list.count(lines, fn(line) {
    let sublists = create_sublists(line)
    list.any(sublists, check_line_safe)
  })
}
