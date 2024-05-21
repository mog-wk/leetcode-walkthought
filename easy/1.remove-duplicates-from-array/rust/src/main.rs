fn main() {
    println!("Hello, world!");

    let mut arr = [-3, -3, -2, -1, -1, 0, 0, 0, 0, 0];
    //let mut arr = [1, 2, 2, 3];
    let r = remove_dup(&mut arr);

    dbg!(arr, r);
}

fn remove_dup(arr: &mut [i32]) -> i32 {
    let mut l = 1;
    for i in 1..arr.len() {
        if arr[i] != arr[i - 1] {
            arr[l] = arr[i];
            l += 1;
        }
    }
    l as i32
}
