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
    void deleteNode(ListNode* node) {
        if (node == nullptr) {
            return;
        }
        
        ListNode* prev = node;
        ListNode* cur = node->next;

        while (cur) {
            prev->val = cur->val;
            if (cur->next == nullptr) {
                break;
            }
            prev = cur;
            cur = cur->next;
        }
        prev->next = nullptr;
    }

    void deleteNode2(ListNode* node) {
        node->val = node->next->val;
        ListNode* t = node->next;
        node->next = node->next->next;
        delete t;
    }
};