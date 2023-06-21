pub mod list;

use core::num;
use list::reverse;
use std::borrow::Borrow;
use std::cell::RefCell;
use std::cmp::min;
use std::collections::HashMap;
use std::rc::Rc;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

fn main() {
    // pub fn max_path_sum(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
    //     let mut ans = i32::MIN;

    //     fn dfs(root: &Option<Rc<RefCell<TreeNode>>>, ans: &mut i32) -> i32 {
    //         if root.is_none() {
    //             return 0;
    //         }
    //         let node = root.as_ref().unwrap().borrow();
    //         let l = std::cmp::max(0, dfs(&node.left, ans));
    //         let r = std::cmp::max(0, dfs(&node.right, ans));
    //         *ans = std::cmp::max(*ans, l + r + node.val);
    //         return std::cmp::max(l, r) + root.as_ref().unwrap().borrow().val;
    //     }
    //     dfs(&root, &mut ans);
    //     return ans;
    // }

    // let res = max_path_sum(getTreeNode());
    // println!("{:?}", res);
}

#[test]
fn test_124() {
    // let mut ans: i32 = i32::MIN;
    // fn dfs(root: &Option<Rc<RefCell<TreeNode>>>, ans: &mut i32) -> i32 {
    //     match root {
    //         None => return 0,
    //         Some(x) => {
    //             let node = &x.borrow();
    //             let l = i32::max(0, dfs(&node.left, ans));
    //             let r = i32::max(0, dfs(&node.right, ans));
    //             *ans = i32::max(*ans, l + r + node.val);
    //             return i32::max(l, r) + node.val;
    //         }
    //     }
    // }

    // let root = getTreeNode();
    // dfs(&root, &mut ans);
    // println!("{:?}", ans);
}

fn getTreeNode() -> Option<Rc<RefCell<TreeNode>>> {
    let node1 = Some(Rc::new(RefCell::new(TreeNode::new(15))));
    let node2 = Some(Rc::new(RefCell::new(TreeNode::new(7))));
    let mut node3 = TreeNode::new(20);
    node3.left = node1;
    node3.right = node2;
    let mut node4 = TreeNode::new(-10);
    node4.left = Some(Rc::new(RefCell::new(TreeNode::new(9))));
    node4.right = Some(Rc::new(RefCell::new(node3)));

    return Some(Rc::new(RefCell::new(node4)));
}

#[test]
fn test_fib() {
    fn fib(n: i32) -> i32 {
        if n == 0 {
            return 0;
        }
        if n <= 2 {
            return 1;
        }

        return fib(n - 2) + fib(n - 1);
    }
    let result = fib(30);
    println!("{:?}", result);
}

#[test]
fn test_fib1() {
    fn fib(n: i32) -> i32 {
        let n: usize = n as usize;
        let mut memo: Vec<i32> = vec![0; n + 1];
        return helper(&mut memo, n);
    }
    fn helper(memo: &mut Vec<i32>, n: usize) -> i32 {
        if n == 0 {
            return 0;
        }
        if n <= 2 {
            return 1;
        }
        if memo[n] != 0 {
            return memo[n];
        }
        memo[n] = helper(memo, n - 2) + helper(memo, n - 1);
        return memo[n];
    }
    let result = fib(10);
    println!("{:?}", result);
}

#[test]
fn test_fib2() {
    fn fib(n: i32) -> i32 {
        if n == 0 {
            return 0;
        }
        if n <= 2 {
            return 1;
        }
        let n: usize = n as usize;
        let mut dp: Vec<i32> = vec![0; n + 1];
        dp[0] = 0;
        dp[1] = 1;
        for i in 2..=n {
            dp[i] = dp[i - 2] + dp[i - 1];
        }
        return dp[n];
    }

    let result = fib(30);
    println!("{:?}", result);
}

#[test]
fn test_fib3() {
    fn fib(n: i32) -> i32 {
        if n < 2 {
            return n;
        }
        let (mut prev, mut curr) = (0, 1);
        for _ in 0..n - 1 {
            let sum = prev + curr;
            prev = curr;
            curr = sum;
        }
        return curr;
    }

    let result = fib(30);
    println!("{:?}", result);
}

#[test]
fn test_coin1() {
    let coins: [i32; 3] = [1, 2, 5];
    fn dp(n: i32, coins: [i32; 3]) -> i32 {
        if n == 0 {
            return 0;
        }
        if n < 0 {
            return -1;
        }
        let mut result = i32::MAX;
        for coin in coins {
            let sub_problem = dp(n - coin, coins);
            if sub_problem == -1 {
                continue;
            }
            result = i32::min(result, sub_problem + 1);
        }

        if result == i32::MAX {
            return -1;
        }
        return result;
    }

    let res = dp(11, coins);
    println!("{:?}", res);
}

#[test]
fn test_coin2() {
    let coins: [i32; 3] = [1, 2, 5];
    let mut memo: HashMap<i32, i32> = HashMap::new();
    fn helper(n: i32, coins: [i32; 3], memo: &mut HashMap<i32, i32>) -> i32 {
        match memo.get(&n) {
            Some(x) => {
                return *x;
            }
            None => {
                if n == 0 {
                    return 0;
                }
                if n < 0 {
                    return i32::MAX;
                }
                let mut res = i32::MAX;
                for coin in coins {
                    let sub_problem = helper(n - coin, coins, memo);
                    if sub_problem == i32::MAX {
                        continue;
                    }
                    res = i32::min(res, 1 + sub_problem);
                }
                if res == i32::MAX {
                    res = -1
                }
                memo.insert(n, res);
                return res;
            }
        }
    }

    let res = helper(11, coins, &mut memo);
    println!("{:?}", res);
}

#[test]
fn test_coin3() {
    let mut res = -1;
    let amount: i32 = 11;
    let coins: [i32; 3] = [1, 2, 5];
    let len = (amount + 1) as usize;
    let mut dp: Vec<i32> = vec![amount + 1; len];
    dp[0] = 0;
    for i in 0..len {
        for coin in &coins {
            if i as i32 - *coin < 0 {
                continue;
            }
            dp[i] = i32::min(dp[i], dp[i - *coin as usize] + 1);
        }
    }

    if dp[amount as usize] != amount + 1 {
        res = dp[amount as usize];
    }
    println!("{:?}", res);
}

#[test]
fn test_coin4() {
    let amount = 11;
    let coins = vec![1, 2, 5];
    let amount = amount as usize;
    // dp[i] 表示组成 i 所需的最少硬币数量
    let mut dp = Vec::with_capacity(amount + 1);
    let mut min;
    dp.push(0);
    for i in 1..=amount {
        min = amount;
        for &n in coins.iter() {
            if i >= (n as usize) {
                min = min.min(dp[i - n as usize]);
            }
        }
        dp.push(min + 1);
    }
    let res = if dp[amount] > amount {
        -1
    } else {
        dp[amount] as i32
    };

    println!("{:?}", res);
}

#[test]
fn test_min() {
    let head = list::ListNode::sample();
    let res = reverse::Solution::reverse_list(head);
    println!("{:?}", res);
}
