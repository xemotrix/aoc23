use regex::Regex;

pub fn run() -> [u32; 2] {
    let input: &mut str = &mut include_str!("../inputs/input03.txt").to_owned();

    println!("{}", input);
    let parsed = parse(input);
    [1, 2]
}

fn part01(numbers: Vec<Number>, symbols: Vec<Symbol>) -> u32 {
    symbols.iter().map(|s| {
        numbers.iter().map(|n| {
            if n.from <= s.idx && s.idx <= n.to {
                n.num
            } else {
                0
            }
        })
    });
    0
}

#[derive(Debug)]
struct Number {
    num: i32,
    from: usize,
    to: usize,
}

#[derive(Debug)]
struct Symbol {
    sym: char,
    idx: usize,
}

fn parse(input: &mut str) -> (Vec<Number>, Vec<Symbol>) {
    let r_num = Regex::new(r"\d+").unwrap();

    let v = r_num
        .captures_iter(input)
        .filter_map(|m| match m.get(0) {
            Some(mch) => {
                let num = Number {
                    num: mch.as_str().parse::<i32>().unwrap(),
                    from: mch.start(),
                    to: mch.end(),
                };
                Some(num)
            }
            None => None,
        })
        .collect::<Vec<Number>>();

    let s = input
        .chars()
        .enumerate()
        .filter_map(|(i, c)| {
            if is_symbol(&c) {
                Some(Symbol { sym: c, idx: i })
            } else {
                None
            }
        })
        .collect::<Vec<Symbol>>();

    (v, s)
}

fn is_symbol(c: &char) -> bool {
    *c != '.' && !('0'..='9').contains(c)
}

// fn is_symbol(c: &char) -> bool {
//     *c != '.' && !('0'..='9').contains(c)
// }
