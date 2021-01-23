#include "../tree.h"
#include <vector>
using namespace std;

class Solution_654{
public:
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
};
