pub mod day01;
pub mod day02;
pub mod day03;
use std::time::Instant;

fn main() {
    let start = Instant::now();

    day03::run();
    // run_all();

    println!("Time elapsed: {:?}", start.elapsed());
}

#[allow(dead_code)]
fn run_all() {
    let mut res_str = "".to_string();
    res_str += &fmt_res(1, day01::run());
    res_str += &fmt_res(2, day02::run());
    println!("{}", res_str);
}

fn fmt_res(day: i32, res: [u32; 2]) -> String {
    format!("DAY {}:\n- part1: {}\n- part2: {}\n", day, res[0], res[1])
}
