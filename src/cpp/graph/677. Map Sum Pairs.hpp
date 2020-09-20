#include <string>
#include <unordered_map>
using namespace std;

class Charnode{
public:
    Charnode(){
        for(int i = 0;i < 26;i++){
            children[i] = nullptr;
        }
    }
    Charnode* children[26];
    int count = 0;
    bool is_word = false;
};

class MapSum {
public:
    /** Initialize your data structure here. */
    MapSum() {
        root = new Charnode();
    }

    void insert(string key, int val) {
        bool exist = false;
        int old_val = 0;
        if(record.find(key) != record.end()){
            old_val = record[key];
            exist = true;
        }
        record[key] = val;
        int diff = val - old_val;
        Charnode* visit = root;
        int l = key.size();
        for(int i = 0;i < l;i++){
            if(visit->children[key[i] - 'a'] == nullptr){
                visit->children[key[i] - 'a'] = new Charnode();
            }
            if(exist){
                visit->children[key[i] - 'a']->count += diff;
            }else{
                visit->children[key[i] - 'a']->count += val;
            }
            visit = visit->children[key[i] - 'a'];
        }
    }

    int sum(string prefix) {
        int l = prefix.size();
        Charnode* visit = root;
        for (int i = 0;i < l;i++){
            if(visit->children[prefix[i] - 'a'] != nullptr){
                visit = visit->children[prefix[i] - 'a'];
            }else{
                return 0;
            }
        }
        if(visit == nullptr)
            return 0;
        return visit->count;
    }

    std::unordered_map<string,int> record;
    Charnode* root;
};
