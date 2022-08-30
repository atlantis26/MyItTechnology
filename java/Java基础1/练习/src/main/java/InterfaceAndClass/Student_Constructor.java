package InterfaceAndClass;

/**
 * @Author: yangyong
 * @Date: 2020/4/9 10:20
 * @Version 1.0
 */
public class Student_Constructor {
    private String name;
    private String gender;
    private long ID;
    private long birthday;


    /**
     * 构造方法的主要作用就是创建对象,构造方法分为有参数和没有参数两种。有参数的构造方法主要就是用于对创建出来的对象进行初始化
     * 当你需要对创建出来的类的对象进行初始化赋值的时候，构造参数就比较有用
     * 构造方法与类同名且没有返回值
     */
    Student_Constructor(String name, String gender, long ID, long birthday) {
        this.name = name;
        this.gender = gender;
        this.ID = ID;
        this.birthday = birthday;
    }

    /**
     * 另一种是没有参数的构造方法，又叫缺省构造方法
     * 其实，你的类如果没有自己手动去定义任何构造方法，Java的编译器会配上一个自动缺省构造方法，这个构造方法是空的，不做任何事情，只是为了满足编译需要
     */
    Student_Constructor(){

    }

    public static void main(String[] args) {

        /**
         * 这样在new一个对象出来的时候，就对它进行了初始化，避免了某些属性忘记初始化的问题；同时可以看出代码长度远小于上一个例子，提高了程序的可阅读性
         * 只要你new一个对象出来，就会相应地调用这个类的构造方法。
         * 有参数的构造方法可以对对象进行初始化，但建议写了有参的构造方法后再写一个无参的空的构造方法，便于创建一个对象而不给它的成员变量赋初值
         * 要注意，如果自己写了有参的构造方法，编译器是不会再补充一个缺省构造方法的
         */
        Student_Constructor s = new Student_Constructor("Lily", "女", 100001, 200000226);
        System.out.println("姓名：" + s.name +  " 性别：" + s.gender + " 学号：" + s.ID + " 生日：" + s.birthday);

        Student_Constructor ss = new Student_Constructor();

    }
}
