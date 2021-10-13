#include <string>
#include <map>
#include <vector>
using namespace std;

class Solution_1418 {
public:
    class SortId
    {
    public:
        bool operator () (const std::string& id1, const std::string& id2) const
        {
            int l1 = id1.length();
            int l2 = id2.length();
            if (l1 < l2) {
                return true;
            }
            if (l1 > l2) {
                return false;
            }
            return id1.compare(id2) < 0;
        }
    };

    typedef std::map < string, std::map<string, int>,SortId> my_map;
    vector<vector<string>> displayTable(vector<vector<string>>& orders) {
        std::vector<std::vector<string>> res;
        my_map record;
        std::set<string> foods;
        int len = orders.size();
        for (int i = 0; i < len; i++) {
            foods.insert(orders[i][2]);
            if (record.find(orders[i][1]) != record.end()) {
                if (record[orders[i][1]].find(orders[i][2]) != record[orders[i][1]].end()) {
                    record[orders[i][1]][orders[i][2]]++;
                }
                else {
                    record[orders[i][1]][orders[i][2]] = 1;
                }
            }
            else {
                std::map<string, int> m;
                m[orders[i][2]] = 1;
                record[orders[i][1]] = m;
            }
        }
        int food_cnt = foods.size();
        std::vector<string> title;
        title.push_back("Table");
        for (auto it_food = foods.begin(); it_food != foods.end(); it_food++)
            title.push_back(*it_food);
        res.push_back(title);
        for (auto it = record.begin(); it != record.end(); it++) {
            std::vector<string> row;// (food_cnt);
            row.push_back(it->first);
            for (auto it_food = foods.begin(); it_food != foods.end(); it_food++) {
                if (it->second.find(*it_food) == it->second.end()) {
                    row.push_back("0");
                    continue;
                }
                row.push_back(std::to_string(it->second[*it_food]));
            }
            res.push_back(row);
        }
        return res;
    }
};