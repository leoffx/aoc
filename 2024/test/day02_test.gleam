import days/day02/day02
import gleeunit/should

pub fn check_line1_safe_test() {
  let line = [7, 6, 4, 2, 1]
  day02.check_line_safe(line)
  |> should.equal(True)
}

pub fn check_line2_safe_test() {
  let line = [1, 2, 7, 8, 9]
  day02.check_line_safe(line)
  |> should.equal(False)
}

pub fn check_line3_safe_test() {
  let line = [9, 7, 6, 2, 1]
  day02.check_line_safe(line)
  |> should.equal(False)
}

pub fn check_line4_safe_test() {
  let line = [1, 3, 2, 4, 5]
  day02.check_line_safe(line)
  |> should.equal(False)
}

pub fn check_line5_safe_test() {
  let line = [8, 6, 4, 4, 1]
  day02.check_line_safe(line)
  |> should.equal(False)
}

pub fn check_line6_safe_test() {
  let line = [1, 3, 6, 7, 9]
  day02.check_line_safe(line)
  |> should.equal(True)
}
