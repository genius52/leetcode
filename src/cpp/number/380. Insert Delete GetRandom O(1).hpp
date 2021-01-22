#include <unordered_set>

class RandomizedSet {
public:
    std::unordered_set<int> s;
    /** Initialize your data structure here. */
    RandomizedSet() {

    }

    /** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
    bool insert(int val) {
        if(s.find(val) != s.end()){
            return false;
        }
        s.insert(val);
        return true;
    }

    /** Removes a value from the set. Returns true if the set contained the specified element. */
    bool remove(int val) {
        if(s.find(val) == s.end()){
            return false;
        }
        s.erase(val);
        return true;
    }

    /** Get a random element from the set. */
    int getRandom() {
        int len = s.size();
        int res = 0;
        int cnt = 1;
        for(auto it = s.begin();it != s.end();it++){
            int r = rand() % cnt;
            if(r % cnt++ == 0){
                res = *it;
            }
        }
        return res;
    }
};

