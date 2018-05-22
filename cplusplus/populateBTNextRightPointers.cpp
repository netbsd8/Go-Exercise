#include <iostream>
#include <vector>
#include <unordered_set>
#include <stack>
#include <string>
#include <algorithm>
#include <bits/stdc++.h>

using namespace std;

struct TreeLinkNode {
    int val;
    TreeLinkNode *left, *right, *next;
    TreeLinkNode(int x) : val(x), left(NULL), right(NULL), next(NULL) {}
};

class Solution {
public:
    void connect(TreeLinkNode *root) {
        if (root == NULL) {
            return;
        }

        queue<TreeLinkNode*> q;
        q.push(root);

        while (!q.empty()) {
            TreeLinkNode *prev;
            int count = q.size();
            for (int i=0; i<count; i++) {
                TreeLinkNode *t = q.front();
                q.pop();
                if (i == 0) {
                    prev = t;
                    if (t->left) q.push(t->left);
                    if (t->right) q.push(t->right);
                    continue;
                }
                prev->next = t;
                if (t->left) q.push(t->left);
                if (t->right) q.push(t->right);
                prev = t;
                if (i == count-1) {
                    t->next = NULL;
                }
            }
        }
    }

    // recursion
    void connect2(TreeLinkNode *root) {
        if (root == NULL) {
            return;
        }

        if (root->left) root->left->next = root->right;
        if (root->right) {
            if (root->next) {
                root->right->next = root->next->left;
            } else {
                root->right->next = NULL;
            }
        }
        connect2(root->left);
        connect2(root->right);
    }
};