package ExceptionAndAssert;

import java.util.Scanner;

/**
 * @Author: yangyong
 * @Date: 2020/4/10 17:24
 * @Version 1.0
 */
public class ExceptHandle1 {
    public static void main(String[] args) {
        //控制台输入字符
        Scanner sc = new Scanner(System.in);
        try {
            //要求输入的必须是整数，如果不是整数，则会抛出异常
            int number = sc.nextInt();
            System.out.println("你输入的值是："+number);
        } catch (Exception ex) {

            //e.getMessage(); 只会获得具体的异常名称. 比如说NullPoint 空指针,就告诉你说是空指针...
            System.out.println(ex.getMessage());

            //e.toString()获取的信息包括异常类型和异常详细消息
            System.out.println(ex.toString());

            //e.printStackTrace();会打出详细异常,异常名称,出错位置,便于调试用..
            ex.printStackTrace();

        }
    }
}
