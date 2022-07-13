#include <set>

class SmallestInfiniteSet {
    std::set<int> added_;
    int smallest_ = 1;
public:
    SmallestInfiniteSet() {

    }

    int popSmallest() {
        if(added_.empty()){
            return smallest_++;
        }else{
            auto res = *added_.begin();
            added_.erase(added_.begin());
            return res;
        }
    }

    void addBack(int num) {
        if(num < smallest_){
            added_.insert(num);
        }
    }
};