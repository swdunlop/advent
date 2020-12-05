use std::collections::HashSet;
use std::fs;
use std::path;

fn main() {
    solve2();
}

fn solve1() {
    let input: HashSet<i32> = load_integer_set("2020/01-report-repair/input.txt");
    for entry in &input {
        let other = 2020 - entry;
        if input.contains(&other) {
            println!("{}", other * entry);
            return;
        }
    }
}

fn solve2() {
    let input: HashSet<i32> = load_integer_set("2020/01-report-repair/input.txt");
    for entry1 in &input {
        for entry2 in &input {
            let entry3 = 2020 - entry1 - entry2;
            if input.contains(&entry3) {
                println!("{}", entry3 * entry1 * entry2);
                return;
            }
        }
    }
}

fn load_integer_set<P: AsRef<path::Path>>(path: P) -> HashSet<i32> {
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
