#include "../mylist.h"

class Solution_138{
public:
    Node* copyRandomList(Node* head) {
        Node* v1 = head;
        Node* root = nullptr;
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
};