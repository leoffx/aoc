import argv
import days/day1/day1
import days/day2/day2
import days/day3/day3
import gleam/int
import gleam/io

fn run_day(day: String) {
  case int.parse(day) {
    Ok(1) -> day1.run()
    Ok(2) -> day2.run()
    Ok(3) -> day3.run()
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
