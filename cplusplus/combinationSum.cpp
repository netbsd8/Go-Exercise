#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    vector<vector<int>> combinationSum2(vector<int>& candidates, int target) {
        vector<vector<int>> ans;
        vector<int> out;

        sort(candidates.begin(), candidates.end());

        combinSumHelper(ans, out, candidates, target, 0);

        return ans;        
    }

    void combinSumHelper(vector<vector<int>> &ans, vector<int> &out, vector<int> &candidates, int target, int start) {
        if (target == 0) {
            ans.push_back(out);
            return;
        }
        if (target < 0) {
            return;
        }

        for (int i=start; i<candidates.size(); i++) {
            if (i > start && candidates[i] == candidates[i-1]) {
                continue;
            }

            out.push_back(candidates[i]);
            combinSumHelper(ans, out, candidates, target-candidates[i], i+1);
            out.pop_back();
        }
    }
};