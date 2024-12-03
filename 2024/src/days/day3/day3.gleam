import gleam/int
import gleam/io
import gleam/list
import gleam/option
import gleam/regexp
import utils

pub fn process_file(file: String) {
  let assert Ok(content) = utils.read_file(file)
  content
}

pub fn run() {
  let file = "./src/days/day3/input.txt"
  let input = process_file(file)
  let res = part1(input)
  // let res = part2()
  io.print("Res: " <> int.to_string(res))
  Nil
}

pub fn find_mul(input: String) {
  let assert Ok(re) = regexp.from_string("mul\\((\\d+),(\\d+)\\)")
  let matches = regexp.scan(re, input)
  list.map(matches, fn(s) {
    option.all(s.submatches)
    |> option.unwrap([])
    |> list.filter_map(int.parse)
  })
}

pub fn part1(input: String) -> Int {
  let matches = find_mul(input)
  list.fold(matches, 0, fn(acc, match) {
    acc + list.fold(match, 1, fn(a, b) { a * b })
  })
}

fn part2(input) -> Int {
  1
}
