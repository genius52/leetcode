#include <string>
using namespace std;

class Solution_1410 {
public:
    string entityParser(string text) {
        std::string res;
        int len = text.length();
        int i = 0;
        while (i < len) {
            if (text[i] != '&') {
                res += text[i];
                i++;
            }
            else {
                int tmp = i;
                int j = tmp + 4;
                bool find = false;
                if (j <= len) {
                    string sub = text.substr(tmp, 4);
                    if (sub == "&gt;" ) {
                        res += '>';
                        find = true;
                    }
                    else if (sub == "&lt;" ) {
                        res += '<';
                        find = true;
                    }
                    if(find)
                        i = j;
                }
                if (!find) {
                    j = tmp + 5;
                    if (j <= len) {
                        string sub = text.substr(tmp, 5);
                        if (sub == "&amp;") {
                            res += '&';
                            find = true;
                        }
                        if (find)
                            i = j;
                    }
                }
                if (!find) {
                    j = tmp + 6;
                    if (j <= len) {
                        string sub = text.substr(tmp, 6);
                        if (sub == "&quot;") {
                            res += '\"';
                            find = true;
                        }
                        else if (sub == "&apos;") {
                            res += '\'';
                            find = true;
                        }
                        if (find)
                            i = j;
                    }
                }
                if (!find) {
                    j = tmp + 7;
                    if (j <= len) {
                        string sub = text.substr(tmp, 7);
                        if (sub == "&frasl;") {
                            res += '/';
                            find = true;
                        }
                        if (find)
                            i = j;
                    }
                }
                if (!find) {
                    res += '&';
                    i++;
                }
            }
        }
        return res;
    }
};