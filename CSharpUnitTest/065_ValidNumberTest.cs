using _065_ValidNumber;
using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace CSharpUnitTest {
    [TestClass]
    public class _065_ValidNumberTest {
        Solution solution;
        public _065_ValidNumberTest() {
            solution = new Solution();
        }
        [TestMethod]
        public void IsNumberTest() {
            Assert.AreEqual(true, solution.IsNumber("0"));
            Assert.AreEqual(true, solution.IsNumber(" 0.1 "));
            Assert.AreEqual(false, solution.IsNumber("abc"));
            Assert.AreEqual(false, solution.IsNumber("1 a"));
            Assert.AreEqual(true, solution.IsNumber("2e10"));
            Assert.AreEqual(false, solution.IsNumber(""));
            Assert.AreEqual(false, solution.IsNumber(" "));
            Assert.AreEqual(true, solution.IsNumber(".1"));
            Assert.AreEqual(false, solution.IsNumber("e"));
            Assert.AreEqual(true, solution.IsNumber("3."));
            Assert.AreEqual(false, solution.IsNumber("e100"));
            Assert.AreEqual(false, solution.IsNumber("."));
            Assert.AreEqual(true, solution.IsNumber("2e0"));
            Assert.AreEqual(true, solution.IsNumber("-1."));
            Assert.AreEqual(true, solution.IsNumber(" +.8"));
            Assert.AreEqual(false, solution.IsNumber(" +0e-"));
            Assert.AreEqual(true, solution.IsNumber("46.e3"));
            Assert.AreEqual(false, solution.IsNumber(".e1"));
            Assert.AreEqual(false, solution.IsNumber("6e6.5"));
        }
    }
}
