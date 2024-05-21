fn main() {
    println!("Hello, world!");
    let mut arr = vec![1, 2, 3, 4, 5, 6, 7];

    rotate(&mut arr, 3);
    dbg!(arr);
}

fn rotate(nums: &mut Vec<i32>, k: i32) {
    let r = k as usize % nums.len();

    let mid = nums.len() - r;

    let remainder = &nums[0..mid];
    let shifted = &nums[mid..nums.len()];
    let arr = [shifted, remainder].concat();
    *nums = arr.to_vec();
}
