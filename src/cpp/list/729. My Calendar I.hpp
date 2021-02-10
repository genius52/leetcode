#include <set>
using namespace std;

class MyCalendar {
    std::set<std::pair<int,int>> start_end;
public:
    MyCalendar() {

    }

    bool book(int start, int end) {
        //在已有开始时间中查找小于等于当前的结束时间
        for(auto it = start_end.begin();it != start_end.end();it++){
            if(it->first >= end){
                break;
            }
            if(max(it->first,start) < min(it->second,end)){
                return false;
            }
        }
        start_end.insert(std::pair<int,int>{start,end});
        return true;
    }
};

