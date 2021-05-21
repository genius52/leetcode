#include <map>
using namespace std;

class MyCalendar {
    std::map<int,int> start_end;
public:
    MyCalendar() {

    }

    bool book(int start, int end) {
        auto it = start_end.lower_bound(end);
        if(it == start_end.begin()){
            start_end.insert({start,end});
            return true;
        }
        it--;
        if((*it).second <= start){
            start_end.insert({start,end});
            return true;
        }
        return false;
    }
};

