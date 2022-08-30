package StaticMethod;

/**
 * @Author: yangyong
 * @Date: 2020/4/9 13:40
 * @Version 1.0
 */
public class StaticBianLiang {

    /*
    static修饰的变量为静态变量，所有类的对象共享该变量
    静态成员属于整个类，当系统第一次使用该类时，就会为其分配内存空间直到该类被卸载才会进行资源回收
     */
    private static String bianliang = "yyyy";

    public static void main(String[] args){

        //静态变量可以直接使用类名来访问，无需创建类的对象
        System.out.println("通过类名访问bianliang: " +StaticBianLiang.bianliang);

        //创建类的对象
        StaticBianLiang bl = new StaticBianLiang();

        //使用对象名来访问静态变量
        System.out.println("通过对象名访问bianliang: " +bl.bianliang);

        //使用对象名的形式修改静态变量的值
        bl.bianliang = "aaaa";

        //再次使用类名访问静态变量，值已被修改
        System.out.println("通过类名访问修改后的bianliang: " +StaticBianLiang.bianliang);

    }
}
