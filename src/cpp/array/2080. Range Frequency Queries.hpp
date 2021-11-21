#include <vector>
#include <unordered_map>
using namespace std;

class RangeFreqQuery {
    std::unordered_map<int, std::vector<int>> record_;
public:
    RangeFreqQuery(vector<int>& arr) {
        int len = arr.size();
        for (int i = 0; i < len; i++) {
            record_[arr[i]].push_back(i);
        }
    }

    int query(int left, int right, int value) {
        if (record_.find(value) == record_.end())
            return 0;
        auto start = std::lower_bound(record_[value].begin(), record_[value].end(),left);
        auto end = std::upper_bound(record_[value].begin(), record_[value].end(), right);
        return std::distance(start, end);
    }
};