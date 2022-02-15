#include <vector>
using namespace std;

class Bank {
    vector<long long> balance_;
public:
    Bank(vector<long long>& balance) {
        balance_ = balance;
    }

    bool transfer(int account1, int account2, long long money) {
        if(account1 > balance_.size() || account2 > balance_.size())
            return false;
        if(balance_[account1 - 1] < money)
            return false;
        balance_[account1 - 1] -= money;
        balance_[account2 - 1] += money;
        return true;
    }

    bool deposit(int account, long long money) {
        if(account > balance_.size())
            return false;
        balance_[account - 1] += money;
        return true;
    }

    bool withdraw(int account, long long money) {
        if(account > balance_.size())
            return false;
        if(balance_[account - 1] < money)
            return false;
        balance_[account - 1] -= money;
        return true;
    }
};