#include <vector>
#include <map>
#include <unordered_map>
#include <set>
#include <queue>
using namespace std;

//entries[i] = [shopi, moviei, pricei]
class MovieRentingSystem {
    //
    //std::unordered_map<int,std::unordered_map<int,int>> rent_shop_;//shop:movie:price
    std::set<std::pair<int,int>> unrent_movie_[10001];//当前商店可用的电影 movie:shop,price;
    std::map<std::pair<int,int>,int> fix_;// shop:movie:price

    //std::map<int,std::set<std::pair<int,int>>> rent_movie_; //借出去的电影，price:shop,movie
    set<pair<int, pair<int, int>>> rent_movie_;
    //set<pair<int, pair<int, int>>> rt;
public:
    MovieRentingSystem(int n, vector<vector<int>>& entries) {
        for(auto entry : entries){
            fix_[std::make_pair(entry[0],entry[1])] = entry[2];
            unrent_movie_[entry[1]].insert(std::make_pair(entry[2],entry[0]));
        }
    }
    //找最便宜的没借出的5家店
    vector<int> search(int movie) {
        std::vector<int> res;
        int cnt = 0;
        for(auto it = unrent_movie_[movie].begin();it != unrent_movie_[movie].end() && cnt < 5;it++){
            res.push_back(it->second);
            cnt++;
        }
        return res;
    }
    //从商店shop借电影movie
    void rent(int shop, int movie) {
        int price = fix_[std::make_pair(shop,movie)];
        unrent_movie_[movie].erase(std::make_pair(price,shop));
        //rent_movie_[price].insert(std::make_pair(shop,movie));
        rent_movie_.insert(std::make_pair(price,std::make_pair(shop,movie)));
    }
    //归还
    void drop(int shop, int movie) {
        int price = fix_[std::make_pair(shop,movie)];
        //rent_movie_[price].erase(std::make_pair(shop,movie));
        rent_movie_.erase(std::make_pair(price,std::make_pair(shop,movie)));
        unrent_movie_[movie].insert(std::make_pair(price,shop));
    }
    //最便宜的 5 部已借出电影
    vector<vector<int>> report() {
        std::vector<std::vector<int>> res;
        int cnt = 0;
        for(auto it = rent_movie_.begin();it != rent_movie_.end() && cnt < 5;it++){
            res.push_back({it->second.first,it->second.second});
            cnt++;
        }
        return res;
    }
};