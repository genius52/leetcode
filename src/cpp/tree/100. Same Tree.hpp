#include "../define.h"

class Solution_100 {
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        if(nullptr == p && nullptr == q)
            return true;
        if(nullptr == p && nullptr != q)
            return false;
        if(nullptr != p && nullptr == q)
            return false;
        if(p->val != q->val)
            return false;
        return isSameTree(p->left,q->left) && isSameTree(p->right,q->right);
    }
};