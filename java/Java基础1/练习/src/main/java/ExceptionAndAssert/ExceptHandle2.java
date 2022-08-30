package ExceptionAndAssert;

import java.util.Scanner;

/**
 * @Author: yangyong
 * @Date: 2020/4/10 16:17
 * @Version 1.0
 */


public class ExceptHandle2 {

    public static void main(String[] args){
        //控制台输入字符
        Scanner sc = new Scanner(System.in);
        //要求输入的必须是整数
        int age = sc.nextInt();

        //由于该方法已经声明了异常，那么你在使用这个方法的时候就必须处理这个异常，不处理直接调用，代码会报错
       //enter(age);

        //提供异常处理代码，使用try-catch或者try-catch-finally结构
        try {
            enter(age);
        }
        catch(Exception e) {
            //打印出enter方法中手动抛出的异常
            System.out.println(e.toString());
        }
        finally {
            //不论是否有异常，最后都会执行finally里面的代码
            System.out.println("=========结束===========");
        }

    }

    //throws 声明了此方法可能有异常
    public static void enter(int age) throws Exception {

        if (age < 18) {
            System.out.println("年龄不合适~");

            //throw 手动向外抛出异常
            throw new Exception("to small");

        } else {
            System.out.println("欢迎你，成年人~");
        }
    }
}
