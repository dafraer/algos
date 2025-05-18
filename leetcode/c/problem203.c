 /**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* removeElements(struct ListNode* head, int val) {
    if (head== NULL) return NULL;
    struct ListNode *prev, *h;
    struct ListNode d;
    d.next = head;
    h = head;
    prev = &d;
    bool isFirst = true;
    while (head) {
        if (head->val == val) {
            prev->next = head->next;   
            head = head->next; 
            if (isFirst) h = head;
        } else {
            prev = head;
            head = head->next; 
            isFirst = false;
        }
    }
    return h;
}
