#include <vector>
#include <queue>
using namespace std;

class Solution_786 {
public:
    void dfs_kthSmallestPrimeFraction(vector<int>& arr,int left,int right, int k,std::priority_queue<std::pair<float,std::pair<int,int>>, std::deque<std::pair<float,std::pair<int,int>>>, std::less<std::pair<float,std::pair<int,int>>>>& bigtop_queue){
        if(left > right)
            return;
        while(left <= right){
            if(!bigtop_queue.empty()){
                if(float(arr[left]/arr[right]) > bigtop_queue.top().first)
                    return;
            }
            for(int i = left;i < right;i++){
                float cur = float(arr[i])/float(arr[right]);
                if(bigtop_queue.size() < k){
                    bigtop_queue.push({cur,{arr[i],arr[right]}});
                }else{
                    if(cur > bigtop_queue.top().first)
                        break;
                    bigtop_queue.pop();
                    bigtop_queue.push({cur,{arr[i],arr[right]}});
                }
            }
            for(int i = right - 1;i > left;i--){
                float cur = float(arr[left])/float(arr[i]);
                if(bigtop_queue.size() < k){
                    bigtop_queue.push({cur,{arr[left],arr[i]}});
                }else{
                    if(cur > bigtop_queue.top().first)
                        break;
                    bigtop_queue.pop();
                    bigtop_queue.push({cur,{arr[left],arr[i]}});
                }
            }
            left++;
            right--;
        }
    }

    vector<int> kthSmallestPrimeFraction(vector<int>& arr, int k) {
        int len = arr.size();
        std::priority_queue<std::pair<float,std::pair<int,int>>, std::deque<std::pair<float,std::pair<int,int>>>, std::less<std::pair<float,std::pair<int,int>>> > bigtop_queue;
        int left = 0;
        int right = len - 1;
        dfs_kthSmallestPrimeFraction(arr,left,right,k,bigtop_queue);
        std::vector<int> res(2);
        res[0] = bigtop_queue.top().second.first;
        res[1] = bigtop_queue.top().second.second;
        return res;
    }
};