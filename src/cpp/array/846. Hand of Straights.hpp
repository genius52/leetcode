#include <vector>
#include <map>
#include <set>
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

    bool isNStraightHand2(vector<int>& hand, int W){
        std::multiset<int> s;
        for(auto n: hand){
            s.insert(n);
        }
        int len = s.size();
        if(len % W != 0 || len < W){
            return false;
        }
        while(!s.empty()){
            int first = *s.begin();
            for(int i = first;i < first + W;i++){
                auto it = s.find(i);
                if(it == s.end())
                    return false;
                else{
                    s.erase(it);
                }
            }
        }
        return true;
    }
};