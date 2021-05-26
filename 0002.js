/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
    let s1 = l1.val
    let s2 = l2.val
    let n = l1
    let n2 = l2
    let carry = 0
    let total = 0
    let rln = new ListNode()
    let head = rln
    while(n || n2) {
        total = s1 + s2 + carry
        carry = tMath.floor(total / 10)
        head.next = new ListNode(total % 10)
        head = head.next
        if(n) {
            n = n.next
        }
        if(n2) {
            n2 = n2.next
        }
        s1 = n ? n.val : 0
        s2 = n2 ? n2.val : 0
    }
    if (carry > 0 ) { head.next = new ListNode(1) }

    return rln.next
};
