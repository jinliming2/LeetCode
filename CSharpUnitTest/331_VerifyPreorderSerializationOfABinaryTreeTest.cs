using _331_VerifyPreorderSerializationOfABinaryTree;
using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace CSharpUnitTest {
    [TestClass]
    public class _331_VerifyPreorderSerializationOfABinaryTreeTest {
        Solution solution;
        public _331_VerifyPreorderSerializationOfABinaryTreeTest() {
            solution = new Solution();
        }
        [TestMethod]
        public void IsValidSerializationTest() {
            Assert.AreEqual(true, solution.IsValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"));
            Assert.AreEqual(false, solution.IsValidSerialization("1,#"));
            Assert.AreEqual(false, solution.IsValidSerialization("9,#,#,1"));
            Assert.AreEqual(true, solution.IsValidSerialization("#"));
            Assert.AreEqual(false, solution.IsValidSerialization("9,3,4,#,#,1,#,#,#,2,#,6,#,#"));
        }
    }
}
