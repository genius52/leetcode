#include <unordered_map>
#include <vector>
#include <list>
using namespace std;


class LRUCache {
    std::list<int> lru;
    std::unordered_map<int,std::list<int>::iterator> key_position;//key -> index
    std::unordered_map<int,int> key_value;
    int cap = 0;
public:
    LRUCache(int capacity) {
        this->cap = capacity;
    }

    int get(int key) {
        if(key_value.find(key) == key_value.end())
            return -1;
        if (!lru.empty())
            lru.erase(key_position[key]);
        lru.push_front(key);
        key_position[key] = lru.begin();
        return key_value[key];
    }

    void put(int key, int value) {
        if(key_value.find(key) != key_value.end()){//已存在，更新value和位置
            if (key_position.size() > 0)
                lru.erase(key_position[key]);
            lru.push_front(key);
            key_position[key] = lru.begin();
            key_value[key] = value;
        }else{
            if(cap == key_value.size()){//不存在，去除旧的value
                if(!lru.empty() && !key_position.empty() && !key_value.empty()){
                    key_position.erase(lru.back());
                    key_value.erase(lru.back());
                    lru.pop_back();
                }
                lru.push_front(key);
                key_value[key] = value;
                key_position[key] = lru.begin();
            }else{//不存在，添加新的
                lru.push_front(key);
                key_position[key] = lru.begin();
                key_value[key] = value;
            }
        }
    }
};