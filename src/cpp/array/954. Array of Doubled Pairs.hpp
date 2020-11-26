#include <vector>
#include <map>
using namespace std;

class Solution_954 {
public:
    bool canReorderDoubled(vector<int>& A) {
        std::map<int,int> positive;
        std::map<int,int> negtive;
        int zero_cnt = 0;
        int len = A.size();
        for(int i = 0;i < len;i++){
            if (A[i] > 0){
                positive[A[i]]++;
            }else if(A[i] < 0){
                negtive[A[i]]++;
            }else{
                zero_cnt++;
            }
        }
        if (zero_cnt > 0){
            if ((zero_cnt | 1) == zero_cnt)
                return false;
        }
        for(auto it = positive.begin();it != positive.end();it++){
            if(it->second == 0)
                continue;
            auto it2 = positive.find(it->first * 2);
            if (it2 == positive.end())
                return false;
            if (it->second > it2->second)
                return false;
            it2->second -= it->second;
        }
        for(auto it = negtive.rbegin();it != negtive.rend();it++){
            if(it->second == 0)
                continue;
            auto it2 = negtive.find(it->first * 2);
            if (it2 == negtive.end())
                return false;
            if (it->second > it2->second)
                return false;
            it2->second -= it->second;
        }
        return true;
    }
};