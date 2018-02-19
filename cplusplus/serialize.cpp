#include <iostream>
#include <string>
#include <sstream>
#include <queue>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

// recursive
class Codec {
public:
    string serialize(TreeNode* root) {
        ostringstream out;
        shelper(root, out);
        return out.str();
    }

    void shelper(TreeNode *root, ostringstream &out) {
        if (root == NULL)
            out << "# ";
        else {
            out << root->val << " ";

            shelper(root->left, out);
            shelper(root->right, out);    
        }
    }

    TreeNode *deserialize(string data) {
        istringstream in(data);
        return dhelper(in);
    }

    TreeNode *dhelper(istringstream &in) {
        string val;
        in >> val; 
        if (val == "#")
            return nullptr;

        TreeNode *ret = new TreeNode(stoi(val));

        ret->left = dhelper(in);
        ret->right = dhelper(in);

        return ret;
    }
};

// iterative
class Codec2 {
public:
    string serialize(TreeNode* root) {
        if (root == NULL)
            return "";

        ostringstream out;
        queue<TreeNode*> q;
        q.push(root);

        while (!q.empty()) {
            TreeNode *cur;
            cur = q.front();
            q.pop();
            if (cur) {
                out << cur->val << " ";
                q.push(cur->left);
                q.push(cur->right);
            }
            else {
                out << "# ";
            }
        }

        return out.str();
    }

    TreeNode* deserialize(string data) {
        if (data.empty()) {
            return nullptr;
        }

        istringstream in(data);
        string cur;
        in >> cur; 

        TreeNode* root;
        root = new TreeNode(stoi(cur));

        queue<TreeNode*> q;
        q.push(root);

        while(!q.empty()) {
            TreeNode* node;
            node = q.front();
            q.pop();

            in >> cur;
            if (cur != "#") {
                TreeNode* leftNode = new TreeNode(stoi(cur));
                node->left = leftNode;
                q.push(leftNode);
            }
            
            in >> cur;
            if (cur != "#") {
                TreeNode* rightNode = new TreeNode(stoi(cur));
                node->right = rightNode;
                q.push(rightNode);
            }
        }
        return root;

    }

};