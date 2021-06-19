#include <set>

class ExamRoom {
    int used = 0;
    int size = 0;
    std::set<int> s;
public:
    ExamRoom(int n) {
        size = n;
    }
    int seat() {
        if(used == 0){
            s.insert(0);
            used++;
            return 0;
        }
        int idx = -1;
        int pre_seated = -1;
        int max_distance = 0;
        for(auto it = s.begin();it != s.end();it++){
            if(pre_seated == -1){
                max_distance = *it;
                if(*it != 0){
                    idx = 0;
                }
            }else{
                int cur_dis = (*it - pre_seated)/2;
                if(cur_dis > max_distance){
                    max_distance = cur_dis;
                    idx = (*it + pre_seated)/2;
                }
            }
            pre_seated = *it;
        }
        int cur_dis = size - 1 - pre_seated;
        if(cur_dis > max_distance){
            idx = size - 1;
        }
        s.insert(idx);
        used++;
        return idx;
    }

    void leave(int p) {
        if(s.find(p) != s.end()){
            used--;
            s.erase(p);
        }
    }
};