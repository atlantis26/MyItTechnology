package StaticMethod;

/**
 * @Author: yangyong
 * @Date: 2020/4/9 14:13
 * @Version 1.0
 * 成员变量：在类中定义，方法外的变量，描述类的属性
 * 局部变量：在类的方法中定义，在方法中临时保存数据
 */
public class StaticAndShiLiBianLiang {
    /**
     * 静态变量和实例变量可以统称为成员变量
     * 静态变量也叫做类变量，独立于方法之外的变量，有static修饰。实例变量同样独立也是独立于方法之外的变量，但没有static修饰
     */
    private static int staticInt = 2;//静态变量
    private int random = 2;//实例变量

    public StaticAndShiLiBianLiang() {
        staticInt++;
        random++;
        System.out.println("staticInt = "+staticInt+"  random = "+random);
    }

    /**
     * 区别总结如下：
     *
     * 实例变量属于某个对象的属性，必须创建了实例对象，其中的实例变量才会被分配空间，才能使用这个实例变量。结合上述给出的例子。
     * 每创建一个实例对象，就会分配一个random，实例对象之间的random是互不影响的，所以就可以解释为什么输出的两个random值是相同的了。
     *
     * 静态变量不属于某个实例对象，而是属于整个类。只要程序加载了类的字节码，不用创建任何实例对象，静态变量就回被分配空间，静态变量就可以被使用了。
     * 结合上述给出的例子，无论创建多少个实例对象，永远都只分配一个staticInt 变量，并且每创建一个实例对象,staticInt就会加一。
     */

    public static void main(String[] args) {
        StaticAndShiLiBianLiang test = new StaticAndShiLiBianLiang();
        StaticAndShiLiBianLiang test2 = new StaticAndShiLiBianLiang();
    }
}