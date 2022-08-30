package StaticMethod;

/**
 * @Author: yangyong
 * @Date: 2020/4/9 15:11
 * @Version 1.0
 */
public class StaticMethod {

    // 定义静态变量score1
    static Integer score1 = 86;
    // 定义静态变量score2
    static Integer score2 = 92;

    // 定义静态方法sum，计算成绩总分，并返回总分
    private static Integer sum() {
        return score1+score2;
    }

    public static void main(String[] args) {

        // 调用静态方法sum并接收返回值
        int allScore = StaticMethod.sum();
        System.out.println("总分：" + allScore);
    }
}
