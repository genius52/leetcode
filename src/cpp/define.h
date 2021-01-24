//
// Created by 陶诚程 on 2021/1/24.
//

#ifndef CPP_DEFINE_H
#define CPP_DEFINE_H
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(nullptr) {}
};
#endif //CPP_DEFINE_H
