#include <stack>

class MyQueue {
    std::stack<int> s1;
    std::stack<int> s2;
public:
    /** Initialize your data structure here. */
    MyQueue() {

    }

    /** Push element x to the back of queue. */
    void push(int x) {
        s1.push(x);
    }

    /** Removes the element from in front of queue and returns that element. */
    int pop() {
        if(s1.empty())
            return 0;
        while(!s1.empty()){
            s2.push(s1.top());
            s1.pop();
        }
        int val = s2.top();
        s2.pop();
        while(!s2.empty()){
            s1.push(s2.top());
            s2.pop();
        }
        return val;
    }

    /** Get the front element. */
    int peek() {
        if(s1.empty())
            return 0;
        while(!s1.empty()){
            s2.push(s1.top());
            s1.pop();
        }
        int val = s2.top();
        while(!s2.empty()){
            s1.push(s2.top());
            s2.pop();
        }
        return val;
    }

    /** Returns whether the queue is empty. */
    bool empty() {
        return s1.empty();
    }
};