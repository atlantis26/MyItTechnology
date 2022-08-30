package StaticMethod;

/**
 * @Author: yangyong
 * @Date: 2020/4/9 15:25
 * @Version 1.0
 */
public class StaticMethodFinal {

    //非静态变量
    String name = "多点";
    //静态变量
    static String yoo = "交易组";//

    //在静态方法调用非静态变量
    public static void print(){
        System.out.println("======静态方法开始=========");
//        System.out.println("欢迎您："+name); //静态方法中不能直接调用非静态变量
        System.out.println("欢迎您："+yoo);  //静态方法中可以直接调用同类中的静态变量

        //如果希望在静态方法中调用非静态变量，可以通过创建类的对象，然后通过对象来访问非静态变量
        StaticMethodFinal demo = new StaticMethodFinal();
        System.out.println("欢迎您："+demo.name);
    }

    //实例方法中，则可以直接访问同类的非静态变量和静态变量
    public void show(){
        System.out.println("======实例方法开始=========");
        System.out.println("欢迎您："+name); //实例方法中可以直接调用非静态变量
        System.out.println("欢迎您："+yoo);  //实例方法中可以直接调用同类中的静态变量
    }

    //静态方法中不能直接调用非静态方法，需要通过对象来访问非静态方法
    public static void main(String[] args){
        print();//静态方法中可以直接调用静态方法
//        show();//静态方法中不能直接调用非静态方法，必须通过对象来调用

        //通过对象来调用非静态方法
        StaticMethodFinal demo = new StaticMethodFinal();
        demo.show();

    }
}
