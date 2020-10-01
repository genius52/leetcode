#include <vector>
#include <deque>
using namespace std;


class Solution_735 {
public:
    vector<int> asteroidCollision(vector<int>& asteroids) {
        int len = asteroids.size();
        std::deque<int> q;
        for (int i = 0; i < len; ++i) {
            if(q.empty()){
                q.push_front(asteroids[i]);
            }else{
                while(!q.empty()){
                    auto top = q.front();
                    if( top > 0 && asteroids[i] < 0 ){
                        if(abs(top) < abs(asteroids[i])){
                            q.pop_front();
                            if(q.empty()){
                                q.push_front(asteroids[i]);
                                break;
                            }
                        }else if(top == -asteroids[i]){
                            q.pop_front();
                            break;
                        }
                        else{
                            break;
                        }
                    }else{
                        q.push_front(asteroids[i]);
                        break;
                    }
                }
            }
        }
        std::vector<int> res;
        while(!q.empty()){
            res.push_back(q.back());
            q.pop_back();
        }
        return res;
    }
};