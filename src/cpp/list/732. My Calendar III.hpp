#include <map>
using namespace std;

//Input
//["MyCalendarThree", "book", "book", "book", "book", "book", "book"]
//[[], [10, 20], [50, 60], [10, 40], [5, 15], [5, 10], [25, 55]]
//Output
//[null, 1, 1, 2, 3, 3, 3]
class MyCalendarThree {
    std::map<int,int> timestamp_cnt;
public:
    MyCalendarThree() {

    }

    int book(int start, int end) {
        timestamp_cnt[start]++;
        timestamp_cnt[end]--;
        int res = 0;
        int cur_dup = 0;
        for(auto it = timestamp_cnt.begin();it != timestamp_cnt.end();it++){
            cur_dup += (*it).second;
            res = max(res,cur_dup);
        }
        return res;
    }
};

