#include <vector>
#include <map>
#include <unordered_map>
using namespace std;

//Input: barcodes = [1,1,1,1,2,2,3,3]
//Output: [1,3,1,3,1,2,1,2]
class Solution_1054 {
public:
    vector<int> rearrangeBarcodes(vector<int>& barcodes){
        int len = barcodes.size();
        std::vector<int> res;
        res.resize(len);
        std::unordered_map<int,int> num_cnt;
        for(int i = 0;i < len;i++){
            if(num_cnt.find(barcodes[i]) == num_cnt.end()){
                num_cnt[barcodes[i]] = 1;
            }else{
                num_cnt[barcodes[i]]++;
            }
        }
        std::multimap<int, int, std::greater<int>> cnt_num;
        for(auto it : num_cnt){
            cnt_num.insert(std::make_pair(it.second, it.first));
        }
        int total = 0;
        int pos = 0;
        while(total < len){
            std::multimap<int,int>::iterator top = cnt_num.begin();
            int times = top->first;
            int num = top->second;
            cnt_num.erase(top);
            while(times > 0){
                while (res[pos] != 0){
                    pos++;
                    pos = pos % len;
                }
                res[pos] = num;
                pos += 2;
                pos = pos % len;
                times--;
                total++;
            }
        }
        return res;
    }
};