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
int postorder(int *arr, struct TreeNode *root, int i) {
  if (root == NULL) {
    return i - 1;
  }
  int x = postorder(arr, root->left, i);
  int ans = postorder(arr, root->right, x + 1);
  arr[ans + 1] = root->val;
  return ans + 1;
}

int count(struct TreeNode *root) {
  if (root == NULL) {
    return 0;
  }
  return count(root->left) + count(root->right) + 1;
}

int *postorderTraversal(struct TreeNode *root, int *returnSize) {
  int len = count(root);
  *returnSize = len;
  int *arr = (int *)malloc(sizeof(int) * len);
  postorder(arr, root, 0);
  return arr;
}
