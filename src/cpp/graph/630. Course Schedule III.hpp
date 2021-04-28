#include <vector>
#include <queue>
using namespace std;

class Solution_630 {
public:
    //  TLE
//    int dp_scheduleCourse(vector<vector<int>>& courses,int len,int cur,int timestamp,
//                          std::vector<bool>& visited){
//        if(cur >= len)
//            return 0;
//        if(visited[cur])
//            return 0;
//        if(timestamp > courses[cur][1])
//            return 0;
//        timestamp += courses[cur][0];
//        visited[cur] = true;
//        int max_course = 1;
//        for(int i = 0;i < len;i++){
//            if(visited[i])
//                continue;
//            if((timestamp + courses[i][0]) > courses[i][1]){
//                continue;
//            }
//            int course_cnt = 1 + dp_scheduleCourse(courses,len,i,timestamp,visited);
//            if(course_cnt > max_course)
//                max_course = course_cnt;
//        }
//        visited[cur] = false;
//        return max_course;
//    }
//
//    int scheduleCourse(vector<vector<int>>& courses){
//        int len = courses.size();
//        int max_course = 0;
//        for(int i = 0;i < len;i++){
//            std::vector<bool> visited(len);
//            int course_cnt = dp_scheduleCourse(courses,len,i,0,visited);
//            if(course_cnt > max_course)
//                max_course = course_cnt;
//        }
//        return max_course;
//    }

    int scheduleCourse(vector<vector<int>>& courses) {
        std::sort(courses.begin(),courses.end(),[](const std::vector<int>& c1, const std::vector<int>& c2){
            return c1[1] < c2[1];
        });
        std::priority_queue<int> q;//big top stack
        int cur_day = 0;
        int len = courses.size();
        for(int i = 0;i < len;i++){
            q.push(courses[i][0]);
            cur_day += courses[i][0];
            if(cur_day > courses[i][1]){
                cur_day -= q.top();
                q.pop();
            }
        }
        return q.size();
    }
};