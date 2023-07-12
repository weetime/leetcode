#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

/**
 * 内置常见的链表操作
 * 1.构建链表的两种形式，头插法和尾插法
 * 2.取链表的长度
 * 3.找中点
 * 4.反转链表
 * 5.合并两个链表
 * 6.按索引删除节点
 * 7.按值删除节点
 * 8.排序
 */
impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }

    // 头插法构建链表 [1,2,3,4,5] => 1->2->3->4->5
    pub fn from_array_by_head(arr: Vec<i32>) -> Option<Box<ListNode>> {
        let mut arr = arr;
        let mut head = None;
        while let Some(x) = arr.pop() {
            let mut node = Box::new(Self::new(x));
            node.next = head;
            head = Some(node);
        }
        head
    }

    // 尾插法构建链表 [1,2,3,4,5] => 1->2->3->4->5
    pub fn from_array_by_tail(arr: Vec<i32>) -> Option<Box<ListNode>> {
        let mut dummy = ListNode { val: 0, next: None };
        let mut tail = &mut dummy;
        for x in arr.iter() {
            let node = Box::new(ListNode::new(*x));
            tail.next = Some(node);
            tail = tail.next.as_mut().unwrap().as_mut();
        }
        dummy.next
    }
}
