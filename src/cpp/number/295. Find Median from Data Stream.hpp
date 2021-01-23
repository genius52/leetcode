#include <queue>
#include <vector>
#include <algorithm>

class MedianFinder {
public:
    /** initialize your data structure here. */
    MedianFinder() {
    }

    void addNum(int num) {
        auto small_size = smalltop_queue.size();
        auto big_size = bigtop_queue.size();
        if (small_size == 0){
            smalltop_queue.push(num);
            return;
        }
        if (small_size > big_size){
            if (num > smalltop_queue.top()){
                bigtop_queue.push(smalltop_queue.top());
                smalltop_queue.pop();
                smalltop_queue.push(num);
            }else{
                bigtop_queue.push(num);
            }
        }else if (small_size < big_size){
            if (num < bigtop_queue.top()){
                smalltop_queue.push(bigtop_queue.top());
                bigtop_queue.pop();
                bigtop_queue.push(num);
            }else{
                smalltop_queue.push(num);
            }
        }else{
            if (num <= bigtop_queue.top()){
                bigtop_queue.push(num);
            }else{
                smalltop_queue.push(num);
            }
        }
    }

    double findMedian() {
        if (smalltop_queue.size() == bigtop_queue.size()){
            return double(smalltop_queue.top() + bigtop_queue.top()) / 2;
        }else if(smalltop_queue.size() > bigtop_queue.size()){
            return smalltop_queue.top();
        }else{
            return bigtop_queue.top();
        }
    }
    std::priority_queue<int,std::vector<int>,std::greater<int> > smalltop_queue;
    std::priority_queue<int,std::vector<int>,std::less<int> > bigtop_queue;
};