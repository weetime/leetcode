use crate::list::list_node::ListNode;

pub struct Solution {}
impl Solution {
    /**
     * 解题思路，由于不能影响head，所以需要引入curr，操作指针(引用)
     * 常规边界如果head是空，直接返回
     * 解题的关键是先取第一个节点的val值，并绑定到dp变量中，依次遍历链表节点
     * 若遇到节点val=dp，则继续遍历节点，直到遇到节点val!=dp时，更新dp的值为当前节点val，然后继续按上面的策略继续遍历，直到遍历完整个链表
     * 整个过程时间复杂度O(n)，空间复杂度O(1)
     */
    pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut head = head;
        let mut curr = head.as_mut()?;
        let mut dp = curr.val;
        while let Some(next) = curr.next.as_mut() {
            if next.val == dp {
                curr.next = next.next.take();
            } else {
                dp = next.val; // 重新复制对比
                curr = curr.next.as_mut()?; // 往后移动一步
            }
        }

        head
    }

    /**
     * 为了练习所有权转移，尝试了
     * delete_duplicates3 &mut head 对应 &mut Option<Box<ListNode>>
     * delete_duplicates2 head.as_mut() 对应 Option<&mut Box<ListNode>>
     * delete_duplicates1 head.as_mut().unwrap().as_mut() 对应 &mut ListNode
     */
    pub fn delete_duplicates1(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut head = head;
        let mut curr = head.as_mut().unwrap().as_mut();
        let mut dp = curr.val;
        while let Some(next) = curr.next.take() {
            if next.val == dp {
                curr.next = next.next;
            } else {
                dp = next.val; // 重新复制对比
                curr.next = Some(next); // 放回去
                curr = curr.next.as_mut().unwrap().as_mut();
            }
        }

        head
    }

    // 每次比对前后节点，这样能减少对比的次数
    pub fn delete_duplicates2(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut head = head;
        let mut curr = head.as_mut(); // 引用，指针复制，否则所有权就转移了

        while curr.as_ref().unwrap().next.is_some() {
            if curr.as_ref().unwrap().val == curr.as_ref().unwrap().next.as_ref().unwrap().val {
                let next = curr.as_mut().unwrap().next.as_mut().unwrap().next.take();
                curr.as_mut().unwrap().next = next;
            } else {
                curr = curr.unwrap().next.as_mut();
            }
        }
        head
    }

    pub fn delete_duplicates3(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut head = head;
        let mut curr: &mut Option<Box<ListNode>> = &mut head;

        while curr.as_ref().unwrap().next.is_some() {
            if curr.as_ref().unwrap().val == curr.as_ref().unwrap().next.as_ref().unwrap().val {
                let next = curr.as_mut().unwrap().next.as_mut().unwrap().next.take();
                curr.as_mut().unwrap().next = next;
            } else {
                curr = &mut curr.as_mut().unwrap().next;
            }
        }
        head
    }

    // 还可以用递归的方式处理，时间复杂度O(n)，空间复杂度O(n)
    pub fn delete_duplicates4(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        Self::dfs(head, None)
    }

    fn dfs(head: Option<Box<ListNode>>, prev_val: Option<i32>) -> Option<Box<ListNode>> {
        match head {
            Some(mut head) => {
                let curr_val = Some(head.val);
                let next_node = Self::dfs(head.next.take(), curr_val);
                if curr_val == prev_val {
                    next_node
                } else {
                    head.next = next_node;
                    Some(head)
                }
            }
            None => None,
        }
    }

    // 借助于ChatGPT实现的代码
    pub fn delete_duplicates5(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut head = head;
        let mut prev = &mut head;

        while let Some(curr) = prev.as_mut().and_then(|node| node.next.take()) {
            if prev.as_ref().unwrap().val == curr.val {
                prev.as_mut().unwrap().next = curr.next;
            } else {
                prev = &mut prev.as_mut().unwrap().next;
                *prev = Some(curr);
            }
        }

        head
    }
}

#[test]
fn test() {
    let head_a = ListNode::from_array_by_head(vec![1, 1, 2, 3, 3, 4]);
    let target = ListNode::from_array_by_head(vec![1, 2, 3, 4]);
    let res = Solution::delete_duplicates(head_a.clone());
    // println!("{:?}", res);
    assert_eq!(res, target);

    let res = Solution::delete_duplicates1(head_a.clone());
    assert_eq!(res, target);

    let res = Solution::delete_duplicates2(head_a.clone());
    assert_eq!(res, target);

    let res = Solution::delete_duplicates3(head_a.clone());
    assert_eq!(res, target);

    let res = Solution::delete_duplicates4(head_a.clone());
    assert_eq!(res, target);

    let res = Solution::delete_duplicates5(head_a.clone());
    assert_eq!(res, target);
}
