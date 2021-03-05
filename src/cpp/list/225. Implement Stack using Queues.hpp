#include <queue>

class MyStack {
public:
    std::queue<int> q;
    std::queue<int> q2;
    /** Initialize your data structure here. */
    MyStack() {

    }

    /** Push element x onto stack. */
    void push(int x) {
        q.push(x);
    }

    /** Removes the element on top of the stack and returns that element. */
    int pop() {
        if(q.empty())
            return 0;
        int size = q.size();
        while(size > 1){
            q2.push(q.front());
            q.pop();
            size--;
        }
        int val = q.front();
        q.pop();
        while(!q2.empty()){
            q.push(q2.front());
            q2.pop();
        }
        return val;
    }

    /** Get the top element. */
    int top() {
        if(q.empty())
            return 0;
        int size = q.size();
        while(size > 1){
            q2.push(q.front());
            q.pop();
            size--;
        }
        int val = q.front();
        q2.push(val);
        q.pop();
        while(!q2.empty()){
            q.push(q2.front());
            q2.pop();
        }
        return val;
    }

    /** Returns whether the stack is empty. */
    bool empty() {
        return q.empty();
    }
};