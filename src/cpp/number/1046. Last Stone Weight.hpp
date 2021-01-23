#include <set>
#include <vector>
using namespace std;

int lastStoneWeight(std::vector<int>& stones) {
    std::multiset<int,std::greater<int>> s;
    for(auto it = stones.begin();it != stones.end();it++){
        s.insert(*it);
    }
    while(s.size()>0){
        auto first = s.begin();
        auto biggest = *first;
        s.erase(first++);
        if(first == s.end())
            return biggest;
        else{
            if(*first != biggest)
                s.insert(biggest - *first);
            s.erase(first);
        }
    }
    return 0;
}