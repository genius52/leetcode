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

TreeNode* constructMaximumBinaryTree(std::vector<int>& nums);

//297
class Codec {
public:
    void level_visit(TreeNode* node,int parent_num,bool is_left,std::stringstream& res){
        if(nullptr == node){
            return;
        }
        int cur_num = is_left?parent_num + 1:parent_num + 2;
        res << parent_num << "_" << cur_num  << "," << is_left << "," << node->val << "|";
        level_visit(node->left,cur_num,true,res);
        level_visit(node->right,cur_num,false,res);
    }
    // Encodes a tree to a single string.
    std::string serialize(TreeNode* root) {
        std::stringstream ss;
        level_visit(root,0,true,ss);
        return ss.str();
    }
//"0_1,3|1_2,7|1_3,9|3_4,2|3_5,1|"
    // Decodes your encoded data to tree.
    TreeNode* deserialize(std::string data) {
        if(data.empty()){
            return nullptr;
        }
        std::map<int,TreeNode*> record;
        std::vector<std::string> v;
        split(data,v,'|');
        for (auto i = 0;i < v.size();i++){
            std::vector<std::string> v2;
            split(v[i],v2,',');
            std::vector<std::string> v3;
            split(v2[0],v3,'_');

            TreeNode* t = new TreeNode(std::stoi(v2[2]));
            record[std::stoi(v3[1])] = t;
            auto it = record.find(std::stoi(v3[0]));
            if (it != record.end()){
                if(std::stoi(v2[1]) == 1){
                    it->second->left = t;
                }else{
                    it->second->right = t;
                }
            }
        }
        return record.begin()->second;
    }

private:
    void split(const std::string& s,std::vector<std::string>& sv,const char flag = ' ') {
        sv.clear();
        std::istringstream iss(s);
        std::string temp;

        while (getline(iss, temp, flag)) {
            sv.push_back(temp);
        }
        return;
    }
};

class Solution_1379 {
public:
    TreeNode* getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target);
};
#endif //LEET_TREE_H
