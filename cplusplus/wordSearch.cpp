#include <iostream>
#include <vector>
#include <unordered_set>
#include <stack>
#include <algorithm>
#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    bool exist(vector<vector<char> >& board, string word) {
        if (word.empty()) return true;
        if (board.empty() || board[0].empty()) return false;
        vector<vector<bool> > visited(board.size(), vector<bool>(board[0].size(), false));

        for (int i=0; i<board.size(); i++) {
            for (int j=0; j<board[0].size(); j++) {
                if (search(0, word, board, visited, i, j) == true) return true;
            }
        }
        return false;
    }

    bool search(int index, string word, vector<vector<char> >& board, vector<vector<bool> >& visited, int i, int j) {
        if (index == word.size()) return true;
        if (i<0 || j<0 || i>=board.size() || j>=board[0].size() || visited[i][j] || word[index]!=board[i][j])
            return false;
        visited[i][j] = true;
        bool res = search(index+1, word, board, visited, i+1, j) ||
                search(index+1, word, board, visited, i-1, j) ||
                search(index+1, word, board, visited, i, j+1) ||
                search(index+1, word, board, visited, i, j-1);
                

        visited[i][j] = false;
        return res;
    }
};