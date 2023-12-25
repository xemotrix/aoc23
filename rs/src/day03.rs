pub fn run() -> [u32; 2] {
    let input: &mut str = &mut include_str!("../inputs/input03.txt").to_owned();

    let parsed = parse(input);
    let p1 = part1(parsed);
    let p2 = part2(parse(input));
    [p1, p2]
}

fn parse(input: &mut str) -> Vec<char> {
    input.chars().collect()
}

fn part2(mut input: Vec<char>) -> u32 {
    let mut sum = 0;
    for i in 0..input.len() {
        if input[i] == '*' {
            let gear_nums = search_gears(i, &mut input);
            if gear_nums.len() != 2 {
                continue;
            }
            sum += gear_nums[0] * gear_nums[1];
        }
    }
    sum
}

fn search_gears(idx: usize, input: &mut [char]) -> Vec<u32> {
    let mut nums = vec![];
    let width = input.iter().position(|c| *c == '\n').unwrap() as i32;
    for i in get_indices(idx, width as usize) {
        if let Some(c) = input.get(i) {
            if c.is_ascii_digit() {
                nums.push(get_num(i, input));
            }
        }
    }
    nums
}

fn part1(mut input: Vec<char>) -> u32 {
    let mut sum = 0;
    for i in 0..input.len() {
        if !is_symbol(&input[i]) {
             continue;
        }
        sum += search_numbers(i, &mut input)
            .iter()
            .fold(0, |acc, x| acc + *x);
    }
    sum
}

fn search_numbers(idx: usize, input: &mut [char]) -> Vec<u32> {
    let width = input.iter().position(|c| *c == '\n').unwrap() as i32;
    let indices = get_indices(idx, width as usize);
    let mut nums = vec![];
    for i in indices {
        if let Some(c) = input.get(i) {
            if c.is_ascii_digit() {
                nums.push(get_num(i, input));
            }
        }
    }
    nums
}

fn get_indices(i: usize, width: usize) -> Vec<usize> {
    vec![
        i - (width + 1) - 1,
        i - (width + 1),
        i - (width + 1) + 1,
        i - 1,
        i + 1,
        i + (width + 1) - 1,
        i + (width + 1),
        i + (width + 1) + 1,
    ]
}

fn get_num(idx: usize, input: &mut [char]) -> u32 {
    let offset = input[idx..]
        .iter()
        .position(|c| !c.is_ascii_digit())
        .unwrap();

    input[..idx + offset]
        .iter_mut()
        .rev()
        .take_while(|c| c.is_ascii_digit())
        .enumerate()
        .fold(0, |acc, (i, c)| {
            let num = acc + c.to_digit(10).unwrap() * u32::pow(10, i as u32);
            *c = '.';
            num
        })
}


fn is_symbol(c: &char) -> bool {
    *c != '.' && !c.is_ascii_digit() && *c != '\n'
}
