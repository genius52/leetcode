#include <vector>
#include <string>
using namespace std;

class Solution_748 {
public:
    string shortestCompletingWord(string licensePlate, vector<string>& words) {
        int char_cnt = 0;
        std::vector<int> record;
        record.resize(26);
        for(auto s : licensePlate){
            if((s >= 'a' && s <= 'z')){
                record[s - 'a']++;
                char_cnt++;
            }
            if((s >= 'A' && s <= 'Z')){
                record[s - 'A']++;
                char_cnt++;
            }
        }
        std::string res;
        int len = words.size();
        int min_length = INT32_MAX;
        for(int i = 0;i < len;i++){
            int l = words[i].length();
            if(l < char_cnt)
                continue;
            std::vector<int> copy_record = record;
            int cnt = 0;
            for(int j = 0;j < l;j++){
                if(copy_record[words[i][j] -  'a'] != 0){
                    cnt++;
                    copy_record[words[i][j] - 'a']--;
                }
            }
            if(cnt == char_cnt){
                if(l < min_length){
                    min_length = l;
                    res = words[i];
                }
            }
        }
        return res;
    }
};