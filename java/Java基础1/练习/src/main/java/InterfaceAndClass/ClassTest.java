package InterfaceAndClass;

/**
 * @Author: yangyong
 * @Date: 2020/4/7 19:34
 * @Version 1.0
 */
public class ClassTest {

    //声明类的属性，也就是成员变量
    int age = 23;

    //声明类的行为，也就是方法
    public void getAge(){

        System.out.print("the age is "+age);
    }

    /**
     * public static void main(String[] args)，是java程序的入口地址，java虚拟机运行程序的时候首先找的就是main方法
     * 参数String[] args是一个字符串数组，接收来自程序执行时传进来的参数
     * public 表示程序的访问权限，表示的是任何的场合可以被引用
     * static表示方法是静态的，不依赖类的对象的
     * void 表示方法是不需要返回值的
     */
    public static void main(String[] args) {

        //用这个类来实例化一个对象
        ClassTest test= new ClassTest();

        //访问类中的方法
        test.getAge();
    }
}
