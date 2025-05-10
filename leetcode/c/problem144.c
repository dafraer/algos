/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int preorder(int *arr, struct TreeNode *root, int i) {
  if (root == NULL) {
    return i - 1;
  }
  arr[i] = root->val;
  int x = preorder(arr, root->left, i + 1);
  int ans = preorder(arr, root->right, x + 1);
  return ans;
}

int count(struct TreeNode *root) {
  if (root == NULL) {
    return 0;
  }
  return count(root->left) + count(root->right) + 1;
}

int *preorderTraversal(struct TreeNode *root, int *returnSize) {
  int len = count(root);
  *returnSize = len;
  int *arr = (int *)malloc(sizeof(int) * len);
  preorder(arr, root, 0);
  return arr;
}
