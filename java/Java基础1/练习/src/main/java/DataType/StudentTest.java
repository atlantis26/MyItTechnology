package DataType;

import InterfaceAndClass.Student;

/**
 * @Author: yangyong
 * @Date: 2020/4/9 18:31
 * @Version 1.0
 */
public class StudentTest {
    private String name;//学生姓名
//    private int score; //学生成绩, 如果使用基础类型int
    private Integer score;//使用包装类型Integer

//    public StudentTest(String name, int score) {
//        this.name = name;
//        this.score = score;
//    }

    public StudentTest(String name, Integer score) {
        this.name = name;
        this.score = score;
        System.out.println("姓名："+name+", 成绩："+score);
    }

    public static void main(String[] args){

        //小明考试很认真考了95分
        StudentTest xiaoming = new StudentTest("小明", 95);

        //小红考试睡着了考了0分
        StudentTest xiaohong = new StudentTest("小红", 0);

        /**
         * 小刚考试没来...？相当于缺考，和0分是两个概念，为了区分开，缺考需要填null。
         * 但是基础类型int不允许值为null, 用包装类型Integer才允许值为null，所以这种场景只能让score使用包装类型
         */
        StudentTest xiaogang = new StudentTest("小明", null);
    }
}
