#include <vector>
#include <list>
using namespace std;

class MyHashSet {
    vector<std::list<int>> table;
    int factor = 10086;
public:
    /** Initialize your data structure here. */
    MyHashSet() {
        table.resize(factor);
        //v = vector<int>(1000000,-1);
    }

    void add(int key) {
        int pos = key % factor;
        if(table[pos].size() == 0){
            table[pos].push_back(key);
            return;
        }
        for(auto it = table[pos].begin();it != table[pos].end();it++){
            if(*it == key){
                return;
            }
        }
        table[pos].push_back(key);
    }

    void remove(int key) {
        int pos = key % factor;
        if(table[pos].size() == 0)
            return;
        for(auto it = table[pos].begin();it != table[pos].end();it++){
            if(*it == key){
                table[pos].erase(it);
                break;
            }
        }
    }

    /** Returns true if this set contains the specified element */
    bool contains(int key) {
        int pos = key % factor;
        if(table[pos].size() == 0)
            return false;
        for(auto it = table[pos].begin();it != table[pos].end();it++){
            if(*it == key){
                return true;
            }
        }
        return false;
    }
};