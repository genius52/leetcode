#include <string>
#include <unordered_map>
using namespace std;

class Solution_535 {
public:
    std::unordered_map<std::string,std::string> short2long;
    std::unordered_map<std::string,std::string> long2short;
    std::string random_key = "abcdefghijklmnopqrstuvwxyz";
    // Encodes a URL to a shortened URL.
    string encode(string longUrl) {
        std::string key;
        for(int i = 0;i < 6;i++){
            int pos = rand() % 26;
            key += random_key[pos];
        }
        std::string encoded_url = "http://tinyurl.com/" + key;
        short2long[encoded_url] = longUrl;
        long2short[longUrl] = encoded_url;
        return encoded_url;
    }

    // Decodes a shortened URL to its original URL.
    string decode(string shortUrl) {
        return short2long[shortUrl];
    }
};

