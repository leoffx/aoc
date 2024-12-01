import gleam/io
import utils

pub fn run() {
  part1()
}

fn part1() {
  let file = "./src/days/example.txt"
  let assert Ok(content) = utils.read_file(file)
  Nil
}
