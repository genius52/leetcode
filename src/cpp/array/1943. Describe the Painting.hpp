#include <vector>
#include <map>
using namespace std;

class Solution_1943 {
public:
    vector<vector<long long>> splitPainting(vector<vector<int>>& segments) {
        int len = segments.size();
        std::vector<std::vector<long long>> res;
        std::map<long long,long long> record;
        for(int i = 0;i < len;i++){
            record[segments[i][0]] += segments[i][2];
            record[segments[i][1]] -= segments[i][2];
        }
//        long long pre = 0;
//        long long mix_color = 0;
//        for(auto it = record.begin();it != record.end();it++){
//            long long pos = it->first;
//            long long color = it->second;
//            if(mix_color != 0){
//                std::vector<long long> cur{pre,pos,mix_color};
//                res.push_back(cur);
//            }
//            mix_color += color;
//            pre = pos;
//        }
//
//        return res;

        long long start = -1;
        long long mix_color = 0;
        for(auto it = record.begin();it != record.end();it++){
            int pos = it->first;
            int color = it->second;
            if(start == -1){
                start = pos;
                mix_color += color;
            }else{
                std::vector<long long> cur{start,pos,mix_color};
                res.push_back(cur);
                mix_color += color;
                if(mix_color == 0){
                    start = -1;
                }else{
                    start = pos;
                }
            }
        }
        return res;
    }
};