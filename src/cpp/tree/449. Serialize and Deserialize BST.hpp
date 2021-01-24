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
#include "../define.h"


//class Codec_449 {
//public:
//    void level_visit(TreeNode* node,int parent_index,bool is_left,std::stringstream& s){
//        if(nullptr == node)
//            return;
//        int index = is_left?(parent_index * 2):(parent_index * 2 + 1);
//        s<<"|"<<index << ":"<<node->val;
//        level_visit(node->left,index,true,s);
//        level_visit(node->right,index,false,s);
//    }
//    // Encodes a tree to a single string.
//    string serialize(TreeNode* root) {
//        std::stringstream s;
//        if(nullptr == root)
//            return s.str();
//        s << "1:"<<root->val;
//        level_visit(root->left,1,true,s);
//        level_visit(root->right,1,false,s);
//        return s.str();
//    }
//
//    // Decodes your encoded data to tree.
//    TreeNode* deserialize(string data) {
//        if(data.empty())
//            return nullptr;
//        std::vector<std::string> v;
//        split(data,v,'|');
//        std::map<int,int> graph;
//        int m = 0;
//        for(auto it = v.begin();it != v.end();it++){
//            std::vector<std::string> node;
//            split(*it,node,':');
//            auto index = std::stoi(node[0],nullptr,10);
//            auto val = std::stoi(node[1],nullptr,10);
//            graph[index] = val;
//            m = max(m,index);
//        }
//        std::map<int,TreeNode*> memo;
//        for(long i = m;i >= 1;i--){
//            if(graph.find(i) == graph.end())
//                continue;
//            TreeNode* node = new TreeNode(graph[i]);
//            long left_child = i * 2;
//            long right_child = i * 2 + 1;
//            if(left_child <= m ){
//                if(memo.find(left_child) != memo.end()){
//                    node->left = memo[left_child];
//                }
//            }
//            if(right_child <= m){
//                if(memo.find(right_child) != memo.end()){
//                    node->right = memo[right_child];
//                }
//            }
//            memo[i] = node;
//        }
//        return memo[1];
//    }
//private:
//    void split(const std::string& s,std::vector<std::string>& sv,const char flag = ' ') {
//        sv.clear();
//        std::istringstream iss(s);
//        std::string temp;
//
//        while (getline(iss, temp, flag)) {
//            sv.push_back(temp);
//        }
//        return;
//    }
//};

class Codec_449 {
public:
    // Encodes a tree to a single string.
    void my_serialize(TreeNode* node,std::stringstream& ss){
        if(node == nullptr)
            return;
        static bool start = true;
        if(start){
            start = false;
            ss<<node->val;
        }else{
            ss<<","<<node->val;
        }
        my_serialize(node->left,ss);
        my_serialize(node->right,ss);
    }

    string serialize(TreeNode* root) {
        std::stringstream s;
        if(nullptr == root)
            return s.str();
        my_serialize(root,s);
        return s.str();
    }

    TreeNode* my_deserialize(std::vector<int>& v,int len,int pos,long low,long high){
        if(pos >= len)
            return nullptr;

        if(v[pos] < low || v[pos] > high)
            return nullptr;

        TreeNode* node = new TreeNode(v[pos]);
        int edge = pos + 1;
        while(edge < len){
            if(v[edge] > v[pos])
                break;
            edge++;
        }
        node->left = my_deserialize(v,len,pos + 1,low,v[pos]);
        node->right = my_deserialize(v,len,edge,v[pos],high);
        return node;
    }
    //5,3,2,4,6,7
    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        std::vector<int> v;
        split(data,v,',');
        int len = v.size();
        long low = 1 << 31;
        TreeNode* root = my_deserialize(v,len,0,low,0x7fffffff);
        return root;
    }

    void split(const std::string& s,std::vector<int>& sv,const char flag = ' ') {
        sv.clear();
        std::istringstream iss(s);
        std::string temp;

        while (getline(iss, temp, flag)) {
            try{
                auto val = std::stoi(temp,nullptr,10);
                sv.push_back(val);
            }
            catch (...){
                continue;
            }
        }
        return;
    }
};

