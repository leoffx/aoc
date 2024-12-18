import days/day2/day2
import gleeunit/should

pub fn check_line1_safe_test() {
  let line = [7, 6, 4, 2, 1]
  day2.check_line_safe(line)
  |> should.equal(True)
}

pub fn check_line2_safe_test() {
  let line = [1, 2, 7, 8, 9]
  day2.check_line_safe(line)
  |> should.equal(False)
}

pub fn check_line3_safe_test() {
  let line = [9, 7, 6, 2, 1]
  day2.check_line_safe(line)
  |> should.equal(False)
}

pub fn check_line4_safe_test() {
  let line = [1, 3, 2, 4, 5]
  day2.check_line_safe(line)
  |> should.equal(False)
}

pub fn check_line5_safe_test() {
  let line = [8, 6, 4, 4, 1]
  day2.check_line_safe(line)
  |> should.equal(False)
}

pub fn check_line6_safe_test() {
  let line = [1, 3, 6, 7, 9]
  day2.check_line_safe(line)
  |> should.equal(True)
}

pub fn remove_one_at_test() {
  let list = [1, 2, 3, 4, 5]
  day2.remove_one_at(list, 2)
  |> should.equal([1, 2, 4, 5])
}

pub fn create_sublists_test() {
  let list = [1, 2, 3, 4, 5]
  day2.create_sublists(list)
  |> should.equal([
    [1, 2, 3, 4, 5],
    [2, 3, 4, 5],
    [1, 3, 4, 5],
    [1, 2, 4, 5],
    [1, 2, 3, 5],
    [1, 2, 3, 4],
  ])
}
