use std::collections::HashSet;
use std::fs;
use std::iter::FromIterator;
use std::path;

fn main() {
    solve2();
}

fn solve1() {
    let input = load_integers("2020/01-report-repair/input.txt");
    let table: HashSet<&i32> = HashSet::from_iter(&input);
    for entry in &input {
        let other = 2020 - entry;
        if table.contains(&other) {
            println!("{}", other * entry);
            return;
        }
    }
}

fn solve2() {
    let input = load_integers("2020/01-report-repair/input.txt");
    let table: HashSet<&i32> = HashSet::from_iter(&input);
    for (skip, entry1) in input.iter().enumerate() {
        for entry2 in input[skip..].iter() {
            let entry3 = 2020 - entry1 - entry2;
            if table.contains(&entry3) {
                println!("{}", entry3 * entry1 * entry2);
                return;
            }
        }
    }
}

fn load_integers<P: AsRef<path::Path>>(path: P) -> Vec<i32> {
    return fs::read_to_string(path)
        .expect("could not read input")
        .trim_end_matches("\n")
        .split("\n")
        .map(|line| {
            line.parse::<i32>()
                .expect("could not parse line as integer")
        })
        .collect();
}
