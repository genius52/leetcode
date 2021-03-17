#include <vector>
using namespace std;

class SummaryRanges {
    //std::priority_queue<std::pair<int,int>> q;
    std::vector<std::vector<int>> data;
public:
    /** Initialize your data structure here. */
    SummaryRanges() {

    }

    void addNum(int val) {
        int len = data.size();
        if(len == 0){
            data.push_back(std::vector<int>{val,val});
        }else{
            bool find = false;
            for(auto it = data.begin();it != data.end();it++){
                if(val >= (*it)[0] && val <= (*it)[1]){
                    find = true;
                    break;
                }else if(val < (*it)[0] - 1){
                    data.insert(it,std::vector<int>{val,val});
                    find = true;
                    break;
                }else if(val == (*it)[0] - 1){
                    (*it)[0] = val;
                    find = true;
                    break;
                }else if(val == (*it)[1] + 1){
                    find = true;
                    auto next = it;
                    next++;
                    if(next == data.end()){
                        (*it)[1]++;
                    }else{
                        if((*next)[0] == val + 1){
                            (*it)[1] = (*next)[1];
                            data.erase(next);
                        }else{
                            (*it)[1]++;
                        }
                        break;
                    }
                }
            }
            if(!find){
                data.push_back(std::vector<int>{val,val});
            }
        }
    }

    vector<vector<int>> getIntervals() {
        return data;
    }
};