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
    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
        return helper(inorder, 0, inorder.size(), postorder, 0, postorder.size());
    }

    TreeNode* helper(vector<int>& inorder, int i_left, int i_right, vector<int>& postorder, int p_left, int p_right) {
        if (i_left >= i_right || p_left >= p_right) {
            return NULL;
        }

        int value = postorder[p_right-1];

        TreeNode* root = new TreeNode(value);

        auto f = find(inorder.begin() + i_left, inorder.begin() + i_right, value);
        int dis = f - inorder.begin() - i_left;

        vector<int> l_inorder(inorder.begin(), inorder.begin()+position-1);
        vector<int> r_inorder(inorder.begin()+position+1, inorder.end()-1);
        vector<int> l_postorder(postorder.begin(), postorder.begin()+position);
        vector<int> r_postorder(postorder.begin()+position+1, postorder.end()-2);

        root->left = helper(inorder, i_left, i_left+dis, postorder, p_left, p_left+dis);
        root->right = helper(inorder, i_left+dis+1, i_right, postorder, p_left+dis, p_right-1);
        return root;

    }

    int findIndex(vector<int>& data, int value) {
        for (int i=0; i<data.size(); i++) {
            if (data[i] == value) {
                return i;
            }
        }
        return -1;
    }
};