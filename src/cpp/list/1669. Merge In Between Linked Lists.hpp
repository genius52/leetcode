#include "../mylist.h"
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

class Solution_1669 {
public:
    ListNode* mergeInBetween(ListNode* list1, int a, int b, ListNode* list2) {
        ListNode* l2end = list2;
        while(l2end->next != nullptr){
            l2end = l2end->next;
        }
        ListNode* l1_a = list1;
        ListNode* l1_b = list1;
        while(a > 1){
            l1_a = l1_a->next;
            l1_b = l1_b->next;
            a--;
            b--;
        }
        while(b >= 0){
            l1_b = l1_b->next;
            b--;
        }
        l1_a->next = list2;
        l2end->next = l1_b;
        return list1;
    }
};