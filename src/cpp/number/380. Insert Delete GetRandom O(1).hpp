#include <unordered_set>
#include <unordered_map>
#include <vector>

class RandomizedSet {
public:
    std::unordered_map<int,int> val_idx_;
    std::vector<int> data_;
    /** Initialize your data structure here. */
    RandomizedSet() {

    }

    /** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
    bool insert(int val) {
        if(val_idx_.find(val) != val_idx_.end()){
            return false;
        }
        data_.emplace_back(val);
        val_idx_[val] = data_.size() - 1;
        return true;
    }

    /** Removes a value from the set. Returns true if the set contained the specified element. */
    bool remove(int val) {
        if(val_idx_.find(val) == val_idx_.end()){
            return false;
        }
        int idx = val_idx_[val];
        int tail = data_[data_.size() - 1];
        std::swap(data_[idx],data_[data_.size() - 1]);
        data_.pop_back();
        val_idx_[tail] = idx;
        val_idx_.erase(val);
        return true;
    }

    /** Get a random element from the set. */
    int getRandom() {
        int idx = rand() % data_.size();
        return data_[idx];
    }
};

//class RandomizedSet {
//public:
//    std::unordered_set<int> s;
//    /** Initialize your data structure here. */
//    RandomizedSet() {
//
//    }
//
//    /** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
//    bool insert(int val) {
//        if(s.find(val) != s.end()){
//            return false;
//        }
//        s.insert(val);
//        return true;
//    }
//
//    /** Removes a value from the set. Returns true if the set contained the specified element. */
//    bool remove(int val) {
//        if(s.find(val) == s.end()){
//            return false;
//        }
//        s.erase(val);
//        return true;
//    }
//
//    /** Get a random element from the set. */
//    int getRandom() {
//        int len = s.size();
//        int res = 0;
//        int cnt = 1;
//        for(auto it = s.begin();it != s.end();it++){
//            int r = rand() % cnt;
//            if(r % cnt++ == 0){
//                res = *it;
//            }
//        }
//        return res;
//    }
//};