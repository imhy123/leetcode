fn sort(nums: &mut Vec<i32>, l: usize, r: usize) {
    if l >= r {
        return;
    }

    let mut i = l;
    let mut j = r;
    let mid = nums[ (l+r) / 2 ];

    while i <= j {
        while nums[i] < mid {
            i += 1;
        }
        while nums[j] > mid {
            j -= 1;
        }

        if i <= j {
            let t = nums[i];
            nums[i] = nums[j];
            nums[j] = t;

            i += 1;
            if j > 0 {
                j -= 1;
            }
        }
    }

    sort(nums, l, j);
    sort(nums, i, r);
}

impl Solution {
    

    pub fn next_permutation(nums: &mut Vec<i32>) {
        if nums.len() <= 1 {
            return
        }

        let mut found = false;
        let mut cur_val = nums[ nums.len() - 1 ];
        let mut steal_idx: usize = 0;

        for i in (0..nums.len()-1).rev() {
            if nums[i] < cur_val {
                found = true;

                cur_val = nums[i];
                steal_idx = i;
                break;
            } else {
                cur_val = nums[i];
            }
        }

        if !found {
            sort(nums, 0, nums.len() - 1);
            return;
        }

        let mut min_val = -1;
        let mut min_idx: usize = 0;

        let start = (steal_idx +1) as usize;

        for i in start..nums.len() {
            if min_val == -1 || (nums[i] < min_val && nums[i] > cur_val) {
                min_val = nums[i];
                min_idx = i;
            }
        }

        // swap
        let t = nums[min_idx];

        nums[min_idx] = nums[steal_idx as usize];
        nums[steal_idx as usize] = t;

        // sort
        sort(nums, (steal_idx + 1) as usize, nums.len() - 1);
    }
}

struct Solution{}

fn main() {
    let mut v = vec![1,1,5];

    Solution::next_permutation(&mut v);

    println!("{:?}", v);
}