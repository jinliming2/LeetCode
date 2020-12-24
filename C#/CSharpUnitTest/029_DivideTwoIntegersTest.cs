using _029_DivideTwoIntegers;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using System;

namespace CSharpUnitTest {
    [TestClass]
    public class _029DivideTwoIntegersTest {
        Solution solution;
        public _029DivideTwoIntegersTest() {
            solution = new Solution();
        }
        [TestMethod]
        public void DivideTest() {
            Random random = new Random();
            for(int i = 0; i < 10; i++) {
                int dividend = random.Next();
                int divisor = random.Next();
                Assert.AreEqual(dividend / divisor, solution.Divide(dividend, divisor));
            }
            Assert.AreEqual(int.MaxValue, solution.Divide(0, 0));
            Assert.AreEqual(int.MaxValue, solution.Divide(10, 0));
            Assert.AreEqual(1 / 2, solution.Divide(1, 2));
            Assert.AreEqual(4 / 2, solution.Divide(4, 2));
            Assert.AreEqual(5 / 2, solution.Divide(5, 2));
            Assert.AreEqual(-1 / 2, solution.Divide(-1, 2));
            Assert.AreEqual(-4 / 2, solution.Divide(-4, 2));
            Assert.AreEqual(-5 / 2, solution.Divide(-5, 2));
            Assert.AreEqual(1 / -2, solution.Divide(1, -2));
            Assert.AreEqual(4 / -2, solution.Divide(4, -2));
            Assert.AreEqual(5 / -2, solution.Divide(5, -2));
            Assert.AreEqual(-1 / -2, solution.Divide(-1, -2));
            Assert.AreEqual(-4 / -2, solution.Divide(-4, -2));
            Assert.AreEqual(-5 / -2, solution.Divide(-5, -2));
            Assert.AreEqual(0 / 2, solution.Divide(0, 2));
            Assert.AreEqual(2 / 2, solution.Divide(2, 2));
            Assert.AreEqual(2147483647, solution.Divide(-2147483648, -1));
            Assert.AreEqual(-2147483648 / -2, solution.Divide(-2147483648, -2));
            Assert.AreEqual(1073741823, solution.Divide(2147483647, 2));
        }
    }
}
