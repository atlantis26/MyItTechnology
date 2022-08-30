package ExceptionAndAssert;

/**
 * @Author: yangyong
 * @Date: 2020/4/10 18:15
 * @Version 1.0
 */

public class AssertVerify {

    public static void main(String[] args)
    {
        System.out.println("----start---");

        boolean isOpen = false;

        assert isOpen=true;//如果开启了断言，会将isOpen的值改为true

        System.out.println("是否开启了断言"+isOpen);  //打印是否开启了断言

        if (isOpen)
        {
            int value=-1;
            //assert为假的时候会中断程序的执行，并且输出了详细的错误信息，比try catch更简洁，发布程序的时候只需要去掉-ea就行，在调试程序的时候非常有用
            assert 0 < value:"value="+value;
        }
        System.out.println("----end----");
    }
}