#include <map>
#include <set>
#include <stack>
using namespace std;

class DinnerPlates {
    std::map<int,std::vector<int>> m_;
    std::set<int> avaliable_;
    int cap_ = 0;
public:
    DinnerPlates(int capacity) {
        cap_ = capacity;
    }

    void push(int val) {
        auto it = avaliable_.begin();
        if(it == avaliable_.end()){
            std::vector<int> v;
            v.push_back(val);
            int len = m_.size();
            m_[len] = v;
            if(v.size() < cap_)
                avaliable_.insert(len);
        }else{
            m_[*it].push_back(val);
            if(m_[*it].size() >= cap_){
                avaliable_.erase(*it);
            }
        }
    }

    int pop() {
        auto it = m_.rbegin();
        if(it == m_.rend()){
            return -1;
        }else{
            int res = m_[m_.rbegin()->first].back();
            m_[m_.rbegin()->first].pop_back();
            auto idx = m_.rbegin()->first;
            if(m_[m_.rbegin()->first].size() == 0)
                m_.erase(m_.rbegin()->first);
            avaliable_.insert(idx);
            return res;
        }
    }

    int popAtStack(int index) {
        if(index >= m_.size())
            return -1;
        if(m_[index].empty())
            return -1;
        else{
            auto res = m_[index].back();
            m_[index].pop_back();
            if(m_[index].empty())
                m_.erase(index);
            avaliable_.insert(index);
            return res;
        }
    }
};