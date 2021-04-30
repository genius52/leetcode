#include <string>
#include <queue>
#include <set>
using namespace std;

//第一个参议员来自  Radiant
class Solution_649 {
public:
    string predictPartyVictory(string senate) {
        //Radiant and the Dire
        std::set<int> set_r;
        std::set<int> set_d;
        int len = senate.size();
        for(int i = 0;i < len;i++){
            if(senate[i] == 'R'){
                set_r.insert(i);
            }else{
                set_d.insert(i);
            }
        }
        //bool r_round = true;
        int cur_pos = 0;
        auto it_r = set_r.begin();
        auto it_d = set_d.begin();
        while(!set_r.empty() && !set_d.empty()){
            while(set_d.find(cur_pos) == set_d.end() && set_r.find(cur_pos) == set_r.end()){
                cur_pos = (cur_pos + 1) % len;
            }
            if(senate[cur_pos] == 'R'){
                //set_d.erase(set_d.begin());
                auto to_del = set_d.lower_bound(cur_pos);
                if(to_del != set_d.end()){
                    set_d.erase(to_del);
                }else{
                    set_d.erase(set_d.begin());
                }
            }else{
                //set_r.erase(set_r.begin());
                auto to_del = set_r.lower_bound(cur_pos);
                if(to_del != set_r.end()){
                    set_r.erase(to_del);
                }else{
                    set_r.erase(set_r.begin());
                }
            }
            cur_pos = (cur_pos + 1) % len;
        }
        if(set_r.empty())
            return "Dire";
        else
            return "Radiant";
    }
};