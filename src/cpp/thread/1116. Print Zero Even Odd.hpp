#include <functional>
#include <mutex>
#include <condition_variable>
using namespace std;

//Input: n = 5
//Output: "0102030405"
class ZeroEvenOdd {
private:
    int n;
    std::mutex m_;
    std::condition_variable cond_;
    int cur = 1;
    bool is_zero = true;
    bool end = false;
public:
    ZeroEvenOdd(int n) {
        this->n = n;
    }

    // printNumber(x) outputs "x", where x is an integer.
    void zero(function<void(int)> printNumber) {
        while(cur <= n){
            std::unique_lock<std::mutex> lk(m_);
            cond_.wait(lk,[this](){
                return end || is_zero;
            });
            if(cur > n){
                end = true;
                cond_.notify_all();
                break;
            }
            printNumber(0);
            //std::cout<<"0";
            is_zero = false;
            //cur++;
            cond_.notify_one();
        }
    }

    void even(function<void(int)> printNumber) {
        while(cur <= n) {
            std::unique_lock<std::mutex> lk(m_);
            cond_.wait(lk, [this]() {
                return end || (!is_zero && (cur | 1) != cur);
            });
            if(cur > n){
                end = true;
                cond_.notify_all();
                break;
            }

            printNumber(cur);
            //std::cout<<cur;
            is_zero = true;
            cur++;
            cond_.notify_one();
        }
    }

    void odd(function<void(int)> printNumber) {
        while(cur <= n) {
            std::unique_lock<std::mutex> lk(m_);
            cond_.wait(lk, [this]() {
                return end || (!is_zero && (cur | 1) == cur);
            });
            if(cur > n){
                end = true;
                cond_.notify_all();
                break;
            }

            printNumber(cur);
            //std::cout<<cur;
            is_zero = true;
            cur++;
            cond_.notify_one();
        }
    }
};