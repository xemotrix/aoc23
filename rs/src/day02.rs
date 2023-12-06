pub fn run() -> [u32; 2] {
    let input = include_str!("../inputs/input02.txt");
    let parsed = parse(input);
    let p1 = part01(&parsed);
    let p2 = part02(&parsed);
    [p1, p2]
}

const THRES: [u32; 3] = [12, 13, 14];

fn part01(input: &[Game]) -> u32 {
    input
        .iter()
        .filter(|g| {
            g.results
                .iter()
                .all(|r| r.iter().zip(THRES.iter()).all(|(x, th)| x <= th))
        })
        .map(|g| g.num)
        .sum()
}

fn part02(input: &[Game]) -> u32 {
    input
        .iter()
        .map(|g| {
            g.results.iter().fold([0, 0, 0], |mut acc, r| {
                (0..3usize).for_each(|i| {
                    if r[i] > acc[i] {
                        acc[i] = r[i];
                    }
                });
                acc
            })
        })
        .map(|r| r.iter().product::<u32>())
        .sum()
}

#[derive(Debug)]
struct Game {
    num: u32,
    results: Vec<[u32; 3]>,
}

fn parse(input: &str) -> Vec<Game> {
    input
        .lines()
        .map(parse_line)
        .into_iter()
        .enumerate()
        .map(|(i, mut g)| {
            g.num = u32::try_from(i).unwrap() + 1u32;
            g
        })
        .collect()
}

fn parse_line(line: &str) -> Game {
    Game {
        num: 0,
        results: line
            .split(':')
            .last()
            .unwrap()
            .split(';')
            .map(parse_result)
            .collect(),
    }
}

fn parse_result(line: &str) -> [u32; 3] {
    let mut result = [0, 0, 0];

    line.split(',').for_each(|x| {
        let digit = parse_one_digit(x);
        if x.ends_with("red") {
            result[0] = digit;
        } else if x.ends_with("green") {
            result[1] = digit;
        } else if x.ends_with("blue") {
            result[2] = digit;
        }
    });

    result
}

fn parse_one_digit(line: &str) -> u32 {
    line.split_whitespace()
        .next()
        .unwrap()
        .parse::<u32>()
        .unwrap()
}
