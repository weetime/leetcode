use crate::list::list_node::ListNode;

// 给定一个数组，采用头插法构建链表，给链表进行升序排序，去重后反转链表
#[test]
fn test_demo01() {
    // step1. 头插法构建链表
    let mut head = from_array_by_head(vec![5, 1, 2, 1, 4, 2, 3]);
    let mut target = ListNode::from_array_by_head(vec![5, 1, 2, 1, 4, 2, 3]);
    assert_eq!(head, target);

    // step2.排序
    head = insert_sort(head);
    target = ListNode::from_array_by_head(vec![1, 1, 2, 2, 3, 4, 5]);
    assert_eq!(head, target);

    // step3.去重
    head = delete(head);
    target = ListNode::from_array_by_head(vec![1, 2, 3, 4, 5]);
    assert_eq!(head, target);

    // step4.反转
    head = reverse(head);
    target = ListNode::from_array_by_head(vec![5, 4, 3, 2, 1]);
    assert_eq!(head, target);
}

#[test]
fn test_demo02() {
    // step1. 头插法构建链表
    let mut head = from_array_by_head(vec![5, 1, 2, 1, 4, 2, 3]);
    let mut target = ListNode::from_array_by_head(vec![5, 1, 2, 1, 4, 2, 3]);
    assert_eq!(head, target);

    // step2. 归并排序 优化时间复杂度
    head = sort(head);
    target = ListNode::from_array_by_head(vec![1, 1, 2, 2, 3, 4, 5]);
    assert_eq!(head, target);

    // step3.去重 & 反转
    head = delete_and_reverse(head);
    target = ListNode::from_array_by_head(vec![5, 4, 3, 2, 1]);
    assert_eq!(head, target);
}

fn from_array_by_head(arr: Vec<i32>) -> Option<Box<ListNode>> {
    let mut arr = arr;
    let mut head = None;
    while !arr.is_empty() {
        match arr.pop() {
            Some(val) => {
                let mut node = Box::new(ListNode::new(val));
                node.next = head;
                head = Some(node);
            }
            None => todo!(),
        }
    }
    head
}

// 插入排序
fn insert_sort(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    if head.is_none() {
        return head;
    }
    let mut dummy = ListNode::new(0);
    let mut ptr = head;
    while let Some(mut node) = ptr {
        ptr = node.next.take();
        let mut tail = &mut dummy;
        while tail.next.is_some() && tail.next.as_ref().unwrap().val < node.val {
            tail = tail.next.as_mut().unwrap().as_mut();
        }
        node.next = tail.next.take();
        tail.next = Some(node);
    }

    dummy.next
}

// 去重
fn delete(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    if head.is_none() {
        return head;
    }
    let mut head = head;
    let mut dp = head.as_ref().unwrap().val;
    let mut ptr = head.as_mut().unwrap().as_mut();
    while let Some(mut node) = ptr.next.take() {
        if node.val == dp {
            ptr.next = node.next.take();
        } else {
            dp = node.val;
            ptr.next = Some(node);
            ptr = ptr.next.as_mut().unwrap().as_mut();
        }
    }
    head
}

// 反转
fn reverse(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut head = head;
    let mut new_head = None;
    while let Some(mut node) = head {
        head = node.next.take();
        node.next = new_head;
        new_head = Some(node);
    }
    new_head
}

// 归并排序
fn sort(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    fn cut(head: &mut Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        if head.is_none() {
            return None;
        }
        let mut new_head = ListNode::new(0);
        let mut tail = &mut new_head;
        let mut n = n;
        while n > 0 && head.is_some() {
            if let Some(mut node) = head.take() {
                *head = node.next.take(); // 这里是赋值，不是移动指针，将head里面的值改动了
                tail.next = Some(node);
                tail = tail.next.as_mut().unwrap().as_mut();
            }
            n -= 1;
        }
        new_head.next
    }
    fn merge(
        head_a: Option<Box<ListNode>>,
        head_b: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut dummy = ListNode::new(0);
        let mut tail = &mut dummy;
        let (mut h1, mut h2) = (head_a, head_b);
        while h1.is_some() && h2.is_some() {
            let mut temp;
            if h1.as_ref().unwrap().val <= h2.as_ref().unwrap().val {
                temp = h1.take();
                h1 = temp.as_mut().unwrap().next.take();
            } else {
                temp = h2.take();
                h2 = temp.as_mut().unwrap().next.take();
            }
            tail.next = temp.take();
            tail = tail.next.as_mut().unwrap().as_mut();
        }
        tail.next = if h1.is_some() { h1.take() } else { h2.take() };
        dummy.next
    }

    let mut head = head;
    let mut length = 0;
    let mut ptr = head.as_ref();
    while let Some(node) = ptr.take() {
        ptr = node.next.as_ref();
        length += 1;
    }

    let mut dummy = ListNode::new(0);

    let mut size = 1;
    while size < length {
        let (curr, mut ptr) = (&mut head, &mut dummy);
        while curr.is_some() {
            let left = cut(curr, size);
            let right = cut(curr, size);
            ptr.next = merge(left, right);
            while ptr.next.is_some() {
                ptr = ptr.next.as_mut().unwrap().as_mut();
            }
        }
        size = size << 1;
        head = dummy.next.take();
    }
    head
}

// 去重并反转
fn delete_and_reverse(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    if head.is_none() {
        return head;
    }
    let mut new_head = None;
    let mut head = head;
    let mut dp = head.as_ref().unwrap().val - 1;
    while let Some(mut node) = head {
        head = node.next.take();
        if node.val != dp {
            dp = node.val;
            node.next = new_head;
            new_head = Some(node);
        }
    }
    new_head
}
