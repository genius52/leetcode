#include <string>
#include <unordered_map>
#include <queue>
using namespace std;

class Solution_767 {
public:
    string reorganizeString(string s) {
        std::priority_queue<std::pair<int,char>> q;
        std::unordered_map<char,int> char_cnt;
        int len = s.size();
        int max_cnt = 0;
        for(int i = 0;i < len;i++){
            char_cnt[s[i]]++;
            if(char_cnt[s[i]] > max_cnt){
                max_cnt = char_cnt[s[i]];
            }
        }
        if(((len | 1) == len && max_cnt > len/2 + 1) || ((len | 1) != len && max_cnt > len/2)) {
            return "";
        }
        for(auto it = char_cnt.begin();it != char_cnt.end();it++){
            std::pair<int,char> p;
            p.first = (*it).second;
            p.second = (*it).first;
            q.push(p);
        }
        std::vector<char> data(len);
        int index = 0;
        while(!q.empty()){
            auto p = q.top();
            q.pop();
            char c = p.second;
            int cnt = p.first;
            while(cnt > 0){
                if(data[index % len] == 0){
                    data[index] = c;
                    index = (index + 2) % len;
                    cnt--;
                }else{
                    index = (index + 1) % len;
                }
            }
        }
        std::string res;
        for(int i = 0;i < len;i++){
            res.push_back(data[i]);
        }
        return res;
    }
};