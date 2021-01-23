//
// Created by 陶诚程 on 2019/5/31.
//

#ifndef TREE_H
#define TREE_H
#include <vector>
#include <map>
#include <queue>
#include <string>
#include <sstream>

struct tree_node
{
    int val;
    tree_node* left = nullptr;
    tree_node* right = nullptr;
};

struct TreeNode {
     int val;
     TreeNode *left;
     TreeNode *right;
     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 };

#endif //LEET_TREE_H
