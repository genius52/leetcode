#include <math>
#include "./define.h"

class Solution {
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        if(nullptr == headA || nullptr == headB)
            return nullptr;
        int len_a = 0;
        auto node = headA;
        while(nullptr != node){
            len_a++;
            node = node->next;
        }
        int len_b = 0;
        node = headB;
        while(nullptr != node){
            len_b++;
            node = node->next;
        }
        int diff = abs(len_a - len_b);
        if(len_a > len_b){
            while(diff-- > 0){
                headA = headA->next;
            }
        }
        else{
            while(diff-- > 0){
                headB = headB->next;
            }
        }
        while(headA != nullptr && headB!= nullptr && headA != headB){
            headA = headA->next;
            headB = headB->next;
        }
        if(headA == headB)
            return headA;
        return nullptr;
    }
};