#include <vector>
#include <map>
using namespace std;

//0 <= i < j <= n - 1 且
//nums1[i] - nums1[j] <= nums2[i] - nums2[j] + diff.
//nums1[i] = nums2[i] <= nums2[j] - nums2[j] + diff
//record[i] <= record[j] + diff
class Solution_2426 {
public:
    long long numberOfPairs(vector<int>& nums1, vector<int>& nums2, int diff) {
        std::map<int,int> num_cnt;//左边比数字
        int l = nums1.size();
        std::vector<int> record(l);
        for(int i = 0;i < l;i++){
            record[i] = nums1[i] - nums2[i];
        }
        return recursive_numberOfPairs(record,0,l - 1,diff);
    }

    long long count(std::vector<int>& record,int left,int mid,int right,int diff){
        long long res = 0;
        int l = left;
        int r = mid + 1;
        while (l <= mid && r <= right){
            if(record[l] > record[r] + diff){
                r++;
            }else{
                res += right - r + 1;
                l++;
            }
        }
        return res;
    }

    long long recursive_numberOfPairs(std::vector<int>& record,int left,int right,int diff){
        if(left >= right)
            return 0;
        int mid = (left + right)/2;
        long long res = 0;
        res += recursive_numberOfPairs(record,left,mid,diff);
        res += recursive_numberOfPairs(record,mid + 1,right,diff);
        //calculate the number of record[left] <= record[right] + diff
        res += count(record,left,mid,right,diff);
        std::sort(record.begin() + left,record.begin() + right + 1);
        return res;
    }
};