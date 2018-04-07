#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    vector<vector<int>> combinationSum3(int k, int n) {
        vector<vector<int>> ans;
        vector<int> out;

        combinationSum3Helper(ans, out, k, n, 1);

        return ans;
    }

    void combinationSum3Helper(vector<vector<int>> &ans, vector<int> &out, int k, int n, int start) {
        if (out.size() == k && n == 0) {
            ans.push_back(out);
            return;
        }

        for (int i=start; i<=9; i++) {
            out.push_back(i);
            combinationSum3Helper(ans, out, k, n-i, i+1);
            out.pop_back();
        }
    }
};