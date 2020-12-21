#include <vector>
#include <map>
using namespace std;


class SnapshotArray {
    std::vector<std::map<int,int,std::greater<int>>> record;
    int snap_id = 0;
    int len = 0;
public:
    SnapshotArray(int length) {
        record.resize(length);
        len = length;
    }

    void set(int index, int val) {
        record[index][snap_id] = val;
    }

    int snap() {
        return snap_id++;
    }

    int get(int index, int snap_id) {
        auto it = record[index].lower_bound(snap_id);
        if (it == record[index].end())
            return 0;
        return it->second;
    }
};

//class SnapshotArray {
//    //std::vector<std::vector<int>> record;
//    std::vector<std::pair<int,int>> record;
//    int snap_id = -1;
//    int len = 0;
//public:
//    SnapshotArray(int length) {
//        record.resize(length);
//        len = length;
//    }
//
//    void set(int index, int val) {
//        if(record[index].size() > 0){
//            record[index][0] = val;
//        }else{
//            record[index].push_back(val);
//        }
//    }
//
//    int snap() {
//        for(int i = 0;i < len;i++){
//            if (record[i].size() > 0){
//                record[i].push_back(record[i][0]);
//            }
//        }
//        snap_id++;
//        return snap_id;
//    }
//
//    int get(int index, int snap_id) {
//        auto res = record[index][snap_id];
//        return res;
//    }
//};