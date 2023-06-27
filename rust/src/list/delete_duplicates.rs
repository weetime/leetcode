use crate::list::list_node::ListNode;

pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    if head.is_none() {
        return head;
    }
    let mut head = head;
    let mut curr = head.as_mut().unwrap(); // 操作的是box的引用
    let mut dp = curr.val;
    while let Some(next) = curr.next.take() {
        // 先夺 此时curr.next = Node
        if next.val == dp {
            curr.next = next.next;
        } else {
            dp = next.val; // 重新复制对比
            curr.next = Some(next); // 放回去
            curr = curr.next.as_mut().unwrap(); // 往后移动一步
        }
    }

    head
}

pub fn delete_duplicates1(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
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

pub fn delete_duplicates2(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    dfs(head, None)
}

fn dfs(head: Option<Box<ListNode>>, prev_val: Option<i32>) -> Option<Box<ListNode>> {
    match head {
        Some(mut head) => {
            let curr_val = Some(head.val);
            let next_node = self::dfs(head.next.take(), curr_val);
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
