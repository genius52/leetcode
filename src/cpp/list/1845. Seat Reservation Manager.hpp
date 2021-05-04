#include <set>
using namespace std;

class SeatManager {
    std::set<int> s;
public:
    SeatManager(int n) {
        for(int i = 1;i <= n;i++){
            s.insert(i);
        }
    }

    int reserve() {
        auto res = *(s.begin());
        s.erase(s.begin());
        return res;
    }

    void unreserve(int seatNumber) {
        s.insert(seatNumber);
    }
};