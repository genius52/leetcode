#include <vector>
using namespace std;

class Solution_350 {
public:
    vector<int> intersect(vector<int>& nums1, vector<int>& nums2) {
        std::map<int,int> map1;
        int len1 = nums1.size();
        int len2 = nums2.size();
        int record1[1001] = {0};
        for(int i = 0;i < len1;i++){
            record1[nums1[i]]++;
        }
        int record2[1001] = {0};
        for(int i = 0;i < len2;i++){
            if(record1[nums2[i]] > 0){
                record2[nums2[i]]++;
            }
        }
        std::vector<int> res;
        for(int i = 0;i <= 1000;i++){
            if(record2[i] == 0)
                continue;
            int cnt = min(record1[i],record2[i]);
            std::vector<int> add(cnt,i);
            res.insert(res.end(),add.begin(),add.end());
        }
        return res;
    }
};