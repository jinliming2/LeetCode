using System.Collections.Generic;

namespace _236_LowestCommonAncestorOfABinaryTree {
    public class TreeNode {
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int x) { val = x; }
    }
    public class Solution {
        public TreeNode LowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
            while(root != null) {
                if(root == p || root == q) {
                    return root;
                }
                bool ll = Exist(root.left, p);
                bool lr = Exist(root.left, q);
                if(ll != lr) {
                    return root;
                }
                root = ll ? root.left : root.right;
            }
            return null;
        }

        private bool Exist(TreeNode root, TreeNode target) {
            if(root == null) {
                return false;
            }
            if(root == target) {
                return true;
            }
            return Exist(root.left, target) || Exist(root.right, target);
        }
    }
}
