#include <string>
#include <map>
#include <set>
#include <vector>
#include <deque>
#include <unordered_set>
#include <unordered_map>
using namespace std;

//Example 1:
//
//Input: s = "bcabc"
//Output: "abc"
//Example 2:
//
//Input: s = "cbacdcbc"
//Output: "acdb"
class Solution_316 {
public:
    string removeDuplicateLetters(string s){
        int len = s.size();
        int last_occur[26] = {0};
        bool visited[26] = {false};
        std::deque<char> q;
        for(int i = 0;i < len;i++){
            last_occur[s[i] - 'a'] = i;
        }
        for(int i = 0;i < len;i++){
            if(q.empty()){
                q.push_back(s[i]);
                visited[s[i] - 'a'] = true;
            }else{
                if(visited[s[i] - 'a'])
                    continue;
                //char prev_c = q.back();
                while(!q.empty() && s[i] <  q.back() && last_occur[q.back() - 'a'] > i){
                    visited[q.back() - 'a'] = false;
                    q.pop_back();
                }
                q.push_back(s[i]);
                visited[s[i] - 'a'] = true;
            }
        }
        std::string res;
        while (!q.empty()){
            res += q.front();
            q.pop_front();
        }
        return res;
    }

//    void dp_removeDuplicateLetters(string s,int len,int target_len,int pos,std::string prev,std::unordered_set<char>& dup_record,std::unordered_set<char>& choosed,std::vector<std::string>& res){
//        if (pos >= len){
//            if(target_len == 0)
//                res.push_back(prev);
//            return;
//        }
//        if(dup_record.find(s[pos]) == dup_record.end()){
//            prev += s[pos];
//            dp_removeDuplicateLetters(s,len,target_len - 1,pos + 1,prev,dup_record,choosed,res);
//        }else if(choosed.find(s[pos]) != choosed.end()){
//            dp_removeDuplicateLetters(s,len,target_len,pos + 1,prev,dup_record,choosed,res);
//        }else{
//            //auto positions = dup_record[s[pos]];
//            auto choose_cur = prev + s[pos];
//            std::unordered_set<char> copy_choosed = choosed;
//            copy_choosed.insert(s[pos]);
//            dp_removeDuplicateLetters(s,len,target_len - 1,pos + 1,choose_cur,dup_record,copy_choosed,res);
//            dp_removeDuplicateLetters(s,len,target_len,pos + 1,prev,dup_record,choosed,res);
//        }
//    }
//
//    string removeDuplicateLetters(string s) {
//        std::unordered_set<char> dup_record;
//        std::unordered_map<char,int> record;
//        int len = s.size();
//        int dup_count = 0;
//        for(int i = 0;i < len;i++){
//            if(record.find(s[i]) != record.end()){
//                dup_record.insert(s[i]);
//                dup_count++;
//            }else{
//                record[s[i]] = i;
//            }
//        }
//        std::vector<std::string> res;
//        std::unordered_set<char> choosed;
//        int target_len = len - dup_count;
//        dp_removeDuplicateLetters(s,len,target_len,0,"",dup_record,choosed,res);
//        std::partial_sort(res.begin(),res.begin(),res.end());
//        return res[0];
//    }
};