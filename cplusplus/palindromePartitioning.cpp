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
    vector<vector<string> > partition(string s) {
        vector<vector<string> > res;
        vector<string> out;

        helper(res, out, s);
        return res;
    }

    void helper(vector<vector<string> >& res, vector<string> &out, string s) {
        if (s.empty()) {
            res.push_back(out);
            return;
        }

        for (int i=1; i<=s.size(); i++) {
            string t = s.substr(0, i);
            if (isPalindrome(t)) {
                out.push_back(t);
                helper(res, out, s.substr(i));
                out.pop_back();
            }
        }
    }

    bool isPalindrome(string s) {
        int len = s.size();
        if (len == 1) {
            return true;
        }

        int l = 0;
        int r = len-1;
        while (l < r) {
            if (s[l] != s[r]) {
                return false;
            }
            l++;
            r--;
        }
        return true;
    }
};