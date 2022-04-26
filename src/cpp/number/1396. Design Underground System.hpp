#include <unordered_map>
#include <string>
#include <algorithm>

class UndergroundSystem {
    std::unordered_map<int, std::pair<std::string, int> > checkin_;
    std::unordered_map<std::string,std::unordered_map<std::string,std::pair<int,int>>> record_;
public:
    UndergroundSystem() {

    }

    void checkIn(int id, string stationName, int t) {
        checkin_[id] = {stationName,t};
    }

    void checkOut(int id, string stationName, int t) {
        auto duration = t - checkin_[id].second;
        auto pre = record_[checkin_[id].first][stationName];
        record_[checkin_[id].first][stationName].first += duration;
        record_[checkin_[id].first][stationName].second++;
    }

    double getAverageTime(string startStation, string endStation) {
        return double(record_[startStation][endStation].first) / double(record_[startStation][endStation].second);
    }
};

//class UndergroundSystem {
//public:
//    std::unordered_map<int, std::pair<std::string, int> > checkin_map;
//    std::unordered_map<string, std::vector< int> > checkout_map;
//    UndergroundSystem() {
//
//    }
//
//    void checkIn(int id, string stationName, int t) {
//        std::pair<std::string, int> p{ stationName, t };
//        checkin_map[id] = p;
//    }
//
//    void checkOut(int id, string stationName, int t) {
//        std::string start_end = checkin_map[id].first  + ","  + stationName;
//        int duration = t - checkin_map[id].second;
//        if (checkout_map.find(start_end) != checkout_map.end()) {
//            checkout_map[start_end].push_back(duration);
//        }
//        else {
//            std::vector<int> v{ duration };
//            checkout_map[start_end] = v;
//        }
//        checkin_map.erase(id);
//    }
//
//    double getAverageTime(string startStation, string endStation) {
//        int total = 0;
//        std::string start_end = startStation + "," + endStation;
//        if (checkout_map.find(start_end) != checkout_map.end()) {
//            for (int i = 0; i < checkout_map[start_end].size(); i++) {
//                total += checkout_map[start_end][i];
//            }
//        }
//        else {
//            return 0;
//        }
//        double d = double(total) / double(checkout_map[start_end].size());
//        return d;
//    }
//};