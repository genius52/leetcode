#include <vector>
#include <string>
using namespace std;

class Solution_1773 {
public:
    //ruleKey == "type" and ruleValue == typei.
    //ruleKey == "color" and ruleValue == colori.
    //    ruleKey == "name" and ruleValue == namei.
    int countMatches(vector<vector<string>>& items, string ruleKey, string ruleValue) {
        int len = items.size();
        int res = 0;
        for (int i = 0; i < len; i++) {
            if (ruleKey == "type" && items[i][0] == ruleValue) {
                res++;
            }
            else if (ruleKey == "color" && items[i][1] == ruleValue) {
                res++;
            }
            else if (ruleKey == "name" && items[i][2] == ruleValue) {
                res++;
            }
        }
        return res;
    }
};