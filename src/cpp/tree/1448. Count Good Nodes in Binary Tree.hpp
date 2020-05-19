/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
#include "../tree.h"

class Solution_1448 {
public:
    void inorder_goodNodes(TreeNode* node,int max_parent,int& cnt){
        if(nullptr == node)
            return;
        if(node->val >= max_parent){
            max_parent = node->val;
            cnt++;
        }
        inorder_goodNodes(node->left,max_parent,cnt);
        inorder_goodNodes(node->right,max_parent,cnt);
    }

    int goodNodes(TreeNode* root) {
        int cnt = 0;
        inorder_goodNodes(root,INT32_MIN,cnt);
        return cnt;
    }
};