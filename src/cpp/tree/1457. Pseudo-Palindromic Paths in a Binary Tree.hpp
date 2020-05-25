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
#include <deque>
#include <unordered_map>
using namespace std;

class Solution_1457 {
public:
    bool check_palindromic(std::deque<int>& trace){
        int len = trace.size();
        std::unordered_map<int,int> record;
        for(auto it = trace.begin();it != trace.end();it++){
            record[*it]++;
        }
        bool find_single = false;
        for(auto it = record.begin();it != record.end();it++){
            if((*it).second % 2 == 1){
                if(find_single)
                    return false;
                else{
                    find_single = true;
                }
            }
        }
        return true;
    }

    void dfs_pseudoPalindromicPaths(TreeNode* node,std::deque<int>& trace,int& total){
        if(nullptr == node)
            return;
        trace.push_back(node->val);
        if(nullptr == node->left && nullptr == node->right){
            if(check_palindromic(trace)){
                total++;
            }
            trace.pop_back();
            return;
        }
        dfs_pseudoPalindromicPaths(node->left,trace,total);
        dfs_pseudoPalindromicPaths(node->right,trace,total);
        trace.pop_back();
    }

    int pseudoPalindromicPaths (TreeNode* root) {
        int total = 0;
        std::deque<int> trace;
        dfs_pseudoPalindromicPaths(root,trace,total);
        return total;
    }
};