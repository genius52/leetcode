#include <set>

class RecentCounter {
    std::set<int> s;
public:
    RecentCounter() {

    }

    int ping(int t) {
        s.insert(t);
        auto it = s.lower_bound(t - 3000);
        return std::distance(it,s.end());
    }
};

