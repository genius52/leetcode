#include <vector>
#include <bitset>
//#include <unordered_set>
using namespace std;

class Solution_1239 {
public:
    int maxLength(vector<string>& arr) {
        std::vector<std::bitset<26>> total;
        int max_len = 0;
        int len = arr.size();
        for(int i = 0;i < len;i++){
            std::bitset<26> cur;
            int cur_len = arr[i].length();
            for(int j = 0;j < cur_len;j++){
                cur[arr[i][j] - 'a'] = 1;
            }
            if(cur.count() != cur_len)//count函数用来求bitset中1的位数
                continue;
            if(cur_len > max_len)
                max_len = cur_len;
            bool conflict = false;
            for(int j = 0;j < total.size();j++){
                std::bitset<26> pre = total[j];
                for(int i = 0;i < 26;i++){
                    if(cur[i] == 1 && pre[i] == 1){
                        conflict = true;
                        break;
                    }
                }
                if(!conflict){
                    auto next = pre|cur;
                    if(next.count() > max_len)
                        max_len = next.count();
                    total.push_back(next);
                }
                conflict = false;
            }
            total.push_back(cur);
        }
        return max_len;
    }
};