#include <vector>
#include <string>
#include <set>
using namespace std;

//s = "**|**|***|", queries = [[2,5],[5,9]]
class Solution_2055 {
public:
    vector<int> platesBetweenCandles(string s, vector<vector<int>>& queries) {
        std::set<int> candles;
        int len = s.size();
        std::vector<int> prefix_plates(len);
        if(s[0] == '*'){
            prefix_plates[0] = 1;
        }else if(s[0] == '|'){
            candles.insert(0);
        }
        for(int i = 1;i < len;i++){
            if(s[i] == '*'){
                prefix_plates[i] = 1 + prefix_plates[i - 1];
            }else if(s[i] == '|'){
                prefix_plates[i] = prefix_plates[i - 1];
                candles.insert(i);
            }
        }
        int q_len = queries.size();
        std::vector<int> res(q_len);
        for(int i = 0;i < q_len;i++){
            int start = queries[i][0];
            int end = queries[i][1];
            //大于等于start
            auto it = candles.lower_bound(start);
            if(it == candles.end() || *it >= end){
                continue;
            }
            //大于end
            auto it2 = candles.upper_bound(end);
            it2--;
            if(*it >= *it2){
                continue;
            }
            res[i] = prefix_plates[*it2] - prefix_plates[*it];
        }
        return res;
    }
};