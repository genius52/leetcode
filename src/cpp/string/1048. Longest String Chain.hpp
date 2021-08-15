#include <string>
#include <vector>
#include <map>
#include <unordered_map>
using namespace std;

class Solution_1048 {
public:
    bool is_predecessor(std::string& s1,std::string& s2){
        int l1 = s1.length();
        int l2 = s2.length();
        int i = 0;
        int j = 0;
        int diff_cnt = 0;
        while(i < l1 && j < l2){
            if(s1[i] != s2[j]){
                diff_cnt++;
                if(diff_cnt >= 2){
                    return false;
                }
                j++;
            }else{
                i++;
                j++;
            }
        }
        return true;
    }

    int longestStrChain(vector<string>& words) {
        int len = words.size();
        std::map<int,std::vector<string>> record;//length:string
        std::unordered_map<string,int> string_cnt;//string:max chain length
        for(int i = 0;i < len;i++){
            record[words[i].size()].push_back(words[i]);
            string_cnt[words[i]] = 1;
        }
        int res = 1;
        for(auto it = record.begin();it != record.end();it++){
            auto cur_len = it->first;
            if(record.find(cur_len - 1) == record.end())
                continue;
            auto pre_strings = record[cur_len - 1];
            for(auto it2 = it->second.begin();it2 != it->second.end();it2++){
                for(auto it3 = pre_strings.begin();it3 != pre_strings.end();it3++){
                    if (is_predecessor(*it3,*it2)){
                        string_cnt[*it2] = max(string_cnt[*it2],string_cnt[*it3] + 1);
                        if(string_cnt[*it2] > res){
                            res = string_cnt[*it2];
                        }
                    }
                }
            }
        }
        return res;
    }
};