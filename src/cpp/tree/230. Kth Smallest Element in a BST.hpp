#include "../define.h"

class Solution {
    int cnt = 0;
public:
    int kthSmallest(TreeNode* root, int k) {
        int value;
        if(inorder_visit(root,k,value))
            return value;
        return -1;
    }

    bool inorder_visit(TreeNode* node,int k,int& value){
        if(nullptr == node)
            return false;
        if(inorder_visit(node->left,k,value))
            return true;
        cnt++;
        if(cnt == k){
            value = node->val;
            return true;
        }
        if(inorder_visit(node->right,k,value))
            return true;
        return false;
    }
};