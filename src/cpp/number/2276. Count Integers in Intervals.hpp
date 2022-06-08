#include <map>
using namespace std;

class CountIntervals {
    std::map<int,int> data_;//left:right
    int cnt_ = 0;
public:
    CountIntervals() {

    }

    void add(int left, int right) {
        //找到第一个小于left的位置
        int new_left = left;
        int new_right = right;
        auto find = data_.upper_bound(left);//第一个大于
        if(find != data_.begin() && std::prev(find)->second >= left){
            find--;
        }
//        if(find != data_.end())
//            new_left = (*find).first;
        while(find != data_.end() && (*find).first <= right){
            cnt_ -= (*find).second - (*find).first + 1;
            new_left = min(new_left,(*find).first);
            new_right = max(new_right,(*find).second);
            data_.erase(find++);
        }
        cnt_ += new_right - new_left + 1;
        data_[new_left] = new_right;
    }

    int count() {
        return cnt_;
    }
};
