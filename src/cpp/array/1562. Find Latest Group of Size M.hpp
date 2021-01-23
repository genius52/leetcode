#include <vector>
using namespace std;

class Solution_1562 {
public:
    int findLatestStep(vector<int>& arr, int m) {
        int len = arr.size();
        if(len == 1){
            if (m == 1){
			    return 1;
            }else{
                return -1;
            }
        }
        int steps = -1;
        std::vector<int> record(len + 1);
        std::vector<int> length_cnt(len + 1);
        for (int i = 0;i < len;i++) {
            if (arr[i] == 1) {
                record[arr[i]] = record[arr[i] + 1] + 1;
                length_cnt[record[arr[i]]]++;
                if (record[arr[i] + 1] != 0) {
                    int j = arr[i] + 1;
                    length_cnt[record[arr[i] + 1]]--;
                    int old_val = record[j];
                    fill(record.begin() + j,record.begin() + j + old_val,record[arr[i]]);
                }
            } else if (arr[i] == len) {
                record[arr[i]] = record[arr[i] - 1] + 1;
                length_cnt[record[arr[i]]]++;
                if (record[arr[i] - 1] != 0) {
                    int j = arr[i] - 1;
                    length_cnt[record[arr[i] - 1]]--;
                    int old_val = record[j];
                    fill(record.begin() + j - old_val + 1,record.begin() + j + 1,record[arr[i]]);
                }
            } else {
                record[arr[i]] = record[arr[i] - 1] + record[arr[i] + 1] + 1;
                length_cnt[record[arr[i]]]++;
                if (record[arr[i] + 1] != 0) {
                    int j = arr[i] + 1;
                    length_cnt[record[arr[i] + 1]]--;
                    int old_val = record[j];
                    int set_val = record[arr[i]];
                    fill(record.begin() + j,record.begin() + j + old_val,record[arr[i]]);
                }
                if (record[arr[i] - 1] != 0) {
                    int j = arr[i] - 1;
                    length_cnt[record[arr[i] - 1]]--;
                    int old_val = record[j];
                    fill(record.begin() + j - old_val + 1,record.begin() + j + 1,record[arr[i]]);
                }
            }
            if (length_cnt[m] > 0) {
                steps = i + 1;
            }
        }
        return steps;
    }
};