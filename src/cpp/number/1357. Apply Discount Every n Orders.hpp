#include <vector>
#include <unordered_map>
using namespace std;

class Cashier {
    std::unordered_map<int,int> product_price_;
    double discount_ = 0;
    int cur_ = 0;
    int n_ = 0;
public:
    Cashier(int n, int discount, vector<int>& products, vector<int>& prices) {
        int len = products.size();
        for(int i = 0;i < len;i++){
            product_price_[products[i]] = prices[i];
        }
        discount_ = discount;
        n_ = n;
    }

    double getBill(vector<int> product, vector<int> amount) {
        cur_++;
        double res = 0;
        int len = product.size();
        for(int i = 0;i < len;i++){
            res += product_price_[product[i]] * amount[i];
        }
        if(cur_ % n_ == 0){
            res = res * (100 - discount_) / 100;
        }
        return res;
    }
};