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
#include "../define.h"

//Input: head = [1,2,3,4,5], k = 2
//Output: [2,1,4,3,5]
class Solution_25 {
public:
    ListNode* recursive_reverseKGroup(ListNode* node,int k){
        int step = 1;
        if(node == nullptr)
            return node;
        ListNode* begin = node;
        ListNode* end = node;
        while(end != nullptr && step < k){
            end = end->next;
            step++;
        }
        if (step != k || end == nullptr){
            return node;
        }
        auto ret = recursive_reverseKGroup(end->next,k);
        ListNode* pre;
        ListNode* next = begin->next;
        while(step > 1){
            pre = begin;
            begin = next;
            next = next->next;
            begin->next = pre;
            step--;
        }
        node->next = ret;
        return begin;
    }

    ListNode* reverseKGroup(ListNode* head, int k) {
        if(k == 1)
            return head;
        return recursive_reverseKGroup(head,k);
    }
};