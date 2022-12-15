#include <unordered_map>
#include <unordered_set>
#include <set>
using namespace std;

//struct pair_hash
//{
//    template <class T1, class T2>
//    size_t operator () (pair<T1, T2> const &pair) const
//    {
//        size_t h1 = hash<T1>()(pair.first); //用默认的 hash 处理 pair 中的第一个数据 X1
//        size_t h2 = hash<T2>()(pair.second);//用默认的 hash 处理 pair 中的第二个数据 X2
//        return h1 ^ h2;
//    }
//};

class Allocator {
    int len_ = 0;
    std::vector<int> data_;
    std::unordered_map<int,std::unordered_map<int,int>> id_start_size_;
public:
    Allocator(int n) {
        data_.resize(n);
        len_ = n;
    }

    int allocate(int size, int mID) {
        int res = -1;
        int left = 0;
        while(left + size <= len_){
            while(left + size <= len_ && data_[left] != 0)
                left++;
            int right = left;
            while(right < len_ && data_[right] == 0)
                right++;
            if(right - left >= size){
                id_start_size_[mID][left] = size;
                for(int i = 0;i < size;i++)
                    data_[left + i] = 1;
                return left;
            }
            left = right;
        }
        return res;
    }

    int free(int mID) {
        if(id_start_size_.find(mID) == id_start_size_.end())
            return 0;
        int res = 0;
        for(auto it = id_start_size_[mID].begin();it != id_start_size_[mID].end();it++){
            int start = (*it).first;
            int size = (*it).second;
            res += size;
            for(int i = 0;i < size;i++){
                data_[start + i] = 0;
            }
        }
        id_start_size_.erase(mID);
        return res;
    }
};