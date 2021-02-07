#include <string>
using namespace std;

//Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//Output: true
//Example 2:
//
//Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//Output: false
class Solution_97 {
public:
    bool dfs_isInterleave(string& s1,int pos1,int group1,string& s2,int pos2,int group2,string& s3,int pos3,bool prev_choose_s1,std::vector<std::vector<bool>>& memo){
        int l1 = s1.length();
        int l2 = s2.length();
        int l3 = s3.length();
        if(pos1 >= l1 && pos2 >= l2 && pos3 >= l3){
            if(abs(group1 - group2) > 1)
                return false;
            return true;
        }
        //if(pos1 < l1 && pos2 < l2){
            if(memo[pos1][pos2]){
                return false;
            }
        //}
        if(s1[pos1] != s3[pos3] && s2[pos2] != s3[pos3]){
            memo[pos1][pos2] = true;
            return false;
        }
        if(s1[pos1] == s3[pos3]){
            if(prev_choose_s1){
                bool ret = dfs_isInterleave(s1,pos1 + 1,group1,s2,pos2,group2,s3,pos3 + 1,true,memo);
                if(ret)
                    return true;
            }else{
                bool ret = dfs_isInterleave(s1,pos1 + 1,group1 + 1,s2,pos2,group2,s3,pos3 + 1,true,memo);
                if(ret)
                    return true;
            }
        }
        if(s2[pos2] == s3[pos3]){
            if(prev_choose_s1){
                bool ret = dfs_isInterleave(s1,pos1,group1,s2,pos2 + 1,group2 + 1,s3,pos3 + 1,false,memo);
                if(ret)
                    return true;
            }else{
                bool ret = dfs_isInterleave(s1,pos1,group1,s2,pos2 + 1,group2,s3,pos3 + 1,false,memo);
                if(ret)
                    return true;
            }
        }
        memo[pos1][pos2] = true;
        return false;
    }

    bool isInterleave(string s1, string s2, string s3) {
        if(s1.empty()){
            return s2 == s3;
        }
        if(s2.empty()){
            return s1 == s3;
        }
        int l1 = s1.length();
        int l2 = s2.length();
        int l3 = s3.length();
        if(l1 + l2 != l3)
            return false;
        std::vector<std::vector<bool>> memo(l1 + 1,std::vector<bool>(l2 + 1));
        return dfs_isInterleave(s1,0,0,s2,0,0,s3,0,true,memo) || dfs_isInterleave(s1,0,0,s2,0,0,s3,0,false,memo);
    }
};