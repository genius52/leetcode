
//Input: n = 3, k = 1
//Output: 2
//Explanation:
//The array [1,3,2] and [2,1,3] have exactly 1 inverse pair.
class Solution_629 {
public:
    bool check_valid(std::vector<int>& nums,int len,int k){
        int total = 0;
        for(int i = 0;i < len;i++){
            for(int j = i + 1;j < len;j++){
                if(nums[i] > nums[j])
                    total++;
                if(total > k)
                    return false;
            }
        }
        return total == k;
    }

    void perm(std::vector<int>& nums,int len,int pos,int k,int& res){
        if(pos >= len){
            if(check_valid(nums,len,k))
                res++;
            return;
        }
        for(int i = pos;i < len;i++){
            int tmp = nums[pos];
            nums[pos] = nums[i];
            nums[i] = tmp;
            perm(nums,len,pos + 1,k,res);
            tmp = nums[pos];
            nums[pos] = nums[i];
            nums[i] = tmp;
        }
    }

    int kInversePairs(int n, int k) {
        if(k == 0){
            return 1;
        }
        std::vector<int> nums(n);
        for(int i = 0;i < n;i++){
            nums[i] = i + 1;
        }
        int res = 0;
        perm(nums,n,0,k,res);
        return res;
    }
};