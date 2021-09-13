#include <string>
using namespace std;

class Solution_2002 {
public:
    bool check_maxProduct(std::string& s){
        int l = s.length();
        int left = 0;
        int right = l - 1;
        while(left < right){
            if (s[left] != s[right]){
                return false;
            }
            left++;
            right--;
        }
        return true;
    }

    int dfs_maxProduct(std::string& s,int l,int cur_pos,std::string& s1,std::string& s2){
        if(cur_pos >= l){
            if(check_maxProduct(s1) && check_maxProduct(s2))
                return s1.length() * s2.length();
            return 0;
        }
        s1.push_back(s[cur_pos]);
        int res1 = dfs_maxProduct(s,l,cur_pos + 1,s1,s2);
        s1.pop_back();

        s2.push_back(s[cur_pos]);
        int res2 = dfs_maxProduct(s,l,cur_pos + 1,s1,s2);
        s2.pop_back();

        int res3 = dfs_maxProduct(s,l,cur_pos + 1,s1,s2);
        return max(max(res1,res2),res3);
    }

    int maxProduct(string s) {
        int len = s.length();
        std::string s1;
        std::string s2;
        return dfs_maxProduct(s,len,0,s1,s2);
    }
};