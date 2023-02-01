#include <queue>
#include <vector>
using namespace std;

//[2,1,14,12]
//[11,7,13,6]
//3
class Solution_2542 {
public:
    long long maxScore(vector<int>& nums1, vector<int>& nums2, int k) {
        int len = nums1.size();
        std::vector<std::pair<int,int>> data(len);
        for(int i = 0;i < len;i++){
            data[i] = std::pair<int,int>{nums1[i],nums2[i]};
        }
        std::sort(data.begin(),data.end(),[](const std::pair<int,int>& d1,const std::pair<int,int>& d2)->bool{
            if(d1.second == d2.second){
                return d1.first > d2.first;
            }else{
                return d1.second > d2.second;
            }
        });
        long long sum = 0;
        int min_val = 1 << 31 - 1;
        std::priority_queue<std::pair<int,int>,std::vector<std::pair<int,int>>,std::greater<std::pair<int,int>>>  q;
        for(int i = 0;i < k;i++){
            if(data[i].second < min_val){
                min_val = data[i].second;
            }
            sum += data[i].first;
            q.push({data[i].first,data[i].second});
        }
        long long res = sum * min_val;
        for(int i = k;i < len;i++){
            auto top = q.top();
            q.pop();
            sum -= top.first;
            sum += data[i].first;
            long long cur = sum * data[i].second;
            if(cur > res)
                res = cur;
            q.push({data[i].first,data[i].second});
        }
        return res;
    }
};