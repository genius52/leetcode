#include <vector>
#include <map>
using namespace std;

class Solution_846 {
public:
    bool isNStraightHand(vector<int>& hand, int W) {
        int len = hand.size();
        if (len % W != 0 || len < W)
            return false;
        std::map<int,int> record;
        for(auto n: hand){
            record[n]++;
        }
        int groups = len / W;
        while (groups > 0){
            int start = record.begin()->first;
            for (int i = 0;i < W;i++){
                if(record.find(start) == record.end())
                    return false;
                record[start]--;
                if (record[start] == 0)
                    record.erase(start);
                start++;
            }
            groups--;
        }
        if(record.size() != 0)
            return false;
        return true;
    }
};