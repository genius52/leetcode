#include <mutex>
#include <condition_variable>

//1195
//If the number is divisible by 3, output "fizz".
// If the number is divisible by 5, output "buzz".
// If the number is divisible by both 3 and 5, output "fizzbuzz".
// For example, for n = 15, we output: 1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz.
//Implement a multithreaded version of FizzBuzz with four threads. The same instance of FizzBuzz will be passed to four different threads:

// Thread A will call fizz() to check for divisibility of 3 and outputs fizz.
// Thread B will call buzz() to check for divisibility of 5 and outputs buzz.
// Thread C will call fizzbuzz() to check for divisibility of 3 and 5 and outputs fizzbuzz.
// Thread D will call number() which should only output the numbers.
class FizzBuzz {
private:
    int n;
    int cur_num = 1;
    std::mutex mtx;
    std::condition_variable cv;
public:
    FizzBuzz(int n) {
        this->n = n;
        cv.notify_all();
    }

    // printFizz() outputs "fizz".
    void fizz(function<void()> printFizz) {
        while (cur_num <= n){
            std::unique_lock<mutex> lock(mtx);
            cv.wait(lock, [this](){
                return this->cur_num > n ||(this->cur_num % 3 == 0 && this->cur_num % 5 != 0);
            });
            if(this->cur_num > n)
                break;
            printFizz();
            this->cur_num++;
            cv.notify_all();
        }
    }

    // printBuzz() outputs "buzz".
    void buzz(function<void()> printBuzz) {
        while(cur_num <= n){
            std::unique_lock<mutex> lock(mtx);
            cv.wait(lock, [this](){
                return this->cur_num > n || (this->cur_num % 3 != 0 && this->cur_num % 5 == 0);
            });
            if(this->cur_num > n)
                break;
            printBuzz();
            this->cur_num++;
            cv.notify_all();
        }
    }

    // printFizzBuzz() outputs "fizzbuzz".
    void fizzbuzz(function<void()> printFizzBuzz) {
        while(cur_num <= n){
            std::unique_lock<mutex> lock(mtx);
            cv.wait(lock,[this](){
                return this->cur_num > n || (this->cur_num % 3 == 0 && this->cur_num % 5 == 0);
            });
            if(this->cur_num > n)
                break;
            printFizzBuzz();
            this->cur_num++;
            cv.notify_all();
        }
    }

    // printNumber(x) outputs "x", where x is an integer.
    void number(function<void(int)> printNumber) {
        while(cur_num <= n){
            std::unique_lock<mutex> lock(mtx);
            cv.wait(lock,[this](){
                return this->cur_num > n ||(this->cur_num % 3 != 0 && this->cur_num % 5 != 0);
            });
            if(this->cur_num > n)
                break;
            printNumber(cur_num);
            this->cur_num++;
            cv.notify_all();
        }
    }
};
