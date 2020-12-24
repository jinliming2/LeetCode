using System.Collections.Generic;

namespace _144_BinaryTreePreorderTraversal {
    public class TreeNode {
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int x) { val = x; }
    }
    public class Solution {
        public IList<int> PreorderTraversal(TreeNode root) {
            IList<int> list = new List<int>();
            Stack<TreeNode> stack = new Stack<TreeNode>();
            TreeNode current = root;
            while(current != null || stack.Count > 0) {
                while(current != null) {
                    stack.Push(current);
                    list.Add(current.val);
                    current = current.left;
                }
                current = stack.Pop();
                current = current.right;
            }
            return list;
        }
    }
}
