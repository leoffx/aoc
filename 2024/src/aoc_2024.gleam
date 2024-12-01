import argv
import days/day01/day01
import gleam/int
import gleam/io

fn run_day(day: String) {
  case int.parse(day) {
    Ok(1) -> day01.run()
    _ -> io.print("Day not implemented")
  }
}

pub fn main() {
  let args = argv.load().arguments
  case args {
    ["day", day] -> run_day(day)
    _ -> io.print("Usage: gleam run day <day>")
  }
}
