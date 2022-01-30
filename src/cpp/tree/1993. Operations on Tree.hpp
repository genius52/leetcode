#include <vector>
#include <unordered_map>
using namespace std;

class LockingTree {
    std::vector<int> parent_;
    std::vector<std::vector<int>> children_;
    std::unordered_map<int,int> lock_;
public:
    LockingTree(vector<int>& parent) {
        parent_ = parent;
        int len = parent.size();
        children_.resize(len);
        for(int i = 0;i < len;i++){
            if(parent[i] != -1)
                children_[parent[i]].push_back(i);
        }
    }

    bool lock(int num, int user) {
        if(lock_.find(num) == lock_.end()){
            lock_[num] = user;
            return true;
        }
        return false;
    }

    bool unlock(int num, int user) {
        if(lock_.find(num) != lock_.end()){
            if(lock_[num] == user){
                lock_.erase(num);
                return true;
            }
        }
        return false;
    }
//Upgrade：指定用户给指定节点上锁，并且将该节点的所有子孙节点解锁。只有如下 3 个条件全部满足时才能执行升级操作：
//指定节点当前状态为未上锁。
//指定节点至少有一个上锁状态的子孙节点（可以是任意用户上锁的）。
//指定节点没有任何上锁的祖先节点。
    bool dfs_upgrade(int num){
        bool res = false;
        if(lock_.find(num) != lock_.end()){
            lock_.erase(num);
            res = true;
        }
        for(auto c : children_[num]){
            res |= dfs_upgrade(c);
        }
        return res;
    }

    bool upgrade(int num, int user) {
        //指定节点当前状态为未上锁
        if(lock_.find(num) != lock_.end())
            return false;
        //指定节点没有任何上锁的祖先节点
        int visit = parent_[num];
        while(visit != -1){
            if(lock_.find(visit) != lock_.end())
                return false;
            visit = parent_[visit];
        }
        //指定节点至少有一个上锁状态的子孙节点（可以是任意用户上锁的）。
        if(dfs_upgrade(num)){
            lock_[num] = user;
            return true;
        }
        return false;
    }
};