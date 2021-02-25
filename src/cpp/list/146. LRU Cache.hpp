#include <unordered_map>
#include <vector>
#include <list>
using namespace std;

struct Node{
    int key;
    int val;
};

class LRUCache {
    std::list<Node*> l;
    int capacity = 0;
    std::unordered_map<int,Node*> key_pos;//key -> index
public:
    LRUCache(int capacity) {
        this->capacity = capacity;
    }

    int get(int key) {
        auto it = key_pos.find(key);
        if(it == key_pos.end())
            return -1;
        else{
            auto val = it->second->val;
            auto it = l.begin();
            while(it != l.end()){
                if((*it)->key == key){
                    break;
                }
                it++;
            }
            l.erase(it);
            Node* node2 = new Node();
            node2->key = key;
            node2->val = val;
            l.push_back(node2);
            key_pos[key] = node2;
            return val;
        }
    }

    void put(int key, int value) {
        auto cur_size = l.size();
        auto it = key_pos.find(key);
        if(it == key_pos.end()){
            Node* node = new Node();
            node->key = key;
            node->val = value;
            if(cur_size == capacity){
                auto front = l.front();
                key_pos.erase(front->key);
                l.pop_front();
            }
            l.push_back(node);
            key_pos[key] = node;
        }else{
            Node* old = it->second;
            old->val = value;
            auto it = l.begin();
            while(it != l.end()){
                if((*it)->key == key){
                    break;
                }
                it++;
            }
            l.erase(it);
            l.push_back(old);
        }
    }
};