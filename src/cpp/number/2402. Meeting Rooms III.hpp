#include <vector>
#include <queue>

//[[18,19],[3,12],[17,19],[2,13],[7,10]]
class Solution_2402 {
public:
    int mostBooked(int n, vector<vector<int>>& meetings) {
        std::sort(meetings.begin(),meetings.end());
        std::priority_queue<int, std::deque<int>, std::greater<int> > free_rooms;//smalltop;
        for(int i = 0;i < n;i++){
            free_rooms.push(i);
        }
        std::vector<int> record(n);
        std::priority_queue<std::pair<long long,int>,std::vector<std::pair<long long,int>>,std::greater<std::pair<long long,int>>> endtime_roomid;//smalltop
        for(int i = 0;i < meetings.size();i++){
            int start_time = meetings[i][0];
            int end_time = meetings[i][1];
            while (!endtime_roomid.empty() && endtime_roomid.top().first <= start_time){
                free_rooms.push(endtime_roomid.top().second);
                endtime_roomid.pop();
            }
            if(free_rooms.empty()){//没有空闲的会议室
                long long start_time2 = endtime_roomid.top().first;
                long long choose = endtime_roomid.top().second;
                endtime_roomid.pop();
                record[choose]++;
                long long end_time2 = start_time2 + end_time - start_time;
                endtime_roomid.push({end_time2,choose});
            }else{
                auto choose = free_rooms.top();
                free_rooms.pop();
                endtime_roomid.push({end_time,choose});
                record[choose]++;
            }
        }
        int roomid = 0;
        int max_cnt = 0;
        for(int i = 0;i < n;i++){
            if(record[i] > max_cnt){
                max_cnt = record[i];
                roomid = i;
            }

        }
        return roomid;
    }
};