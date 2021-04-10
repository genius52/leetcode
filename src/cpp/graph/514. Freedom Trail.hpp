#include <string>
#include <vector>
#include <set>
#include <unordered_map>
using namespace std;

class Solution_514 {
public:
    int dp_findRotateSteps(int pos_ring,string ring,int len1,string key,int len2,int pos_key,
                           std::vector<std::set<int>>& char_position,std::unordered_map<int,int>& memo){
        if(pos_key >= len2){
            return 0;
        }
        int memo_key = pos_ring * 100 + pos_key;
        if(memo.find(memo_key) != memo.end()){
            return memo[memo_key];
        }
        if(ring[pos_ring] == key[pos_key]){
            return 1 + dp_findRotateSteps(pos_ring,ring,len1,key,len2,pos_key + 1,char_position,memo);
        }
        int char_index = key[pos_key] - 'a';
        int steps = 2147483647;
        for(auto it = char_position[char_index].begin();it != char_position[char_index].end();it++){
            int target_pos = *it;
            int cur_steps = min(abs(pos_ring - target_pos),min(pos_ring,target_pos) + len1 - max(pos_ring,target_pos));
            cur_steps += 1 + dp_findRotateSteps(*it,ring,len1,key,len2,pos_key + 1,char_position,memo);
            if(cur_steps < steps){
                steps = cur_steps;
            }
        }
        memo[memo_key] = steps;
        return steps;
    }

    int findRotateSteps(string ring, string key) {
        std::vector<std::set<int>> char_position(26);
        int len1 = ring.size();
        for(int i = 0;i < len1;i++){
            char_position[ring[i] - 'a'].insert(i);
        }
        int len2 = key.size();
        std::unordered_map<int,int> memo;
        return dp_findRotateSteps(0,ring,len1,key,len2,0,char_position,memo);
    }
};