package InterfaceAndClass;

/**
 * @Author: yangyong
 * @Date: 2020/4/8 14:21
 * @Version 1.0
 */
public class JavaTeacher extends Teacher {
    //实现父类的抽象方法
    @Override
    public void teaching(){
        System.out.println("打开Idea软件");
        System.out.println("书写Java代码");
    }
    //访问父类中的非抽象方法
    @Override
    public void work() {
        super.work();
    }
}
