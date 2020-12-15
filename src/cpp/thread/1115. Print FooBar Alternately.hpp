#include <functional>
#include <thread>
#include <mutex>
#include <condition_variable>


class FooBar {
private:
    int n;
    std::mutex allow_;
    std::condition_variable cond_;
    static bool start;
public:
    FooBar(int n) {
        this->n = n;
    }

    void foo(function<void()> printFoo) {
        for (int i = 0; i < n; i++) {
            std::unique_lock<std::mutex> lk(allow_);
            cond_.wait(lk,[this](){
                return start;
            });
            printFoo();
            start = false;
            cond_.notify_one();
        }
    }

    void bar(function<void()> printBar) {
        for (int i = 0; i < n; i++) {
            std::unique_lock<std::mutex> lk(allow_);
            cond_.wait(lk,[this](){
                return !start;
            });
            printBar();
            start = true;
            cond_.notify_one();
        }
    }
};

bool FooBar::start = true;