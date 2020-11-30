#include <string>
#include <unordered_map>
#include <map>
using namespace std;

class TimeMap {
private:
    std::unordered_map<std::string,std::map<int,std::string>> record;
public:
    /** Initialize your data structure here. */
    TimeMap() {

    }

    void set(string key, string value, int timestamp) {
        record[key][timestamp] = value;
    }

    string get(string key, int timestamp) {
        if (record.find(key) == record.end()){
            return "";
        }
        auto it = record[key].upper_bound(timestamp);
        if (it != record[key].begin()){
            it--;
            return (*it).second;
        }else{
            return "";
        }
    }
};


