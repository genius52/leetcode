#include "../define.h"

class Solution237 {
public:
    void deleteNode(ListNode* node) {
        auto next = node->next;
        node->val = next->val;
        if(next != nullptr)
            node->next = next->next;
        else{
            node->next = nullptr;
        }
    }
};