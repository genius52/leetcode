#include <vector>
#include <map>
#include <unordered_map>
#include <queue>
using namespace std;

////Input: target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]
////Output: 2
////Explanation: We start with 10 liters of fuel.
////We drive to position 10, expending 10 liters of fuel.  We refuel from 0 liters to 60 liters of gas.
////Then, we drive from position 10 to position 60 (expending 50 liters of fuel),
////and refuel from 10 liters to 50 liters of gas.  We then drive to and reach the target.
////We made 2 refueling stops along the way, so we return 2.
class Solution_871 {
public:
    //DP from bottom to top   TO DO
    
    //BFS
    int minRefuelStops(int target, int startFuel, vector<vector<int>>& stations) {
        int len = stations.size();
        int cur_pos = startFuel;
        std::priority_queue<int> q;
        int index = 0;
        int steps = 0;
        while(cur_pos < target){
            while(index < len){
                if(stations[index][0] > cur_pos){
                    break;
                }
                q.push(stations[index][1]);
                index++;
            }
            if(q.empty()){
                return -1;
            }
            cur_pos += q.top();
            q.pop();
            steps++;
        }
        return steps;
    }

    //DP from top to bottom TLE
//    int dp_minRefuelStops(int target,int cur_pos,int fuel,std::map<int,int>& pos_fuel,std::unordered_map<int,std::unordered_map<int,int>>& memo){
//        if(cur_pos >= target){
//            return 0;
//        }
//        if(fuel <= 0)
//            return -1;
//        int max_dis = cur_pos + fuel;
//        if(max_dis >= target){
//            return 0;
//        }
//        if(memo.find(cur_pos) != memo.end()){
//            if(memo[cur_pos].find(fuel) != memo[cur_pos].end()){
//                return memo[cur_pos][fuel];
//            }
//        }
//
//        auto find = pos_fuel.upper_bound(cur_pos);
//        if(find == pos_fuel.end()){
//            return -1;
//        }
//        int next_pos = (*find).first;
//        if((cur_pos + fuel) < next_pos){
//            return -1;
//        }
//        int gas =  (*find).second;
//        //stop at find station
//        int ret1 = dp_minRefuelStops(target,next_pos,fuel - (next_pos - cur_pos) + gas,pos_fuel,memo);
//        if(ret1 != -1)
//            ret1++;
//        //Do not stop at find station
//        int ret2 = dp_minRefuelStops(target,next_pos,fuel - (next_pos - cur_pos) ,pos_fuel,memo);
//        if(ret1 == -1 && ret2 == -1){
//            memo[cur_pos][fuel] = -1;
//        }else if(ret1 == -1){
//            memo[cur_pos][fuel] = ret2;
//        }else if(ret2 == -1){
//            memo[cur_pos][fuel] = ret1;
//        }else{
//            memo[cur_pos][fuel] = min(ret1,ret2);
//        }
//        return memo[cur_pos][fuel];
//    }
//
//    int minRefuelStops(int target, int startFuel, vector<vector<int>>& stations) {
//        std::map<int,int> pos_fuel;
//        int len = stations.size();
//        for(int i = 0;i < len;i++){
//            pos_fuel[stations[i][0]] = stations[i][1];
//        }
//        std::unordered_map<int,std::unordered_map<int,int>> memo;
//        auto res = dp_minRefuelStops(target,0,startFuel,pos_fuel,memo);
//        return res;
//    }
};