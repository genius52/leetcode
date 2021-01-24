#include "../define.h"

class Solution_572 {
public:
    bool isSubtree(TreeNode* s, TreeNode* t) {
        if(nullptr == t && nullptr == s)
            return true;
        if(nullptr == s || nullptr == t)
            return false;
        if(s->val != t->val){
            return isSubtree(s->left,t) || isSubtree(s->right,t);
        }
        else{
            return (isSubtree(s->left,t) || isSubtree(s->right,t)) || (isSubtree(s->left,t->left) && isSubtree(s->right,t->right));
        }
    }
};