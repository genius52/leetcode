#include "../define.h"
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

class Solution_21 {
public:
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        if(nullptr == l1){
            return l2;
        }
        if(nullptr == l2){
            return l1;
        }
        if(l1->val < l2->val){
            auto ret = mergeTwoLists(l1->next,l2);
            l1->next = ret;
            return l1;
        }else{
            auto ret = mergeTwoLists(l1,l2->next);
            l2->next = ret;
            return l2;
        }
    }
};