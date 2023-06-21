pub mod list;
use list::list_node::sample as listSample;
use list::reverse::Solution as listSolution;

fn main() {}

#[test]
// 206. 反转链表
fn test_reverse_list() {
    let res = listSolution::reverse_list2(listSample());
    println!("{:?}", res);
}
