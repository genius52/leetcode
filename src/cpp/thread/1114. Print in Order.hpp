#include <mutex>
#include <condition_variable>
using namespace std;

class Foo {
    std::mutex mtx_;
    std::condition_variable cv_;
    int count_ = 0;
public:
    Foo() {
    }

    void first(function<void()> printFirst) {

        // printFirst() outputs "first". Do not change or remove this line.
        printFirst();
        count_ = 1;
        cv_.notify_all();
    }

    void second(function<void()> printSecond) {
        std::unique_lock <std::mutex> lck(mtx_);
        cv_.wait(lck, [this]() {
            return count_ == 1;
        });
        // printSecond() outputs "second". Do not change or remove this line.
        printSecond();
        count_ = 2;
        cv_.notify_one();
    }

    void third(function<void()> printThird) {
        std::unique_lock <std::mutex> lck(mtx_);
        cv_.wait(lck, [this]() {
            return count_ == 2;
        });
        // printThird() outputs "third". Do not change or remove this line.
        printThird();
        cv_.notify_one();
    }
};