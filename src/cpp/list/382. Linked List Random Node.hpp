
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
#include <random>
#include "../mylist.h"

class Solution_382 {
    ListNode* h;
public:
    /** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
    Solution_382(ListNode* head) {
        h = head;
    }

    /** Returns a random node's value. */
    int getRandom() {
        ListNode* node = h;
        ListNode* res = h;
        int cnt = 1;
        while(node != nullptr){
            if(rand() % cnt++ == 0){
                res = node;
            }
            node = node->next;
        }
        return res->val;
    }
};

/**
 * Your Solution object will be instantiated and called as such:
 * Solution* obj = new Solution(head);
 * int param_1 = obj->getRandom();
 */