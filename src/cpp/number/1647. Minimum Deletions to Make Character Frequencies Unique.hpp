#include <string>
#include <map>
#include <unordered_map>
#include <set>
using namespace std;

class Solution_1647 {
public:
    int minDeletions(string s) {
        int chars[26]={0};
        for(auto c : s){
            chars[c - 'a']++;
        }
        int res = 0;
        std::set<int> used;
        int start = 1;
        for (int i = 0;i < 26;i++){
            if(chars[i] == 0)
                continue;
            if(used.find(chars[i]) == used.end()){
                used.insert(chars[i]);
            }else{
                auto count = chars[i] - 1;
                bool find = false;
                while (count >= 1){
                    if(used.find(count) == used.end()){
                        find = true;
                        break;
                    }
                    count--;
                }
                if(find){
                    res += chars[i] - count;
                    used.insert(count);
                }else{
                    res += chars[i];
                }
            }
        }
        return res;
    }
};