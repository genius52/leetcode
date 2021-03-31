#include <map>
#include <unordered_map>
#include <queue>
#include <list>
#include <string>

using namespace std;

class AllOne {
    std::unordered_map<std::string,int> key_count;
    std::map<int,std::list<std::string>> count_key;
//    std::priority_queue<std::pair<int,std::list<std::string>>> big_top;
//    std::priority_queue<std::pair<int,std::list<std::string>>,std::deque<std::pair<int,std::list<std::string>>>, std::greater<std::pair<int,std::string>>> small_top;
public:
    /** Initialize your data structure here. */
    AllOne() {

    }

    /** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
    void inc(string key) {
        auto it = key_count.find(key);
        if(it == key_count.end()){
            key_count[key] = 1;
            count_key[1].push_back(key);
            return;
        }

        int pre_count = key_count[key];
        for(auto it = count_key[pre_count].begin();it != count_key[pre_count].end();it++){
            if(*it == key){
                if(count_key[pre_count].size() == 1){
                    count_key.erase(pre_count);
                }else{
                    count_key[pre_count].erase(it);
                }
                break;
            }
        }
        count_key[pre_count + 1].push_back(key);
        key_count[key]++;
    }

    /** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
    void dec(string key) {
        auto it = key_count.find(key);
        if(it == key_count.end())
            return;
        int pre_count = key_count[key];
        for(auto it = count_key[pre_count].begin();it != count_key[pre_count].end();it++){
            if(*it == key){
                if(count_key[pre_count].size() == 1){
                    count_key.erase(pre_count);
                }else{
                    count_key[pre_count].erase(it);
                }
                break;
            }
        }
        if(pre_count > 1){
            count_key[pre_count - 1].push_back(key);
            key_count[key]--;
        }else{
            key_count.erase(key);
        }
    }

    /** Returns one of the keys with maximal value. */
    string getMaxKey() {
        if(count_key.empty())
            return "";
        return count_key.rbegin()->second.front();
    }

    /** Returns one of the keys with Minimal value. */
    string getMinKey() {
        if(count_key.empty())
            return "";
        return count_key.begin()->second.front();
    }
};