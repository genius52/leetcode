
class Bitset {
    std::string data_;
    int one_cnt_ = 0;
    bool is_flip_ = false;
public:
    Bitset(int size) {
        data_.resize(size);
        for(int i = 0;i < size;i++){
            data_[i] = '0';
        }
    }

    void fix(int idx) {
        if(data_[idx] == '0' && !is_flip_ ){
            data_[idx] = '1';
            one_cnt_++;
        }else if(data_[idx] == '1' && is_flip_ ){
            data_[idx] = '0';
            one_cnt_++;
        }
    }

    void unfix(int idx) {
        if(data_[idx] == '1' && !is_flip_){
            data_[idx] = '0';
            one_cnt_--;
        }else if(data_[idx] == '0' && is_flip_){
            data_[idx] = '1';
            one_cnt_--;
        }
    }

    void flip() {
        is_flip_ = !is_flip_;
        one_cnt_ = data_.length() - one_cnt_;
    }

    bool all() {
        return one_cnt_ == data_.length();
    }

    bool one() {
        return one_cnt_ > 0;
    }

    int count() {
        return one_cnt_;
    }

    string toString() {
        if (is_flip_){
            std::string res = data_;
            for (int i = 0; i < res.length(); ++i) {
                if (res[i] == '0')
                    res[i] = '1';
                else
                    res[i] = '0';
            }
            return res;
        }else{
            return data_;
        }
    }
};

/**
 * Your Bitset object will be instantiated and called as such:
 * Bitset* obj = new Bitset(size);
 * obj->fix(idx);
 * obj->unfix(idx);
 * obj->flip();
 * bool param_4 = obj->all();
 * bool param_5 = obj->one();
 * int param_6 = obj->count();
 * string param_7 = obj->toString();
 */