#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    vector<vector<int>> subsets(vector<int>& nums) {
        vector<vector<int>> ans;
        vector<int> out;

        sort(nums.begin(), nums.end());
        ans.push_back(vector<int>{});

        subsetHelper(ans, out, nums, 0);

        return ans; 
    }

    void subsetHelper(vector<vector<int>> &ans, vector<int> &out, vector<int> &nums, int start) {
        for (int i=start; i<nums.size(); i++) {
            if (i>start && nums[i] == nums[i-1]) {
                continue;
            }
            out.push_back(nums[i]);
            ans.push_back(out);
            subsetHelper(ans, out, nums, i+1);
            out.pop_back();
        }
    }
};