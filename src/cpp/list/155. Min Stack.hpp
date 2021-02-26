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

