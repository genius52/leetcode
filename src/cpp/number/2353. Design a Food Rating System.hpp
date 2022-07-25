#include <vector>
#include <string>
#include <map>
#include <unordered_map>


class FoodRatings {
    class pair_comp {
    public:
        bool operator() (const std::pair<int, string>& a, const std::pair<int, string>& b)const {
//            if (A.first !=  B.first) {
//                return A.first > B.first;
//            }else {
//                return A.second.compare(B.second) == -1;
//            }
            return (a.first == b.first ? a.second < b.second : a.first > b.first);
        }
    };
    std::unordered_map<std::string,std::pair<std::string,int>> food_;
    std::unordered_map<std::string,std::set<std::pair<int,std::string>,pair_comp>> kind_;
public:
    FoodRatings(vector<string>& foods, vector<string>& cuisines, vector<int>& ratings) {
        int len = foods.size();
        for(int i = 0;i < len;i++){
            food_[foods[i]] = {cuisines[i],ratings[i]};
            //std::pair<int,string> p = {ratings[i],foods[i]};
            //kind_[cuisines[i]].insert({ratings[i],foods[i]});
            kind_[cuisines[i]].insert(std::pair<int,std::string>(ratings[i], foods[i]));
        }
    }

    void changeRating(string food, int newRating) {
        auto old_rate = food_[food].second;
        auto type = food_[food].first;
        kind_[type].erase({old_rate,food});
        kind_[type].insert({newRating,food});
        food_[food].second = newRating;
    }

    string highestRated(string cuisine) {
        auto it = kind_[cuisine].begin();
        return it->second;
    }
};