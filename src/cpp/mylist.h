//
// Created by 陶诚程 on 2019/4/26.
//

#ifndef LEET_LIST_H
#define LEET_LIST_H

struct my_list{
    int val = 0;
    my_list* next = nullptr;
};

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(nullptr) {}
};


class Node {
public:
    int val;
    Node* next;
    Node* random;

    Node(int _val) {
        val = _val;
        next = nullptr;
        random = nullptr;
    }
};

void reverse_print_list(my_list* l);
void reverse_print_list_ex(my_list* l);

ListNode* oddEvenList(ListNode* head);
class Solution_138{
public:
    Node* copyRandomList(Node* head);
};
#endif //LEET_LIST_H