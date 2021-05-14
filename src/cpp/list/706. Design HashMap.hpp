#include <vector>
#include <list>
using namespace std;

class MyHashMap {
    vector<std::list<std::pair<int,int>>> table;
    int factor = 10086;//TO DO,should support rehash when load_factor is too big
public:
    /** Initialize your data structure here. */
    MyHashMap() {
        table.resize(factor);
    }

    int hash(int val){
        return val % factor;
    }
    /** value will always be non-negative. */
    void put(int key, int value) {
        int pos = hash(key);
        if(table[pos].size() == 0){
            std::pair<int,int> p;
            p.first = key;
            p.second = value;
            table[pos].push_back(p);
        }else{
            for(auto it = table[pos].begin();it != table[pos].end();it++){
                if((*it).first == key){
                    (*it).second = value;
                    return;
                }
            }
            std::pair<int,int> p;
            p.first = key;
            p.second = value;
            table[pos].push_back(p);
        }
    }

    /** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
    int get(int key) {
        int pos = hash(key);
        if(table[pos].size() == 0){
            return -1;
        }
        for(auto it = table[pos].begin();it != table[pos].end();it++){
            if((*it).first == key){
                return (*it).second;
            }
        }
        return -1;
    }

    /** Removes the mapping of the specified value key if this map contains a mapping for the key */
    void remove(int key) {
        int pos = hash(key);
        if(table[pos].size() == 0){
            return;
        }
        for(auto it = table[pos].begin();it != table[pos].end();it++){
            if((*it).first == key){
                table[pos].erase(it);
                return;
            }
        }
    }
};