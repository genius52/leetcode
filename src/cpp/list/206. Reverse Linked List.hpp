#include "../define.h"

class Solution206 {
public:
    ListNode* reverseList(ListNode* node) {
        if(node == nullptr)
            return node;
        ListNode* new_head = nullptr;
        reverse(node,&new_head);
        node->next = nullptr;
        return new_head;
    }

    ListNode* reverse(ListNode* node,ListNode** new_head){
        if(node == nullptr){
            return nullptr;
        }
        auto next = reverse(node->next,new_head);
        if(next == nullptr){
            *new_head = node;
        }else{
            next->next = node;
        }
        return node;
    }
};