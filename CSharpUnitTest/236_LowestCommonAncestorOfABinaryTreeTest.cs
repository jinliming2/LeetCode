using _236_LowestCommonAncestorOfABinaryTree;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using System.Collections.Generic;

namespace CSharpUnitTest {
    [TestClass]
    public class _236_LowestCommonAncestorOfABinaryTreeTest {
        Solution solution;
        public _236_LowestCommonAncestorOfABinaryTreeTest() {
            solution = new Solution();
        }
        [TestMethod]
        public void LowestCommonAncestorTest() {
            TreeNode tn = new TreeNode(3);
            tn.left = new TreeNode(5);
            tn.left.left = new TreeNode(6);
            tn.left.right = new TreeNode(2);
            tn.left.right.left = new TreeNode(7);
            tn.left.right.right = new TreeNode(4);
            tn.right = new TreeNode(1);
            tn.right.left = new TreeNode(0);
            tn.right.right = new TreeNode(8);
            Assert.AreEqual(3, solution.LowestCommonAncestor(tn, tn.left, tn.right).val);
            Assert.AreEqual(5, solution.LowestCommonAncestor(tn.left, tn.left, tn.left.right.right).val);
            tn = new TreeNode(1);
            tn.left = new TreeNode(2);
            tn.left.left = new TreeNode(3);
            tn.right = new TreeNode(4);
            Assert.AreEqual(1, solution.LowestCommonAncestor(tn, tn.left.left, tn.right).val);
            tn = new TreeNode(3);
            tn.left = new TreeNode(5);
            tn.right = new TreeNode(1);
            tn.left.left = new TreeNode(6);
            tn.left.right = new TreeNode(2);
            tn.right.left = new TreeNode(0);
            tn.right.right = new TreeNode(8);
            tn.left.right.left = new TreeNode(7);
            tn.left.right.right = new TreeNode(4);
            Assert.AreEqual(5, solution.LowestCommonAncestor(tn, tn.left, tn.left.right.right).val);
        }
    }
}
