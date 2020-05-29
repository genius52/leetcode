#include <vector>
#include <queue>
using namespace std;

//Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
//Output: [[1,2],[1,4],[1,6]]

class Solution_373 {
    struct cmp{
        bool operator()(std::pair<int,int> a, std::pair<int,int> b) {
            if ((a.first + a.second) < (b.first + b.second))
                return false;
            return true;
        }
    };
public:
    vector<vector<int>> kSmallestPairs(vector<int>& nums1, vector<int>& nums2, int k) {
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,cmp > q;
        for(int i = 0;i < nums1.size();i++){
            for(int j = 0;j < nums2.size();j++){
                std::pair<int,int> p;
                p.first = nums1[i];
                p.second = nums2[j];
                q.push(p);
            }
        }
        std::vector<std::vector<int>> res;
        for(int i = 0;i < k && !q.empty();i++){
            res.push_back({q.top().first,q.top().second});
            q.pop();
        }
        return res;
    }
};