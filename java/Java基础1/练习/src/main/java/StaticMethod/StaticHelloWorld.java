package StaticMethod;

/**
 * @Author: yangyong
 * @Date: 2020/4/9 15:20
 * @Version 1.0
 */
public class StaticHelloWorld {
    //使用static关键字声明静态方法
    public static void print(){
        System.out.println("欢迎您：yyyy");
    }

    public static void main(String[] args){
        //直接使用类名调用方法
        StaticHelloWorld.print();

        //也可以通过对象名调用，当然更推荐使用类名调用的方式
        StaticHelloWorld demo = new StaticHelloWorld();
        demo.print();
    }
}
