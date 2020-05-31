//
// Created by 陶诚程 on 2020-05-16.
//
#include <vector>
#include <string>
#include <map>
#include <algorithm>
using namespace std;

class Solution_332 {
public:
    void dfs_findItinerary(std::string start,std::map<std::string,std::vector<std::string>>& record,std::vector<std::string>& res){
        if(record.empty())
            return;
        while(record.find(start) != record.end()){
            std::vector<std::string> cities = record[start];
            auto last_start = start;
            start = cities[0];
            if(cities.size() == 1)
                record.erase(last_start);
            else{
                cities.erase(std::find(cities.begin(),cities.end(),cities[0]));
                record[last_start] = cities;
            }
            res.push_back(start);
        }
        if(record.empty())
            return;
        auto it = record.begin();
        dfs_findItinerary((*it).first,record,res);
    }

    vector<string> findItinerary(vector<vector<string>>& tickets) {
        std::vector<std::string> res;
        std::map<std::string,std::vector<std::string>> record;
        int len = tickets.size();
        for(int i = 0;i < len;i++){
            auto start_city = tickets[i][0];
            auto end_city = tickets[i][1];
            if(record.find(start_city) != record.end()){
                record[start_city].push_back(end_city);
                std::sort(record[start_city].begin(),record[start_city].end());
            }else{
                record[start_city] = std::vector<std::string>{end_city};
            }
        }
        std::string start = "JFK";
        res.push_back(start);
        dfs_findItinerary(start,record,res);
        return res;
    }
};