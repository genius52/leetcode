
// Definition for a Node.
class Node {
public:
    int val;
    Node* prev;
    Node* next;
    Node* child;
};


class Solution_430 {
public:
    Node* dfs_flatten(Node* node){
        if(node == nullptr)
            return node;
        if(node->next == nullptr && node->child == nullptr)
            return node;
        if(node->child == nullptr){
            return dfs_flatten(node->next);
        }
        if(node->next == nullptr){
            node->next = node->child;
            node->next->prev = node;
            node->child = nullptr;
            return dfs_flatten(node->next);
        }
        auto old_next = node->next;
        node->next = node->child;
        node->next->prev = node;
        node->child = nullptr;
        auto child_last = dfs_flatten(node->next);
        child_last->next = old_next;
        old_next->prev = child_last;
        auto next_last = dfs_flatten(old_next);
        return next_last;
    }
    Node* flatten(Node* head) {
        dfs_flatten(head);
        return head;
    }
};