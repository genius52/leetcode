#include <stack>
#include <deque>
using namespace std;

class MinStack {
    std::stack<int,std::deque<int>> data_stack;
    std::stack<int,std::deque<int>> min_stack;
public:
    /** initialize your data structure here. */
    MinStack() {

    }

    void push(int x) {
        data_stack.push(x);
        if(min_stack.empty()){
            min_stack.push(x);
        }else{
            int top = min_stack.top();
            if(top < x){
                min_stack.push(top);
            }else{
                min_stack.push(x);
            }
        }
    }

    void pop() {
        data_stack.pop();
        min_stack.pop();
    }

    int top() {
        return data_stack.top();
    }

    int getMin() {
        return min_stack.top();
    }
};

class MinStack2 {
    std::stack<std::pair<int,int>> s_;//当前值和最小值
public:
    /** initialize your data structure here. */
    MinStack() {

    }

    void push(int x) {
        if(s_.empty()){
            s_.push({x,x});
        }else{
            auto pre_min = s_.top().second;
            if(x < pre_min){
                s_.push({x,x});
            }else{
                s_.push({x,pre_min});
            }
        }
    }

    void pop() {
        if(!s_.empty()){
            s_.pop();
        }
    }

    int top() {
        return s_.top().first;
    }

    int getMin() {
        return s_.top().second;
    }
};