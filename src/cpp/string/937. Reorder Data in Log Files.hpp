#include <vector>
using namespace std;

//Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]
//Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5 1","dig2 3 6"]
class Solution_937 {
public:
    bool check_letterlog(std::string s){
        int len = s.size();
        int pos = 0;
        while(pos < len && s[pos] != ' '){
            pos++;
        }
        return s[pos + 1] >= 'a' && s[pos + 1] <= 'z';
    }

    bool cmp_log_string(std::string& s1,std::string& s2){
        int len1 = s1.size();
        int len2 = s2.size();
        int pos1 = 0;
        int pos2 = 0;
        while(pos1 < len1 && s1[pos1] != ' '){
            pos1++;
        }
        while(pos2 < len2 && s2[pos2] != ' '){
            pos2++;
        }
        int ret = s1.substr(pos1,len1 - pos1).compare(s2.substr(pos2,len2 - pos2));
        if(ret == 0){
            int ret2 = s1.substr(0,pos1).compare(s2.substr(0,pos2));
            return ret2 > 0;
        }
        return ret > 0;
    }

    vector<string> reorderLogFiles(vector<string>& logs) {
        int len = logs.size();
        std::vector<string> res;
        res.resize(len);
        int digital_pos = len - 1;
        int letter_pos = 0;
        for(int i = len - 1;i >= 0;i--){
            if(!check_letterlog(logs[i]))
                res[digital_pos--] = logs[i];
            else
                res[letter_pos++] = logs[i];
        }
        for(int i = 0;i < letter_pos - 1;i++){
            for(int j = 0;j < letter_pos - 1 - i;j++){
                if(cmp_log_string(res[j],res[j+1]) > 0){
                    auto tmp = res[j + 1];
                    res[j + 1] = res[j];
                    res[j] = tmp;
                }
            }
        }
        return res;
    }
};