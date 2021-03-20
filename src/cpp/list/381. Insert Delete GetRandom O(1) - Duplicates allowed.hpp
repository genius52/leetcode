#include <set>

class RandomizedCollection {
    std::multiset<int> data;
public:
    /** Initialize your data structure here. */
    RandomizedCollection() {

    }

    /** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
    bool insert(int val) {
        if(data.find(val) != data.end()){
            data.insert(val);
            return false;
        }else{
            data.insert(val);
            return true;
        }
    }

    /** Removes a value from the collection. Returns true if the collection contained the specified element. */
    bool remove(int val) {
        auto it = data.find(val);
        if(it == data.end()){
            return false;
        }else{
            data.erase(it);
            return true;
        }
    }

    /** Get a random element from the collection. */
    int getRandom() {
        int cnt = 1;
        //std::multiset<int>::iterator res;
        int res;
        for(auto it = data.begin();it != data.end();it++){
            int random = rand() % cnt++;
            if(random == 0){
                res = *it;
            }
        }
        return res;
    }
};

