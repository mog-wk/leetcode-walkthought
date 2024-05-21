#![allow(unused)]
fn main() {
    println!("Hello, world!");

    let tests = [
        vec![7, 1, 5, 3, 6, 4],
        vec![7, 6, 4, 3, 1],
        vec![1, 2, 3, 4, 5],
    ];

    for p in tests {
        let r = max_profit(&p.to_vec()); // 7

        dbg!(r);
    }
}

fn max_profit(prices: &Vec<i32>) -> i32 {
    let mut sum = 0;
    for win in prices.windows(2) {
        if win[1] > win[0] {
            sum += win[1] - win[0];
        }
    }
    sum
}
