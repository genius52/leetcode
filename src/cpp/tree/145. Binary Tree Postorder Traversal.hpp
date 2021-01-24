#include <vector>
#include <deque>
#include "../define.h"
using namespace std;

//
class Solution_145 {
public:
    void preorderTraversal(TreeNode* root){
        std::deque<TreeNode*> stack;
        auto visit = root;
        while(!stack.empty() || visit != nullptr){
            if (visit != nullptr){
                std::cout<<visit->val << std::endl;
                stack.push_back(visit);
                visit = visit->left;
            }else{
                visit = stack.back();
                stack.pop_back();
                if (visit->right != nullptr){
                    visit = visit->right;
                }
            }
        }
    }

    void inorderTraversal(TreeNode* root) {
        std::deque<TreeNode*> stack;
        auto visit = root;
        while(!stack.empty() || visit != nullptr){
            if(visit != nullptr){
                stack.push_back(visit);
                visit = visit.left;
            }
            else{
                visit = stack.back();
                std::cout<<visit->val<<std::endl;
                stack.pop_back();
                visit = visit->right;
            }
        }
    }

    vector<int> postorderTraversal(TreeNode* root) {
        std::vector<int> res;
        std::deque<TreeNode*> stack;
        auto visit = root;
        while(!stack.empty() || visit != nullptr){
            if(visit != nullptr){
                res.push_back(visit->val);
                stack.push_back(visit);
                visit = visit->right;
            }else{
                visit = stack.back();
                stack.pop_back();
                visit = visit->left;
            }
        }
        std::reverse(res.begin(),res.end());
        return res;
    }
};