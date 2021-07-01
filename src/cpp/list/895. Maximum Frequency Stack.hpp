#include <map>
#include <unordered_map>
#include <unordered_set>
#include <vector>
#include <list>
using namespace std;

//push(int x)，将整数 x 推入栈中。
//pop()，它移除并返回栈中出现最频繁的元素。
//如果最频繁的元素不只一个，则移除并返回最接近栈顶的元素。
class FreqStack{
    std::unordered_map<int,int> num_freq;//数字：频率
    std::vector<std::vector<int>> freq_numbers;//频率 ：数字队列
    int max_freq = 0;//最大频率
public:
    FreqStack() {

    }
    void push(int val) {
        num_freq[val]++;
        max_freq = max(max_freq,num_freq[val]);
        if(max_freq > freq_numbers.size())
            freq_numbers.push_back({val});
        else{
            if(num_freq.find(val) == num_freq.end()){
                freq_numbers[0].push_back(val);
            }else{
                freq_numbers[num_freq[val] - 1].push_back(val);
            }
        }
    }
    int pop() {
        int res = freq_numbers[max_freq - 1].back();
        freq_numbers[max_freq - 1].pop_back();
        num_freq[res]--;
        if(freq_numbers[max_freq - 1].size() == 0){
            auto it = freq_numbers.end();
            it--;
            freq_numbers.erase(it);
            max_freq--;
        }
        return res;
    }
};


//TLE
//class FreqStack {
//    std::unordered_map<int,std::list<std::list<int>::iterator>> num_it;
//    std::map<int,std::unordered_set<int>> cnt_num_list;
//    std::list<int> sequence;
//public:
//    FreqStack() {
//
//    }
//
//    void push(int val){
//        sequence.push_back(val);
//        auto it = sequence.end();
//        it--;
//        if(num_it.find(val) == num_it.end()){
//            cnt_num_list[1].insert(val);
//        }else{
//            int pre_cnt = num_it[val].size();
//            cnt_num_list[pre_cnt].erase(val);
//            cnt_num_list[pre_cnt + 1].insert(val);
//        }
//        num_it[val].push_back(it);
//    }
//
//    int pop() {
//        std::unordered_set<int> it = (*cnt_num_list.rbegin()).second;
//        int pre_cnt = (*cnt_num_list.rbegin()).first;
//        int min_dis = 2147483647;
//        std::list<int>::iterator find_it = sequence.end();
//        for(auto it2 = it.begin();it2 != it.end();it2++){
//            int cur_dis = std::distance(*num_it[*it2].rbegin(),sequence.end());
//            if(cur_dis < min_dis){
//                min_dis = cur_dis;
//                find_it = *num_it[*it2].rbegin();
//            }
//        }
//        int res = *find_it;
//        auto del = num_it[res].end();
//        del--;
//        num_it[res].erase(del);
//        if(num_it[res].size() == 0){
//            num_it.erase(res);
//        }
//        sequence.erase(find_it);
//        cnt_num_list[pre_cnt].erase(res);
//        if(cnt_num_list[pre_cnt].size() == 0){
//            cnt_num_list.erase(pre_cnt);
//        }
//        if(pre_cnt > 1){
//            cnt_num_list[pre_cnt - 1].insert(res);
//        }
//        return res;
//    }
//};

//TLE
//class FreqStack {
//    std::unordered_map<int,std::list<std::list<int>::iterator>> num_it;
//    std::map<int,std::unordered_set<int>> cnt_num_list;
//    std::list<int> sequence;
//public:
//    FreqStack() {
//
//    }
//
//    void push(int val){
//        sequence.push_back(val);
//        auto it = sequence.end();
//        it--;
//        if(num_it.find(val) == num_it.end()){
//            cnt_num_list[1].insert(val);
//        }else{
//            int pre_cnt = num_it[val].size();
//            cnt_num_list[pre_cnt].erase(val);
//            cnt_num_list[pre_cnt + 1].insert(val);
//        }
//        num_it[val].push_back(it);
//    }
//
//    int pop() {
//        std::unordered_set<int> it = (*cnt_num_list.rbegin()).second;
//        int pre_cnt = (*cnt_num_list.rbegin()).first;
//        int res = -1;
//        for(auto it2 = sequence.rbegin();it2 != sequence.rend();it2++){
//            if(it.find(*it2) != it.end()){
//                res = *it2;
//                std::list<std::list<int>::iterator>::iterator it3 = num_it[res].end();
//                //std::list<int>::iterator it3 = l2.end();
//                it3--;
//                sequence.erase(*it3);
//                num_it[res].erase(it3);//delete last iterator
//                cnt_num_list[pre_cnt].erase(res);
//                if(cnt_num_list[pre_cnt].size() == 0){
//                    cnt_num_list.erase(pre_cnt);
//                }
//                if(pre_cnt > 1){
//                    cnt_num_list[pre_cnt - 1].insert(res);
//                }
//                if(num_it[res].size() == 0){
//                    num_it.erase(res);
//                }
//                break;
//            }
//        }
//        return res;
//    }
//};