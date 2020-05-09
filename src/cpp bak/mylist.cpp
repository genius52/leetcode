//
// Created by 陶诚程 on 2019/4/26.
//
#include "mylist.h"
#include <iostream>
#include <map>

void reverse_print_list(my_list* l)
{
    if(nullptr == l)
        return;
    my_list* prev = nullptr;
    my_list* cur = l;
    while(nullptr != cur)
    {
        auto next = cur->next;
        cur->next = prev;
        prev = cur;
        cur = next;
    }

    while(prev != nullptr)
    {
        std::cout<<prev->val<<std::endl;
        prev = prev->next;
    }
}

void reverse_print_list_ex(my_list* l)
{
    while(l != nullptr)
    {
        reverse_print_list_ex(l->next);
        std::cout<<l->val<<std::endl;
        return;
        //l = l->next;
    }
}

ListNode* oddEvenList(ListNode* head) {
    if(nullptr == head || nullptr == head->next)
        return head;
    ListNode* odd_node = head;
    ListNode* even_node = head->next;
    ListNode* last_odd = nullptr;
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
    last_odd->next = head->next;
    return head;
}


//138
Node* Solution_138::copyRandomList(Node* head) {
    Node* v1 = head;
    Node* root = nullptr;
    //Node* v2 = nullptr;
    Node* pre = nullptr;
    std::map<Node*,Node*> dup_origin;
    while(v1){
        Node* v2 = new Node(v1->val);
        dup_origin.insert(std::make_pair(v2,v1));
        if (root == nullptr){
            root = v2;
        }else{
            pre->next = v2;
        }
        pre = v2;
        v1 = v1->next;
    }
    Node* v2 = root;
    while (v2){
        auto origin_node = dup_origin[v2];
        if (nullptr != origin_node->random){
            auto it = dup_origin.cbegin();
            for (;it != dup_origin.cend();it++) {
                if (origin_node->random == (*it).second){
                    break;
                }
            }
            v2->random = (*it).first;
        }
        v2 = v2->next;
    }
    return root;
}