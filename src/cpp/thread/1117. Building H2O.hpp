#include <functional>
#include <mutex>
#include <condition_variable>
#include <atomic>
using namespace std;

//Input: "OOHHHH"
//Output: "HHOHHO"


//"HOH"
class H2O {
    std::mutex mtx_;
    std::condition_variable cv_;
    std::atomic<int> h_count_ = 0;
    std::atomic<int> o_count_ = 0;
public:
    H2O() {

    }

    void hydrogen(function<void()> releaseHydrogen) {
        std::unique_lock lk(mtx_);
        cv_.wait(lk,[this](){
            return h_count_ <= (o_count_ * 2);
        });
        releaseHydrogen();
        h_count_++;
        cv_.notify_all();
    }

    void oxygen(function<void()> releaseOxygen) {
        std::unique_lock lk(mtx_);
        cv_.wait(lk,[this](){
            return (o_count_ * 2) <= h_count_;
        });
        releaseOxygen();
        o_count_++;
        cv_.notify_all();
    }
};