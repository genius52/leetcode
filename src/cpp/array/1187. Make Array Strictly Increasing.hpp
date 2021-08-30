#include <vector>
#include <set>
using namespace std;

//Input: arr1 = [1,5,3,6,7], arr2 = [1,3,2,4]
//Output: 1
//Explanation: Replace 5 with 2, then arr1 = [1, 2, 3, 6, 7].
class Solution_1187 {
public:
    int dfs_makeArrayIncreasing(vector<int>& arr1,int len,int pos,std::multiset<int>& s){
        if(pos >= len){
            return 0;
        }
        var res int = -1;
        if(arr1[pos] <= arr1[pos - 1]){
            //change arr1[pos - 1],choose smallest
            if(s.empty())
                return  -1;
            auto it = s.begin();
            if(*it >= arr1[pos])
                return -1;
            int tmp = *it;
            s.erase(it);
            int res1 = -1;
            int cur1 = dfs_makeArrayIncreasing(arr1,len,pos + 1,s);
            if(cur1 != -1){
                res1 = cur1 + 1;
            }
            s.insert(tmp);
            //change arr1[pos]
            auto it2 = s.upper_bound(arr1[pos]);
            int res2 = -1;
            if(it2 != s.end()){
                int tmp2 = *it2;
                int old2 = arr1[pos];
                arr1[pos] = tmp2;
                s.erase(it2);
                int cur = dfs_makeArrayIncreasing(arr1,len,pos + 1,s);
                if(cur != -1){
                    res2 = cur + 1;
                }
                arr1[pos] = old2;
                s.insert(tmp2);
            }
            if(res1 == -1){
                res = res2;
            }else if(res2 == -1){
                res = res1;
            }else{
                res = min(res1,res2);
            }
        }else{
            res = dfs_makeArrayIncreasing(arr1,len,pos + 1,s);
        }
        return res;
    }

    int makeArrayIncreasing(vector<int>& arr1, vector<int>& arr2) {
        int len1 = arr1.size();
        int len2 = arr2.size();
        std::multiset<int> s;
        for(int i = 0;i < len2;i++)
            s.insert(arr2[i]);
        return dfs_makeArrayIncreasing(arr1,len1,1,s);
    }
};