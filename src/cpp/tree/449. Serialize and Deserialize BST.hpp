/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
#include <string>
#include <iostream>
#include <vector>
#include <unordered_map>
#include <map>
#include "../tree.h"


class Codec_449 {
public:
    void level_visit(TreeNode* node,int parent_index,bool is_left,std::stringstream& s){
        if(nullptr == node)
            return;
        int index = is_left?(parent_index * 2):(parent_index * 2 + 1);
        s<<"|"<<index << ":"<<node->val;
        level_visit(node->left,index,true,s);
        level_visit(node->right,index,false,s);
    }
    // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        std::stringstream s;
        if(nullptr == root)
            return s.str();
        s << "1:"<<root->val;
        level_visit(root->left,1,true,s);
        level_visit(root->right,1,false,s);
        return s.str();
    }

    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        if(data.empty())
            return nullptr;
        std::vector<std::string> v;
        split(data,v,'|');
        std::map<int,int> graph;
        int m = 0;
        for(auto it = v.begin();it != v.end();it++){
            std::vector<std::string> node;
            split(*it,node,':');
            auto index = std::stoi(node[0],nullptr,10);
            auto val = std::stoi(node[1],nullptr,10);
            graph[index] = val;
            m = max(m,index);
        }
        std::map<int,TreeNode*> memo;
        for(long i = m;i >= 1;i--){
            if(graph.find(i) == graph.end())
                continue;
            TreeNode* node = new TreeNode(graph[i]);
            long left_child = i * 2;
            long right_child = i * 2 + 1;
            if(left_child <= m ){
                if(memo.find(left_child) != memo.end()){
                    node->left = memo[left_child];
                }
            }
            if(right_child <= m){
                if(memo.find(right_child) != memo.end()){
                    node->right = memo[right_child];
                }
            }
            memo[i] = node;
        }
        return memo[1];
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