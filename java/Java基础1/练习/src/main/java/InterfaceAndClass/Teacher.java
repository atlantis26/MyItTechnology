package InterfaceAndClass;

/**
 * @Author: yangyong
 * @Date: 2020/4/8 14:02
 * @Version 1.0
 */

/**
    设计一个简单的老师授课系统

　　基本需求：老师授课，每个老师教不同的课程。

　　抽象出1个对象：老师（Teacher），有名字、年龄、职业等特征，有上课前准备设备、正式上课、上课后关闭设备等行为；
 */
public abstract class Teacher {
    /**
     * 老师工作的方法是一个流程，也可以认为是一个公共模板方法（即模板方法中的每一个步骤是固定的，老师授课都会按这个模板方法执行）
     */

    public void prepare(){
        System.out.println("准备好白板笔");
        System.out.println("打开投影仪");
    }

    public void end(){
        System.out.println("关闭投影仪");
        System.out.println("锁教室门");
    }

    /**
     * 定义一个抽象方法由其子类去具体实现
     */

    public abstract void teaching();

    /**
     * 上课教学的总流程方法
     */

    public void work(){
        //授课前准备
        prepare();
        //进行授课
        teaching();
        //授课结束
        end();
    }
}
