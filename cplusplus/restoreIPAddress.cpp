#include <iostream>
#include <vector>
#include <unordered_set>
#include <stack>
#include <string>
#include <algorithm>
#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    vector<string> restoreIpAddresses(string s) {
        vector<string> res;
        string out;
        if (s.size()>12) {
            return res;
        }
        helper(res, out, s, 0);
        return res;        
    }

    void helper(vector<string> &res, string &out, string s, int count) {
        if (s.empty() && count == 4) {
            res.push_back(out);
            return;
        }
        for (int j=1; j<=3; j++) {
            string ss = s.substr(0, j);
            if (s.size()>=j && isValid(ss)) {
                string t;
                if (out.size() == 0) {
                    t = ss;
                } else {
                    t = out + "." + ss;
                }
                helper(res, t, s.substr(j), count+1);
            }
        }
    }

    bool isValid(string s) {
        if (s.empty() || s.size() > 3 || (s.size() > 1 && s[0] == '0')) return false;
        int res = atoi(s.c_str());
        return res <= 255 && res >= 0;
    }
    /*
    bool isValid(string s) {
        if (stoi(s, nullptr, 10) >=0 && stoi(s, nullptr, 10) <= 255)
            return true;
        return false;
    }
    */
};