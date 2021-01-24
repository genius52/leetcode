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
#include "../define.h"

class Solution_450 {
    int find_max_inleft(TreeNode* node){
        while(node->right != nullptr){
            node = node->right;
        }
        return node->val;
    }
public:
    TreeNode* deleteNode(TreeNode* root, int key) {
        if(nullptr == root)
            return root;
        if(root->val < key){
            root->right = deleteNode(root->right,key);
        }else if(root->val > key){
            root->left = deleteNode(root->left,key);
        }
        else if(root->val == key){
            if(root->left == nullptr)
                return root->right;
            if(root->right == nullptr)
                return root->left;
            auto max_inleft = find_max_inleft(root->left);
            root->val = max_inleft;
            root->left = deleteNode(root->left,max_inleft);
        }
        return root;
    }
};