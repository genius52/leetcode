#include <string>
#include <queue>
#include <map>
#include <set>
using namespace std;

class Solution_1505 {
public:
    //记录下所有的0-9的位置然后用这些预处理的位置来进行查找
    string minInteger2(string num, int k){
        return "";
    }

    //Brute force by select sort
    string minInteger(string num, int k) {
        int len = num.size();
        std::vector<char> record(len);
        for(int i = 0;i < len;i++){
            record[i] = num[i];
        }
        for(int i = 0;i < len;i++){
            if(record[i] == '0')
                continue;
            int min_idx = i;
            for (int j = i + 1; j < len && j - i <= k; ++j) {
                if (record[j] < record[min_idx]){
                    min_idx = j;
                    if(record[j] == '0')
                        break;
                }
            }
            if(min_idx != i){
                k -= min_idx - i;
                auto tmp = record[min_idx];
                record.erase(record.begin() + min_idx);
                record.insert(record.begin() + i,tmp);
                if (k == 0)
                    break;
            }
        }
        std::string res;
        res.resize(len);
        for (int i = 0; i < len; ++i) {
            res[i] = record[i];
        }
        return res;
    }
};