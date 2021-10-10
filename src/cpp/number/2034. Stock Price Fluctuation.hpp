#include <map>

class StockPrice {
    std::map<int,int> time_price_;
    std::multimap<int,int> price_time_;
public:
    StockPrice() {

    }

    void update(int timestamp, int price) {
        if(time_price_.find(timestamp) != time_price_.end()){
            auto range = price_time_.equal_range(time_price_[timestamp]);
            while (range.first!=range.second)
            {
                if (range.first->second == timestamp)
                {
                    price_time_.erase(range.first);
                    break;
                }
                ++range.first;//不进行删除操作则自增
            }
        }
        time_price_[timestamp] = price;
        price_time_.insert({price,timestamp});
    }

    int current() {
        return time_price_.rbegin()->second;
    }

    int maximum() {
        return price_time_.rbegin()->first;
    }

    int minimum() {
        return price_time_.begin()->first;
    }
};