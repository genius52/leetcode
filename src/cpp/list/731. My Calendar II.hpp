#include <set>
using namespace std;

//MyCalendar();
//MyCalendar.book(10, 20); // returns true
//MyCalendar.book(50, 60); // returns true
//MyCalendar.book(10, 40); // returns true
//MyCalendar.book(5, 15); // returns false
//MyCalendar.book(5, 10); // returns true
//MyCalendar.book(25, 55); // returns true
class MyCalendarTwo {
    std::set<std::pair<int,int>> start_end;
    std::set<std::pair<int,int>> dup_section;
public:
    MyCalendarTwo() {

    }

    bool book(int start, int end) {
        int dup_cnt = 0;
        int dup_start = 0;
        int dup_end = 0;
        for(auto it = dup_section.begin();it != dup_section.end();it++){
            if(max(it->first,start) < min(it->second,end)){
                return false;
            }
        }
        std::vector<std::pair<int,int>> cur_dup;
        for(auto it = start_end.begin();it != start_end.end();it++){
            auto cur_start = max(it->first,start);
            auto cur_end = min(it->second,end);
            if(cur_start < cur_end){
                cur_dup.push_back(std::pair<int,int>{cur_start,cur_end});
            }
        }
        start_end.insert(std::pair<int,int>{start,end});
        for(auto it = cur_dup.begin();it != cur_dup.end();it++){
            dup_section.insert(std::pair<int,int>{it->first,it->second});
        }
        return true;
    }
};

