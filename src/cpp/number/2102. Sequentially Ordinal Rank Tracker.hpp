#include <string>
#include <map>
#include <queue>
using namespace std;

struct my_cmp_smalltop{
    bool operator()(const std::pair<int,std::string>& p1,const std::pair<int,std::string>& p2) {
        if(p1.first == p2.first){
            return p1.second > p2.second;
        }else{
            return p1.first < p2.first;
        }
    }
};

struct my_cmp_bigtop{
    bool operator()(const std::pair<int,std::string>& p1,const std::pair<int,std::string>& p2) {
        if(p1.first == p2.first){
            return p1.second < p2.second;
        }else{
            return p1.first > p2.first;
        }
    }
};

class SORTracker {
    int get_cnt_ = 1;
    std::priority_queue<std::pair<int,std::string>,std::vector<std::pair<int,std::string>>, my_cmp_bigtop> big_nums_;
    std::priority_queue<std::pair<int,std::string>,std::vector<std::pair<int,std::string>>, my_cmp_smalltop> small_nums_;
public:
    SORTracker() {

    }

    void add(string name, int score) {
        big_nums_.push(std::make_pair(score,name));
        if(big_nums_.size() > get_cnt_){
            small_nums_.push(big_nums_.top());
            big_nums_.pop();
        }
    }

    string get() {
        auto ret = big_nums_.top();
        if(!small_nums_.empty()){
            big_nums_.push(small_nums_.top());
            small_nums_.pop();
        }
        get_cnt_++;
        return ret.second;
    }
};