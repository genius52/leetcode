#include "../define.h"


class Solution_138{
public:
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

    Node* copyRandomList(Node* head) {
        Node* v1 = head;
        Node* root = nullptr;
        Node* pre = nullptr;
        std::unordered_map<Node*, Node*> origin_dup;
        while(v1){
            Node* v2 = new Node(v1->val);
            origin_dup.insert({v1,v2});
            if (root == nullptr){
                root = v2;
            }else{
                pre->next = v2;
            }
            pre = v2;
            v1 = v1->next;
        }
        v1 = head;
        Node* v2 = root;
        while (v1 != nullptr){
            v2->random = origin_dup[v1->random];
            v1 = v1->next;
            v2 = v2->next;
        }
        return root;
    }
};