#include <vector>
#include <map>
#include <set>
using namespace std;

class Solution_2122 {
public:
    std::vector<int> check_recoverArray(std::multiset<int>& record,int k,int l){
        std::vector<int> res;
        int cnt = 0;
        while (cnt < l / 2) {
            auto it = record.begin();
            int small = *it;
            int big = *it + k;
            if(record.find(big) == record.end())
                return std::vector<int>{};
            record.erase(record.find(big));
            record.erase(record.find(small));
            res.push_back(small + k / 2);
            cnt++;
        }
        return res;
    }

    vector<int> recoverArray(vector<int>& nums) {
        int len = nums.size();
        if(len == 2){
            int diff = nums[1] - nums[0];
            if((diff | 1) != diff){
                return std::vector<int>{(nums[0] + nums[1])/2};
            }
        }
        std::multiset<int> record;
        std::sort(nums.begin(),nums.end());
        for(int i = 1;i < len;i++){
            record.insert(nums[i]);
        }
        for (int i = 1; i < len; ++i) {
            int diff = nums[i] - nums[0];
            if((diff | 1) == diff || diff == 0){
                continue;
            }
            std::multiset<int> copy = record;
            copy.erase(copy.find(nums[i]));
            auto res = check_recoverArray(copy,diff,len - 2);
            if(res.size() > 0){
                res.push_back(nums[0] + diff/2);
                return res;
            }
        }
        return std::vector<int>{};
    }
};

//	sort.Ints(nums)
//	var l int = len(nums)
//	var res []int
//	var record map[int]int = make(map[int]int)
//	for i := 1; i < l; i++ {
//		record[nums[i]]++
//	}
//	for i := 1; i <= l/2; i++ {
//		diff := nums[i] - nums[0]
//		if diff|1 == diff {
//			continue
//		}
//		record[nums[i]]--
//		res = check_recoverArray(record, diff)
//		if len(res) > 0 {
//			return res
//		}
//		record[nums[i]]++
//	}
//	return res