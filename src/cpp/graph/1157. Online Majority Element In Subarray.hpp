#include <vector>
#include <set>
#include <map>
#include <unordered_map>
using namespace std;

class MajorityChecker {
    std::vector<std::pair<int,std::vector<int>>> record;
public:
    MajorityChecker(vector<int>& arr) {
        std::unordered_map<int, vector<int>> m;
        int len = arr.size();
        for(int i = 0;i < len;i++){
            m[arr[i]].push_back(i);
        }
        for(auto it : m){
            record.push_back({it.first,it.second});
        }
        std::sort(record.begin(),record.end(),[](std::pair<int,std::vector<int>>& p1,std::pair<int,std::vector<int>>& p2)->bool{
            return p1.second.size() > p2.second.size();
        });
    }

    int query(int left, int right, int threshold) {
        for(auto it = record.begin();it != record.end();it++){
            if(it->second.size() < threshold)
                continue;
            auto l = std::lower_bound(std::begin(it->second),std::end(it->second),left);//大于等于left
            auto r = std::upper_bound(std::begin(it->second),std::end(it->second),right);//大于right
            if(std::distance(l,r) >= threshold){
                return it->first;
            }
        }
        return -1;
    }
};

