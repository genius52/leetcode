#include <string>
#include <map>
#include <stack>
#include <deque>
#include <vector>
using namespace std;

class Solution_67 {
    std::map<char,char> tags_;
    std::stack<char> stacks_;
public:
    Solution()
    {
        tags_.insert(std::pair<char,char>('{','}'));
        tags_.insert(std::pair<char,char>('[',']'));
        tags_.insert(std::pair<char,char>('(',')'));
    }
    bool isValid(string s) {
        if(s.empty())
            return true;
        if(s.length() == 1)
            return false;
        auto head = s.at(0);
        auto head_next = s.at(1);
        if(tags_[head] == head_next)
        {
            if(s.length() > 2)
                return isValid(s.substr(2,s.length()));
            else
                return true;
        }

        auto tail = s.at(s.length()-1);
        auto tail_last = s.at(s.length()-2);
        if(tags_[head] == tail)
            return isValid(s.substr(1,s.length()-2));
        if(tags_[tail_last] == tail)
        {
            if(s.length() > 2)
                return isValid(s.substr(0,s.length()-2));
            else
                return true;
        }
        return false;
    }

    string addBinary(string a, string b) {
        int len_a = a.length();
        if(len_a == 0)
            return b;
        int len_b = b.length();
        if(len_b == 0)
            return a;
        std::string sum;
        int index = 0;
        bool upgrade = false;
        int min_len = (len_a < len_b)?len_a:len_b;
        while(index < min_len){
            if(a.at(len_a - 1 - index) == '0' && b.at(len_b - 1 - index) == '0')
            {
                if(upgrade)
                    sum = "1" + sum;
                else
                    sum = "0" + sum;
                upgrade = false;
            }
            else if ((a.at(len_a - 1 - index) == '1' && b.at(len_b - 1 - index) == '0')||
                     (a.at(len_a - 1 - index) == '0' && b.at(len_b - 1 - index) == '1'))
            {
                if(upgrade)
                {
                    sum = "0" + sum;
                    upgrade = true;
                }
                else
                {
                    sum = "1" + sum;
                    upgrade = false;
                }
            }
            else
            {
                if(upgrade)
                {
                    sum = "1" + sum;
                    upgrade = true;
                }
                else
                {
                    sum = "0" + sum;
                    upgrade = true;
                }

            }
            index++;
        }
        while(index < len_b)
        {
            if(b.at(len_b - 1 - index) == '0')
            {
                if(upgrade)
                    sum = "1" + sum;
                else
                    sum = "0" + sum;
                upgrade = false;
            }
            else if (b.at(len_b - 1 - index) == '1')
            {
                if(upgrade)
                {
                    sum = "0" + sum;
                    upgrade = true;
                }
                else
                {
                    sum = "1" + sum;
                    upgrade = false;
                }
            }
            index++;
        }
        while(index < len_a)
        {
            if(a.at(len_a - 1 - index) == '0')
            {
                if(upgrade)
                    sum = "1" + sum;
                else
                    sum = "0" + sum;
                upgrade = false;
            }
            else if (a.at(len_a - 1 - index) == '1')
            {
                if(upgrade)
                {
                    sum = "0" + sum;
                    upgrade = true;
                }
                else
                {
                    sum = "1" + sum;
                    upgrade = false;
                }
            }
            index++;
        }
        if(upgrade)
            sum = "1" + sum;
        return sum;
    }
};
