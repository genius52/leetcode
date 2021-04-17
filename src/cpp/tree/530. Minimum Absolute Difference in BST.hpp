#include "../define.h"

class Solution_530 {
    TreeNode* prev = nullptr;
public:
    int getMinimumDifference(TreeNode* root) {
        if(nullptr == root)
            return 2147483647;
        return inorder_visit(root);
    }

    int inorder_visit(TreeNode* node){
        if(nullptr == node)
            return 2147483647;

        int diff = inorder_visit(node->left);
        if(nullptr != prev)
            diff = min(node->val - prev->val, diff);
        if(diff == 0)
            return 0;
        prev = node;
        return min(inorder_visit(node->right),diff);
    }
};