#include <functional>
#include <thread>
#include <mutex>
#include <condition_variable>


class FooBar {
private:
    int n;
    std::mutex allow_;
    std::condition_variable cond_;
    bool start_ = true;
public:
    FooBar(int n) {
        this->n = n;
    }

    void foo(function<void()> printFoo) {
        for (int i = 0; i < n; i++) {
            std::unique_lock<std::mutex> lk(allow_);
            cond_.wait(lk,[this](){
                return start_;
            });
            printFoo();
            start_ = false;
            cond_.notify_one();
        }
    }

    void bar(function<void()> printBar) {
        for (int i = 0; i < n; i++) {
            std::unique_lock<std::mutex> lk(allow_);
            cond_.wait(lk,[this](){
                return !start_;
            });
            printBar();
            start_ = true;
            cond_.notify_one();
        }
    }
};