#include <set>
#include <list>
using namespace std;

class MKAverage {
    int m_ = 0;
    int k_ = 0;
    //int size_ = 0;
    std::multiset<int> small_;
    std::multiset<int> big_;//len = k
    std::multiset<int> mid_;//len = k
    std::list<int> l_;
    long sum_ = 0;
public:
    MKAverage(int m, int k) {
        m_ = m;
        k_ = k;
    }

    void addElement(int num) {
        l_.push_back(num);
        mid_.insert(num);
        sum_ += num;
        int size = l_.size();
        if(size > m_){
            int val = l_.front();
            l_.pop_front();
            if(small_.find(val) != small_.end()){
                small_.erase(small_.find(val));
            }else if(big_.find(val) != big_.end()){
                big_.erase(big_.find(val));
            }else if(mid_.find(val) != mid_.end()){
                mid_.erase(mid_.find(val));
                sum_ -= val;
            }
        }

        if(l_.size() == m_){
            if(!small_.empty()){
                while(*small_.rbegin() > *mid_.begin()){
                    int val1 = *small_.rbegin();
                    int val2 = *mid_.begin();
                    small_.erase(small_.find(*small_.rbegin()));
                    mid_.erase(mid_.find(*mid_.begin()));
                    small_.insert(val2);
                    mid_.insert(val1);
                    sum_ = sum_ + val1 - val2;
                }
            }else{
                while(small_.size() < k_){
                    small_.insert(*mid_.begin());
                    sum_ -= *mid_.begin();
                    mid_.erase(mid_.find(*mid_.begin()));
                }
                while(*small_.rbegin() > *mid_.begin()){
                    int val1 = *small_.rbegin();
                    int val2 = *mid_.begin();
                    small_.erase(small_.find(*small_.rbegin()));
                    mid_.erase(mid_.find(*mid_.begin()));
                    small_.insert(val2);
                    mid_.insert(val1);
                    sum_ = sum_ + val1 - val2;
                }
            }
            if(!big_.empty()){
                while(*big_.begin() < *mid_.rbegin()){
                    int val1 = *big_.begin();
                    int val2 = *mid_.rbegin();
                    big_.erase(big_.find(*big_.begin()));
                    mid_.erase(mid_.find(*mid_.rbegin()));
                    big_.insert(val2);
                    mid_.insert(val1);
                    sum_ = sum_ + val1 - val2;
                }
            }else{
                while(big_.size() < k_){
                    big_.insert(*mid_.rbegin());
                    sum_ -= *mid_.rbegin();
                    mid_.erase(mid_.find(*mid_.rbegin()));
                }
                while(*big_.begin() < *mid_.rbegin()){
                    int val1 = *big_.begin();
                    int val2 = *mid_.rbegin();
                    big_.erase(big_.find(*big_.begin()));
                    mid_.erase(mid_.find(*mid_.rbegin()));
                    big_.insert(val2);
                    mid_.insert(val1);
                    sum_ = sum_ + val1 - val2;
                }
            }
        }
    }

    int calculateMKAverage() {
        if(l_.size() < m_)
            return -1;
        return sum_/(m_ - k_ * 2);
    }
};