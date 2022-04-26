#include <vector>
#include <map>
using namespace std;

class RangeModule {
    std::map<int,int> range;
public:
    RangeModule() {

    }

    void addRange(int left, int right) {
        auto l = range.upper_bound(left);
        auto r = range.upper_bound(right);
    }

    bool queryRange(int left, int right) {
        return true;
    }

    void removeRange(int left, int right) {
        return;
    }
};