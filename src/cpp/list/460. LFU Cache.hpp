#include <deque>
#include <queue>
#include <map>
#include <unordered_set>
using namespace std;

class LFUCache {
    std::unordered_map<int,int> key_value;//key:value
    std::unordered_map<int,int> key_count;//key:frequency
    std::map<int,std::deque<int>> count_key;//frequncy -> key
    int cap = 0;
public:
    LFUCache(int capacity) {
        cap = capacity;
    }

    int get(int key) {
        if(cap == 0){
            return - 1;
        }
        auto it = key_value.find(key);
        if(it == key_value.end()){
            return -1;
        }
        int pre_count = key_count[key];
        key_count[key]++;
        if(count_key[pre_count].size() == 1){
            count_key.erase(pre_count);
        }else{
            for(auto it = count_key[pre_count].begin();it != count_key[pre_count].end();it++){
                if(*it == key){
                    count_key[pre_count].erase(it);
                    break;
                }
            }
        }
        count_key[pre_count + 1].push_back(key);
        return key_value[key];
    }

    void put(int key, int value) {
        if(cap == 0){
            return;
        }
        auto it = key_value.find(key);
        if(it == key_value.end()){//not exist
            int cur_size = key_value.size();
            if(cur_size >= cap){
                auto it = count_key.begin();
                int k = it->first;
                auto del_key = it->second.front();
                key_value.erase(del_key);
                key_count.erase(del_key);
                if(it->second.size() == 1){
                    count_key.erase(k);
                }else{
                    //it->second.erase(it->second.front());
                    it->second.erase(it->second.begin());
                    //std::deque<int> l = it->second;
                    //l.erase(l.begin());
                }
            }
            key_value[key] = value;
            key_count[key] = 1;
            count_key[1].push_back(key);
        }else{
            key_value[key] = value;
            int pre_count = key_count[key];
            key_count[key]++;
            if(count_key[pre_count].size() == 1){
                count_key.erase(pre_count);
            }else{
                for(auto it = count_key[pre_count].begin();it != count_key[pre_count].end();it++){
                    if(*it == key){
                        count_key[pre_count].erase(it);
                        break;
                    }
                }
            }
            count_key[pre_count + 1].push_back(key);
        }
    }
};