class Solution {
public:
    bool backspaceCompare(string S, string T) {
        std::string s2,t2;
        for(auto it = S.begin(); it != S.end();it++){
            if(*it == '#'){
                if(!s2.empty()){
                    s2.erase(s2.end()-1);
                }
            }
            else{
                s2.push_back(*it);
            }
        }
        for(auto it = T.begin(); it != T.end();it++){
            if(*it == '#'){
                if(!t2.empty()){
                    t2.erase(t2.end()-1);
                }
            }
            else{
                t2.push_back(*it);
            }
        }
        return s2 == t2;
    }
};