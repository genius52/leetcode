#include "../mylist.h"

class Solution_328 {
public:
    ListNode* oddEvenList(ListNode* head) {
        if(nullptr == head || nullptr == head->next)
            return head;
        ListNode* odd_node = head;
        ListNode* even_node = head->next;
        ListNode* last_odd = nullptr;
        ListNode* first_event = head->next;
        while(nullptr != even_node && nullptr != even_node->next)
        {
            odd_node->next = even_node->next;
            odd_node = odd_node->next;
            last_odd = odd_node;
            even_node->next = odd_node->next;
            even_node = even_node->next;
        }
        if(nullptr != even_node){
            if(nullptr != even_node->next)
                even_node->next = nullptr;
        }
        if(nullptr != last_odd)
            last_odd->next = first_event;
        return head;
    }
};