/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        if (!headA && headB)
            return nullptr;
        if (headA && !headB)
            return nullptr;

        int aLen = getLength(headA);
        int bLen = getLength(headB);

        while (true) {
            while (aLen > bLen) {
                headA = headA->next;
                aLen--;
            }
            while (aLen < bLen) {
                headB = headB->next;
                bLen--;
            }
            if (headA != headB) {
                aLen--;
                bLen--;
                headA = headA->next;
                headB = headB->next;
            }
            else {
                break;
            }
        }
        return headA;
    }

    int getLength(ListNode *head) {
        int cnt = 0;
        
        while (head) {
            cnt++;
            head = head->next;
        }

        return cnt;
    }
};

