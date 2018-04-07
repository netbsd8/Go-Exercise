#include <iostream>
#include <vector>
#include <unordered_set>
#include <stack>
#include <algorithm>
#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    vector<int> grayCode(int n) {
        vector<int> res{0};

        for (int i=0; i<n; i++) {
            int size = res.size();
            for (int j=size-1; j>=0; j--) {
                res.push_back(res[j] | 1 << i);
            }
        }

        return res;
    }
    vector<int> grayCodeStack(int n) {
        vector<int> res;
        unordered_set<int> s;
        stack<int> st;
        st.push(0);
        s.insert(0);

        while (!st.empty()) {
            int t = st.top();
            st.pop();
            for (int i=0; i<n; i++) {
                int k = t;
                if (k & (1<<i) == 0)
                    k |= (1<<i);
                else
                    k &= ~(1<<i);
                if (s.count(k))
                    continue;
                s.insert(k);
                st.push(k);
                res.push_back(k);
                break;
            }
        }

        return res;
    }

    vector<int> grayCode2(int n) {
        unordered_set<int> s;
        vector<int> res;

        helper(n, s, 0, res);

        return res;

    }

    void helper(int n, unordered_set<int>& s, int out, vector<int>& res) {
        if (s.count(out) == 0) {
            s.insert(out);
            res.push_back(out);
        }

        for (int i=0; i<n; i++) {
            int t = out;
            if (t&(1<<i) == 0) {
                t = t | 1<<i;
            } else {
                t = t & ~(1<<i);
            }
            if (s.count(t)) {
                continue;
            }

            helper(n, s, t, res);
            break;
        }
    }
};