#include <functional>
#include <mutex>
#include <condition_variable>
#include <atomic>
using namespace std;

//Input: "OOHHHH"
//Output: "HHOHHO"
class H2O {
public:
    H2O() {

    }

    void hydrogen(function<void()> releaseHydrogen) {
        // releaseHydrogen() outputs "H". Do not change or remove this line.
        releaseHydrogen();
    }

    void oxygen(function<void()> releaseOxygen) {
        // releaseOxygen() outputs "O". Do not change or remove this line.
        releaseOxygen();
    }
};