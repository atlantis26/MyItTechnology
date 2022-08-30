package InterfaceAndClass;

/**
 * @Author: yangyong
 * @Date: 2020/4/8 14:23
 * @Version 1.0
 */
public class TeacherTest {
    public static void main(String[] args){

        //分别调用不同的子类
//        System.out.println("=========数据库老师开始教学============");
//        DBTeacher dbtest = new DBTeacher();//使用new 关键字，创建对象
//        dbtest.work();//使用对象，对象名.方法名(), 引用对象的方法
//
//        System.out.println("=========Java老师开始教学============");
//        JavaTeacher javaTest = new JavaTeacher();
//        javaTest.work();

        /**
         * 抽象类不能实例化，但无论DBTeacher还是JavaTeacher，它们都是Teacher类的子类，因此可以向上转型为该类，从而实现多态
         * 多态性是面向对象编程的一个重要特征，它是指在父类中定义的属性和方法被子类继承之后，可以具有不同的数据类型或表现出不同的行为
         * 使程序具有更好的扩展性，并可以对类进行通用的处理
         */

        System.out.println("=========数据库老师开始教学============");
        Teacher tea1 = new DBTeacher();
        tea1.work();

        System.out.println("=========Java老师开始教学============");
        Teacher tea2 = new JavaTeacher();
        tea2.work();

    }

}
