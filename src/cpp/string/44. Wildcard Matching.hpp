#include <string>
#include <vector>
using namespace std;

//Given an input string (s) and a pattern (p),
// implement wildcard pattern matching with support for '?' and '*' where:
//'?' Matches any single character.
//'*' Matches any sequence of characters (including the empty sequence).
//Input: s = "adceb", p = "*a*b"
//Output: true
//Input: s = "acdcb", p = "a*c?b"
//Output: false
class Solution_44 {
public:
//    bool isMatch(string s, string p){
//        int len_s = s.length();
//        int len_p = p.length();
//        std::vector<std::vector<bool>> dp(len_s + 1,std::vector<bool>(len_p + 1));
//        for(int i = 1;i <= len_s;i++){
//            for(int j = 1;j <= len_p;j++){
//                if(p[j - 1] == '*'){
//                    dp[i][j] = dp[i - 1][j - 1] | dp[]
//                }
//                if(s[i - 1] == p[j - 1]){
//
//                }else{
//
//                }
//            }
//        }
//        return true;
//    }

    bool isMatch(string s, string p){
        int l1 = s.length();
        int l2 = p.length();
        int i,j = 0;
        int star_pos = -1;
        int s_stop = -1;
        while(i < l1 && j < l2){
            if(s[i] == p[j] || p[j] == '?'){
                i++;
                j++;
            }else{
                if(star_pos == -1 && p[j] != '*')
                    return false;
                if(p[j] == '*'){
                    s_stop = i;
                    star_pos = j;
                    j++;
                }
                if(star_pos != -1){
                    i++;
                    j++;
                }
            }
        }
        return true;
    }

//    bool isMatch(string s, string p) {
//        int l1 = s.length();
//        int l2 = p.length();
//        int i,j = 0;
//        int star_cnt = 0;
//        int unknown_cnt = 0;
//        while(i < l1 && j < l2){
//            while(j < l2 && (p[j] == '?' || p[j] == '*')){
//                if(p[j] == '?'){
//                    unknown_cnt++;
//                }else if(p[j] == '*'){
//                    star_cnt++;
//                }
//                j++;
//            }
//            if(j >= l2)
//                break;
//            char target = p[j];
//            int pre_i = i;
//            while(i < l1 && s[i] != target){
//                i++;
//            }
//            if(i >= l1)
//                break;
//            int skip_char_cnt = i - pre_i;
//            if(unknown_cnt > skip_char_cnt){//too many '?'
//                return false;
//            }
//            if(star_cnt == 0 && skip_char_cnt > unknown_cnt){// no enough '*' and '?'
//                return false;
//            }
//            star_cnt = 0;
//            unknown_cnt = 0;
//            i++;
//            j++;
//        }
//        if(i < l1){
//            int skip_char_cnt = l1 - i;
//            if(unknown_cnt > skip_char_cnt){//too many '?'
//                return false;
//            }
//            if(star_cnt == 0 && skip_char_cnt > unknown_cnt){// no enough '*' and '?'
//                return false;
//            }
//        }
//        while(j < l2){
//            if(p[j] != '*'){
//                return false;
//            }
//            j++;
//        }
//        return true;
//    }
};