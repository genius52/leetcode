#include <vector>
#include <set>
using namespace std;

//Input: arr1 = [1,5,3,6,7], arr2 = [1,3,2,4]
//Output: 1
//Explanation: Replace 5 with 2, then arr1 = [1, 2, 3, 6, 7].
class Solution_1187 {
public:
    int dfs_makeArrayIncreasing(vector<int>& arr1,int len1,int pos1,int prev,vector<int>& arr2,int len2,int pos2,std::vector<std::vector<int>>& memo){
        if(pos1 >= len1){
            return 0;
        }
        if(pos2 >= len2)
            return -1;
        int old_pos2 = pos2;
        if(memo[pos1][old_pos2] != -2)
            return memo[pos1][old_pos2];
        int res = -1;
        if(arr1[pos1] <= prev){
            std::vector<int>::iterator it;
            int j = pos2 - 1;
            if(j < 0)
                j = 0;
            if(prev >= arr2[pos2]){
                it = std::upper_bound(arr2.begin() + j,arr2.end(),prev);
            }else{
                it = arr2.begin() + pos2;
            }
            if(it == arr2.end()) {
                return -1;
            }
            pos2 = std::distance(arr2.begin(),it);
            res = dfs_makeArrayIncreasing(arr1,len1,pos1 + 1,*it,arr2,len2,pos2,memo);
            if(res != -1)
                res++;
            memo[pos1][old_pos2] = res;
            return res;
        }else{
            int res1 = dfs_makeArrayIncreasing(arr1,len1,pos1 + 1,arr1[pos1],arr2,len2,pos2,memo);
            int res2 = -1;
            std::vector<int>::iterator it;
            int j = pos2 - 1;
            if(j < 0)
                j = 0;
            if(prev >= arr2[pos2]){
                it = std::upper_bound(arr2.begin() + j,arr2.end(),prev);
            }else{
                it = arr2.begin() + pos2;
            }
            if(it == arr2.end()) {
                res2 = -1;
            }else{
                if(*it < arr1[pos1]){
                    pos2 = std::distance(arr2.begin(),it);
                    res2 = dfs_makeArrayIncreasing(arr1,len1,pos1 + 1,*it,arr2,len2,pos2,memo);
                    if(res2 != -1)
                        res2++;
                }
            }
            if(res1 == -1)
                res = res2;
            else if(res2 == -1)
                res = res1;
            else
                res = min(res1,res2);
            memo[pos1][old_pos2] = res;
            return res;
        }
    }

    int makeArrayIncreasing(vector<int>& arr1, vector<int>& arr2) {
        int len1 = arr1.size();
        std::sort(arr2.begin(),arr2.end());
        int len2 = arr2.size();
        std::vector<std::vector<int>> memo(len1,std::vector<int>(len2,-2));
        return dfs_makeArrayIncreasing(arr1,len1,0,-2147483648,arr2,len2,0,memo);
    }
};