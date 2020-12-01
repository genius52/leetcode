#include <vector>
#include <map>
#include <set>
#include "../tree.h"
using namespace std;

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
//Input: [3,9,20,null,null,15,7]
class Solution_987 {
private:
    void visit_verticalTraversal(TreeNode* node,int x_pos,int y_pos,std::map<int,std::map<int,std::vector<int>>>& record){
        if(nullptr == node)
            return;
        record[x_pos][y_pos].push_back(node->val);
        visit_verticalTraversal(node->left,x_pos - 1,y_pos + 1,record);
        visit_verticalTraversal(node->right,x_pos + 1,y_pos + 1,record);
    }
public:
    vector<vector<int>> verticalTraversal(TreeNode* root) {
        std::map<int,std::map<int,std::vector<int>>> record;
        std::vector<std::vector<int>> res;
        visit_verticalTraversal(root,0,0,record);
        for(auto r : record){
            std::vector<int> v;
            for(auto s :r.second){
                std::vector<int> same_pos;
                for(auto ele : s.second){
                    same_pos.push_back(ele);
                }
                std::sort(same_pos.begin(),same_pos.end());
                v.insert(v.end(), same_pos.begin(), same_pos.end());
            }
            res.push_back(v);
        }
        return res;
    }
};