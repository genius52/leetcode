//
// Created by 陶诚程 on 2019/4/18.
//
#include "tree.h"
#include <iostream>

using namespace std;

void ReverseTree(tree_node* root)
{
    if(nullptr == root)
        return;
    auto tmp = root->left;
    root->left = root->right;
    root->right = tmp;
    ReverseTree(root->left);
    ReverseTree(root->right);
}


// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};

#include <math.h>


bool isSubtree(TreeNode* s, TreeNode* t) {
    if(nullptr == t)
        return true;
    if(nullptr == s)
        return false;
    while(nullptr != s && nullptr != t && s->val != t->val){
        if(s->val > t->val)
            s = s->left;
        else
            s = s->right;
    }
    if(nullptr == s)
        return false;
}

bool check_subtree(TreeNode* s,TreeNode* t){
    if(nullptr == t)
        return true;
    if(nullptr == s)
        return false;
    if(s->val != t->val)
        return false;
    return check_subtree(s->left,t->left) && check_subtree(s->right,t->right);
}



//Input: [3,2,1,6,0,5]

TreeNode* head = nullptr;

void recursive_choose_biggest(vector<int>& nums,int begin,int end,TreeNode*& parent){
    if(begin > end)
        return;
    int max = -pow(2,31);
    int i = begin;
    int max_index = begin;
    for(;i <= end;i++){
        if(nums[i] > max){
            max_index = i;
            max = nums[i];
        }
    }
    TreeNode* node = new TreeNode(nums[i]);
    if(nullptr == parent)
        parent = node;
    else{
        if(nullptr == parent->left)
            parent->left = node;
        else
            parent->right = node;
    }
    std::cout<<"begin="<<begin<<std::endl;
    std::cout<<"max_index="<<max_index<<std::endl;
    std::cout<<"end="<<end<<std::endl;
    recursive_choose_biggest(nums,begin,max_index-1,node);
    recursive_choose_biggest(nums,max_index+1,end,node);
}

TreeNode* constructMaximumBinaryTree(vector<int>& nums) {
    int left = 0;
    int right = nums.size() - 1;
    recursive_choose_biggest(nums,left,right,head);
    return head;
}


void preorder_visit(TreeNode* node,int val){
    if(nullptr == node)
        return;

    if(val < node->val){
        if(nullptr == node->left){
            TreeNode* new_node = new TreeNode(val);
            node->left = new_node;
        }
        else
            preorder_visit(node->left,val);
    }
    else{
        if(nullptr == node->right){
            TreeNode* new_node = new TreeNode(val);
            node->right = new_node;
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

vector<TreeNode*> allPossibleFBT(int N) {

}

TreeNode* Solution_1379::getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target){
    if (original == nullptr || target == nullptr){
        return nullptr;
    }
    std::deque<TreeNode*> q;
    q.push_back(original);
    std::deque<TreeNode*> q2;
    q2.push_back(cloned);
    while (!q.empty()){
        auto top = q.front();
        q.pop_front();
        auto clonedtop = q2.front();
        q2.pop_front();
        if (top == target){
            return clonedtop;
        }
        if (top->left != nullptr){
            q.push_back(top->left);
        }
        if (top->right != nullptr){
            q.push_back(top->right);
        }
        if (clonedtop->left != nullptr){
            q2.push_back(clonedtop->left);
        }
        if (clonedtop->right != nullptr){
            q2.push_back(clonedtop->right);
        }
    }
    return nullptr;
}

void preorder_visit(TreeNode* root){
    auto node = root;
    std::deque<TreeNode*> q;
    while(node != nullptr || !q.empty()){
        while (node != nullptr){
            std::cout<<node->val<<std::endl;
            q.push_back(node);
            node = node->left;
        }
        if(!q.empty()){
            node = q.back();
            q.pop_back();
            node = node->right;
        }
    }
}

void inorder_visit(TreeNode* root){
    std::deque<TreeNode*> q;
    auto node = root;
    while (node != nullptr || !q.empty()){
        while (node != nullptr){
            q.push_back(node);
            node = node->left;
        }
        if(!q.empty()){
            node = q.back();
            std::cout<<node->val<<std::endl;
            q.pop_back();
            node = node->right;
        }
    }
}