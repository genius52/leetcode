#include <unordered_map>
#include <map>
#include <vector>
using namespace std;
class Solution_1481 {
public:
    int findLeastNumOfUniqueInts(vector<int>& arr, int k) {
        std::unordered_map<int, int> record;
        int len = arr.size();
        for (int i = 0; i < len; i++) {
            record[arr[i]]++;
        }
        int total = record.size();
        std::map<int, std::vector<int>> cnt_record;
        for (auto it = record.begin(); it != record.end(); it++) {
            cnt_record[(*it).second].push_back((*it).first);
        }
        for (auto it = cnt_record.begin(); it != cnt_record.end(); it++) {
            auto cnt = (*it).first;
            auto num_len = (*it).second.size();
            while (num_len > 0) {
                if (k < cnt)
                    return total;
                k = k - cnt;
                total--;
                num_len--;
            }
        }
        return total;
    }
};