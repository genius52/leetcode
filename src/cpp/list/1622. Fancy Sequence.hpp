#include <vector>
using namespace std;

//Fancy()初始化一个空序列对象。
//void append(val) 将整数val添加在序列末尾。
//void addAll(inc)将所有序列中的现有数值都增加inc。
//void multAll(m)将序列中的所有现有数值都乘以整数m。
//int getIndex(idx) 得到下标为idx处的数值（下标从 0 开始），
// 并将结果对109 + 7取余。如果下标大于等于序列的长度，请返回-1。

//TLE
class Fancy {
    std::vector<int> data_;
    std::vector<std::vector<int>> prefix_sum_;
    std::vector<std::vector<int>> prefix_multiple_;
    int size_ = -1;
public:
    Fancy() {

    }

    void append(int val) {
        data_.push_back(val);
        prefix_sum_.push_back(std::vector<int>{0});
        prefix_multiple_.push_back(std::vector<int>{1});
        size_ = data_.size();
    }

    void addAll(int inc) {
        if (size_ < 0)
            return;
        prefix_sum_[size_ - 1].push_back(inc);
        prefix_multiple_[size_ - 1].push_back(1);
    }

    void multAll(int m) {
        if (size_ < 0)
            return;
        prefix_sum_[size_ - 1].push_back(0);
        prefix_multiple_[size_ - 1].push_back(m);
    }

    int getIndex(int idx) {
        if (idx >= size_)
            return - 1;
        int i = 0;
        long res = data_[idx];
        int MOD = 1e9 + 7;
        for (int i = idx; i < size_; ++i) {
            for (int j = 0; j < prefix_sum_[i].size(); ++j) {
                res += prefix_sum_[i][j];
                res *= prefix_multiple_[i][j];
                res %= MOD;
            }
        }
        return res;
    }
};