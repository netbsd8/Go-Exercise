/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if (root == nullptr) {
            return nullptr;
        }

        if (root->val > max(p->val, q->val)) {
            return lowestCommonAncestor(root->left, p, q);
        }
        if (root->val < min(p->val, q->val)) {
            return lowestCommonAncestor(root->right, p, q);
        }
        return root;
    }

    TreeNode* lowestCommonAncestor2(TreeNode* root, TreeNode* p, TreeNode* q) {
        while (true) {
            if (root->val > max(p->val, q->val)) {
                root = root->left; 
            }
            else if (root->val < min(p->val, q->val)) {
                root = root->right;
            }
            else {
                break;
            }

        }
        return root;
    }    
    
};