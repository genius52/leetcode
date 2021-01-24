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
#include <string>
#include <set>
#include <sstream>
using namespace std;

class Solution_988 {
    void dfs_smallestFromLeaf(TreeNode* node,std::vector<int>& prev,std::set<std::string>& record){
        if (node == nullptr)
            return;
        std::vector<int> v = prev;
        v.push_back(node->val);
        if(nullptr == node->left && nullptr == node->right){
            std::basic_ostringstream<char> ss;
            for(auto it = v.rbegin();it != v.rend();it++){
                char c =  'a' + *it;
                ss << c;
            }
            record.insert(ss.str());
        }else{
            dfs_smallestFromLeaf(node->left,v,record);
            dfs_smallestFromLeaf(node->right,v,record);
        }
    }
public:
    string smallestFromLeaf(TreeNode* root) {
        std::set<std::string> record;
        std::vector<int> v;
        dfs_smallestFromLeaf(root,v,record);
        if(record.size() == 0)
            return "";
        return *record.begin();
    }
};