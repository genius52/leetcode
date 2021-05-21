#include <string>
using namespace std;

//Input: formula = "K4(ON(SO3)2)2"
//Output: "K4N2O14S4"
//Explanation: The count of elements are {'K': 4, 'N': 2, 'O': 14, 'S': 4}.
class Solution_726 {
public:
    std::map<std::string,int> recursive_countOfAtoms(string formula,int left,int right){
        std::map<std::string,int> cur_map;
        if (left >= right)
            return cur_map;
        int visit = left;
        int indgree = 0;
        while(visit <= right && formula[visit] != ')'){
            if(formula[visit] == '('){
                visit++;
                int sub_left = visit;
                indgree++;
                while(visit < right){
                    if(formula[visit] == '(')
                        indgree++;
                    else if(formula[visit] == ')')
                        indgree--;
                    if(indgree == 0)
                        break;
                    visit++;
                }
                auto sub_map = recursive_countOfAtoms(formula,sub_left,visit);
                int cnt = 1;
                visit++;
                int next = visit;
                while (next <= right && (formula[next] >= '0' && formula[next] <= '9')){
                    next++;
                }
                if(next != visit){
                    std::string s_num = formula.substr(visit,next - visit);
                    stringstream convert(s_num);
                    convert >> cnt;
                }
                for(auto it :sub_map){
                    if(it.first.size() > 1){
                        for(auto it2 : it.first){
                            std::string single;
                            single += it2;
                            cur_map[single] += cnt;
                        }
                    }else{
                        cur_map[it.first] += it.second * cnt;
                    }

                }
                visit = right;
            }else if((formula[visit] >= 'a' && formula[visit] <= 'z') ||
                    (formula[visit] >= 'A' && formula[visit] <= 'Z')){
                int cnt = 1;
                int next = visit;
                while(next < right && (formula[next] >= 'a' && formula[next] <= 'z') ||
                      (formula[next] >= 'A' && formula[next] <= 'Z')){
                    next++;
                }
                std::string pre_string = formula.substr(visit,next - visit);
                visit = next;
                while(next < right && formula[next] >= '0' && formula[next] <= '9'){
                    next++;
                }
                if(next != visit){
                    std::string s_num = formula.substr(visit,next - visit);
                    stringstream convert(s_num);
                    convert >> cnt;
                }
                cur_map[pre_string] += cnt;
                visit = next;
            }else{
                visit++;
            }
        }
        return cur_map;
    }

    string countOfAtoms(string formula) {
        int len = formula.length();
        auto char_map = recursive_countOfAtoms(formula,0,len - 1);
        std::string res;
        for(auto it : char_map){
            res += it.first;
            if(it.second != 1)
                res += std::to_string(it.second);
        }
        return res;
    }
};