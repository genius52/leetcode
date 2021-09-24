#include <string>
#include <unordered_map>
#include <set>
using namespace std;

class TweetCounts {
    std::unordered_map<std::string,std::multiset<int>> record_;
public:
    TweetCounts() {

    }

    void recordTweet(string tweetName, int time) {
        record_[tweetName].insert(time);
    }

    vector<int> getTweetCountsPerFrequency(string freq, string tweetName, int startTime, int endTime) {
        int interval = 0;
        if(freq == "minute"){
            interval = 60;
        }else if(freq == "hour"){
            interval = 3600;
        }else if(freq == "day"){
            interval = 86400;
        }
        std::vector<int> res(((endTime - startTime) / interval) + 1);
        if(record_.find(tweetName) == record_.end())
            return res;
        for(auto it = record_[tweetName].begin();it != record_[tweetName].end();it++){
            if((*it) >= startTime && (*it) <= endTime){
                res[((*it) - startTime)/interval]++;
            }
        }
//        std::multiset<int>::iterator pre = record_[tweetName].lower_bound(startTime);
//        int cur_start = startTime;
//        int idx = 0;
//        while(cur_start <= endTime){
//            cur_start += interval;
//            auto find = record_[tweetName].lower_bound(cur_start);
//            int cnt = std::distance(pre,find);
//            //res.push_back(cnt);
//            res[idx] = cnt;
//            idx++;
//            pre = find;
//        }
        return res;
    }
};