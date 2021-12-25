#include <set>
#include <queue>
using namespace std;

class SeatManager {
    std::priority_queue<int,std::vector<int>,std::greater<int>> q_;
public:
    SeatManager(int n) {
        for(int i = 1;i <= n;i++){
            q_.push(i);
        }
    }

    int reserve() {
        int res = q_.top();
        q_.pop();
        return res;
    }

    void unreserve(int seatNumber) {
        q_.push(seatNumber);
    }
};

//class SeatManager {
//    std::set<int> s;
//public:
//    SeatManager(int n) {
//        for(int i = 1;i <= n;i++){
//            s.insert(i);
//        }
//    }
//
//    int reserve() {
//        auto res = *(s.begin());
//        s.erase(s.begin());
//        return res;
//    }
//
//    void unreserve(int seatNumber) {
//        s.insert(seatNumber);
//    }
//};