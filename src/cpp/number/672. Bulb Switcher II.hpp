#include <vector>
#include <string>
#include <list>
#include <unordered_set>

//Button 1: Flips the status of all the bulbs.
//Button 2: Flips the status of all the bulbs with even labels (i.e., 2, 4, ...).
//Button 3: Flips the status of all the bulbs with odd labels (i.e., 1, 3, ...).
//Button 4: Flips the status of all the bulbs with a label j = 3k + 1 where k = 0, 1, 2, ... (i.e., 1, 4, 7, 10, ...).

//Input: n = 3, presses = 1
//Output: 4
//Explanation: Status can be: [off, on, off], [on, off, on], [off, off, off], [off, on, on].
class Solution_672 {
public:
    int flipLights(int n, int presses) {
        if(presses == 0)
            return 1;
        std::vector<bool> init_data(n,true);
        //std::string first(n,'1');
        //existed.insert(first);
        std::list<std::vector<bool>> q;
        q.push_back(init_data);
        std::unordered_set<std::string> prev_existed;
        while(q.size() > 0 && presses > 0){
            int len = q.size();
            std::unordered_set<std::string> existed;
            for(int i = 0;i < len;i++){
                auto pre_data = q.front();
                q.pop_front();
                std::string flip_all_key;
                std::vector<bool> flip_all(n);
                for(int i = 0;i < n;i++){
                    flip_all[i] = !pre_data[i];
                    if(flip_all[i])
                        flip_all_key += "1";
                    else
                        flip_all_key += "0";
                }
                if(existed.find(flip_all_key) == existed.end()){
                    existed.insert(flip_all_key);
                    q.push_back(flip_all);
                }
                std::string flip_even_key;
                std::vector<bool> flip_even(n);
                for(int i = 0;i < n;i++){
                    if((i | 1) != i){
                        flip_even[i] = !pre_data[i];
                    }else{
                        flip_even[i] = pre_data[i];
                    }
                    if(flip_even[i])
                        flip_even_key += "1";
                    else
                        flip_even_key += "0";
                }
                if(existed.find(flip_even_key) == existed.end()){
                    existed.insert(flip_even_key);
                    q.push_back(flip_even);
                }
                std::string flip_odd_key;
                std::vector<bool> flip_odd(n);
                for(int i = 0;i < n;i++){
                    if((i | 1) == i){
                        flip_odd[i] = !pre_data[i];
                    }else{
                        flip_odd[i] = pre_data[i];
                    }
                    if(flip_odd[i])
                        flip_odd_key += "1";
                    else
                        flip_odd_key += "0";
                }
                if(existed.find(flip_odd_key) == existed.end()){
                    existed.insert(flip_odd_key);
                    q.push_back(flip_odd);
                }
                std::string flip_3k_key;
                std::vector<bool> flip_3k(n);
                for(int i = 0;i < n;i++){
                    if((i % 3) == 0){
                        flip_3k[i] = !pre_data[i];
                    }else{
                        flip_3k[i] = pre_data[i];
                    }
                    if(flip_3k[i])
                        flip_3k_key += "1";
                    else
                        flip_3k_key += "0";
                }
                if(existed.find(flip_3k_key) == existed.end()){
                    existed.insert(flip_3k_key);
                    q.push_back(flip_3k);
                }
            }
            if(prev_existed.size() == existed.size()){
                if(prev_existed == existed)
                    break;
            }
            prev_existed = existed;
            presses--;
        }
        return q.size();
    }
};