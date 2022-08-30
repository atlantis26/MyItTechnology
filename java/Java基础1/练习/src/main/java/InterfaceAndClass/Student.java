package InterfaceAndClass;

/**
 * @Author: yangyong
 * @Date: 2020/4/9 10:17
 * @Version 1.0
 * 初学Java的构造方法时，根本不能理解构造方法有什么用，自己写程序的时候也从来没有用过,于是写的程序是这样的：
 */
public class Student {
    private String name;
    private long ID;
    private double score;

    public String getName() {
        return name;
    }
    public void setName(String name) {
        this.name = name;
    }
    public long getID() {
        return ID;
    }
    public void setID(long iD) {
        ID = iD;
    }
    public double getScore() {
        return score;
    }
    public void setScore(double score) {
        this.score = score;
    }

    /**
     * 但是这样写有个弊端，属性很多的时候，实例化对象后我需要set很多属性的值，我可能就会忘记初始化一些属性
     * 同时，这样还需要写许多get和set方法，虽然可以由IDE工具自动生成，但还是显得代码很冗长，不便于阅读。这个时候就显示出构造方法的重要性了
     */
    public static void main(String[] args) {
        Student s = new Student();
        s.setName("张三");
        s.setID(170316);
        s.setScore(86.0);
        System.out.println("学号为" + s.getID() + "的" + s.getName() + "的成绩为" + s.getScore() + "分。");
    }
}