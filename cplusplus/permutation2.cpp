#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        vector<vector<int>> ans;
        vector<int> out;
        vector<bool> visited(nums.size();

        sort(nums.begin(), nums.end()); 

        permuteUniqueHelper(ans, out, nums, visited);

        return ans;
    }

    void permuteUniqueHelper(vector<vector<int>> &ans, vector<int> &out, vector<int> &nums, vector<bool> &visited) {
        if (out.size() == nums.size()) {
            ans.push_back(out);
            return;
        }

        for (int i=0; i<nums.size(); i++) {
            // for the same loop
            if (visited[i] == true) {
                continue;
            }
            // for a different loop
            if (i>0 && nums[i] == nums[i-1] && visited[i-1] == 0) {
                continue;
            }
            visited[i] = true;
            out.push_back(nums[i]);
            permuteUniqueHelper(ans, out, nums, visited);
            out.pop_back();
            visited[i] = false;
        }
    }
};