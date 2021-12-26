#include <vector>
#include <set>
using namespace std;

//rooms[i] = [roomIdi, sizei]
//queries[j] = [preferredj, minSizej]
class Solution_1847 {
public:
    vector<int> closestRoom(vector<vector<int>>& rooms, vector<vector<int>>& queries) {
        int q_len = queries.size();
        std::vector<int> res(q_len);
        for(int i = 0;i < q_len;i++){
            queries[i].push_back(i);
        }
        std::sort(queries.begin(),queries.end(),[](const std::vector<int>& v1,const std::vector<int>& v2){
            return v1[1] > v2[1];
        });
        int room_len = rooms.size();
        std::sort(rooms.begin(),rooms.end(),[](const std::vector<int>& r1,const std::vector<int>& r2){
            return r1[1] > r2[1];
        });
        std::set<int> room_idx;//保存room size 大于要求的room index
        int j = 0;
        for(int i = 0;i < q_len;i++){
            int require_size = queries[i][1];
            int cur_id = queries[i][0];
            while(j < rooms.size() && rooms[j][1] >= require_size){
                room_idx.insert(rooms[j][0]);
                j++;
            }
            if (room_idx.empty()){
                res[queries[i][2]] = -1;
            }else{
                auto find = room_idx.lower_bound(cur_id);
                if(find == room_idx.end()){
                    res[queries[i][2]] =  *room_idx.rbegin();
                }else{
                    int idx = *find;
                    if(find != room_idx.begin()){
                        int big = *find;
                        auto pre = std::prev(find);
                        int small = *pre;
                        if(cur_id - small <= big - cur_id){
                            idx = small;
                        }
                    }
                    res[queries[i][2]] = idx;
                }
            }
        }
        return res;
    }
};