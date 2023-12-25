pub fn run() -> [u32; 2] {
    let input = include_str!("../inputs/input01.txt");
    let p1 = part01(input);
    let p2 = part02(input);
    [p1, p2]
}

fn part01(input: &str) -> u32 {
    input.lines().map(find_num).sum::<u32>()
}

fn find_num(line: &str) -> u32 {
    let first = line
        .chars()
        .find(|c| c.is_ascii_digit())
        .unwrap()
        .to_digit(10)
        .unwrap();
    let last = line
        .chars()
        .rev()
        .find(|c| c.is_ascii_digit())
        .unwrap()
        .to_digit(10)
        .unwrap();
    first * 10 + last
}

const NUMS: [&str; 10] = [
    "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
];

fn part02(input: &str) -> u32 {
    input.lines().map(find_num2).sum::<u32>()
}

fn find_num2(line: &str) -> u32 {
    let first = line
        .chars()
        .scan("".to_string(), |state, c| {
            if c.is_ascii_digit() {
                return Some(c.to_digit(10));
            }
            state.push(c);
            let num: Option<u32> = NUMS
                .into_iter()
                .enumerate()
                .find(|(_, n)| state.ends_with(n))
                .and_then(|(i, _)| i.try_into().ok());

            Some(num)
        })
        .find(|c| c.is_some())
        .unwrap()
        .unwrap();

    let last = line
        .chars()
        .rev()
        .scan("".to_string(), |state, c| {
            if c.is_ascii_digit() {
                return Some(c.to_digit(10));
            }
            state.push(c);
            let num: Option<u32> = NUMS
                .into_iter()
                .map(|n| n.chars().rev().collect::<String>())
                .enumerate()
                .find(|(_, n)| state.ends_with(n))
                .and_then(|(i, _)| i.try_into().ok());

            Some(num)
        })
        .find(|c| c.is_some())
        .unwrap()
        .unwrap();

    first * 10 + last
}
