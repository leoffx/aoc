import days/day3/day3
import gleeunit/should

pub fn find_mul_test() {
  let text = "aaaamul(2,4)bbbbmul(3,1)bb"
  day3.find_mul(text)
  |> should.equal([[2, 4], [3, 1]])
}

pub fn part1_test() {
  let file = "./src/days/day3/example.txt"
  let input = day3.process_file(file)
  day3.part1(input)
  |> should.equal(161)
}
