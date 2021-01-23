#include "../tree.h"

class Solution_701 {
public:
    void insertIntoBST(TreeNode* node,int val){
        if(nullptr == node)
            return;

        if(val < node->val){
            if(nullptr == node->left){
                TreeNode* new_node = new TreeNode(val);
                node->left = new_node;
                return;
            }
            else
                preorder_visit(node->left,val);
        }
        else{
            if(nullptr == node->right){
                TreeNode* new_node = new TreeNode(val);
                node->right = new_node;
                return;
            }
            preorder_visit(node->right,val);
        }
    }

    TreeNode* insertIntoBST(TreeNode* root, int val) {
        if(nullptr == root)
            return root;
        preorder_visit(root,val);
        return root;
    }
};