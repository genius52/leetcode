#include <unordered_map>
#include <set>

class NumberContainers {
public:
    std::unordered_map<int,int> idx_num_;
    std::unordered_map<int,std::set<int>> num_idx_;
    NumberContainers() {

    }

    void change(int index, int number) {
        if(idx_num_.find(index) != idx_num_.end()){
            auto pre_num = idx_num_[index];
            num_idx_[pre_num].erase(index);
        }
        num_idx_[number].insert(index);
        idx_num_[index] = number;
    }

    int find(int number) {
        if(num_idx_.find(number) == num_idx_.end())
            return -1;
        if(num_idx_[number].size() == 0)
            return -1;
        return *num_idx_[number].begin();
    }
};