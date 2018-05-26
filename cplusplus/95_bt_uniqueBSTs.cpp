#include <iostream>
#include <vector>
#include <unordered_set>
#include <stack>
#include <string>
#include <algorithm>
#include <bits/stdc++.h>

using namespace std;

struct TreeNode {
    TreeNode    *left;
    TreeNode    *right;
    int         val;
    TreeNode(int x): val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    int numTrees(int n) {
        vector<int> dp(n+1, 0);
        dp[0] = 1;
        dp[1] = 1;

        for (int i=2; i<=n; i++) {
            for (int j=0; j<i; j++) {
                dp[i] = dp[i] + dp[j]*dp[i-j-1];
            }
        }

        return dp[n];
    }

    vector<TreeNode*> generateTrees(int n) {
        if (n == 0) {
            vector<TreeNode*> res;
            return res;
        }

        return geneHelper(1, n);
    }

    vector<TreeNode*> geneHelper(int l, int r) {
        vector<TreeNode*> res;
        if (l > r) {
            res.push_back(NULL);
            return res;
        }

        for (int i=l; i<=r; i++) {
            vector<TreeNode*> leftTrees = geneHelper(l, i-1);
            vector<TreeNode*> rightTrees = geneHelper(i+1, r);
            for (int m=0; m<leftTrees.size(); m++) {
                for (int n=0; n<rightTrees.size(); n++) {
                    TreeNode* root = new TreeNode(i);
                    root->left = leftTrees[m];
                    root->right = rightTrees[n];
                    res.push_back(root);
                }
            } 
        }
        return res;
    }
};