namespace _029_DivideTwoIntegers {
    public class Solution {
        public int Divide(int dividend, int divisor) {
            if(divisor == 0) {
                return int.MaxValue;
            }
            if(dividend == 0) {
                return 0;
            }
            if(divisor == 1) {
                return dividend;
            }
            if(divisor == -1) {
                return dividend == int.MinValue ? int.MaxValue : -dividend;
            }
            bool flag = (dividend > 0) ^ (divisor > 0);
            long a = dividend < 0 ? -dividend : dividend;
            long b = divisor < 0 ? -divisor : divisor;
            if(a < 0) {
                a = (long)int.MaxValue + 1;
            }
            if(b < 0) {
                b = (long)int.MaxValue + 1;
            }
            int sum = 0, last;
            while(a >= b) {
                long temp = b;
                int i = 0;
                while(a > temp) {
                    temp <<= 1;
                    i++;
                }
                if(a == temp) {
                    last = sum;
                    sum |= (1 << i);
                    if(sum < last) {
                        return int.MaxValue;
                    }
                    break;
                } else {
                    last = sum;
                    sum |= (1 << (i - 1));
                    if(sum < last) {
                        return int.MaxValue;
                    }
                    a -= (temp >> 1);
                }
            }
            return flag ? -sum : sum;
        }
    }
}
