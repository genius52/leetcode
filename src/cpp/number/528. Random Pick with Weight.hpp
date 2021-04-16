#include <vector>
#include <map>
using namespace std;

class Solution_528 {
    std::map<int,int> m_;
    int sum_ = 0;
public:
    Solution(vector<int>& w) {
        int len = w.size();
        for(int i = 0;i < len;i++){
            sum_ += w[i];
            m_[sum_] = i;
        }
    }

    int pickIndex() {
        int r = rand() % sum_;
        auto it = m_.upper_bound(r);
        return it->second;
    }
};