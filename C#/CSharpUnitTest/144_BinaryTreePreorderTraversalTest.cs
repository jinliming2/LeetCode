using _144_BinaryTreePreorderTraversal;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using System.Collections.Generic;

namespace CSharpUnitTest {
    [TestClass]
    public class _144_BinaryTreePreorderTraversalTest {
        Solution solution;
        public _144_BinaryTreePreorderTraversalTest() {
            solution = new Solution();
        }
        [TestMethod]
        public void PreorderTraversalTest() {
            TreeNode tn = new TreeNode(1);
            tn.left = null;
            tn.right = new TreeNode(2);
            tn.right.left = new TreeNode(3);
            tn.right.right = null;
            tn.right.left.left = null;
            tn.right.left.right = null;
            IList<int> list = solution.PreorderTraversal(tn);
            Assert.AreEqual(3, list.Count);
            Assert.AreEqual(1, list[0]);
            Assert.AreEqual(2, list[1]);
            Assert.AreEqual(3, list[2]);
        }
    }
}
