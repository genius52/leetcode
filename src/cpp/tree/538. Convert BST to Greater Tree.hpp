#include "../define.h"

class Solution_538 {
public:
    int sum = 0;
    TreeNode* convertBST(TreeNode* root) {
        if(root == nullptr)
            return nullptr;
        convertBST(root->right);
        auto cur = root;
        sum += cur->val;
        cur->val = sum;
        convertBST(root->left);
        return cur;
    }
};