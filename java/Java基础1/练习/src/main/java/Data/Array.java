package Data;

/**
 * @Author: yangyong
 * @Date: 2020/4/10 14:48
 * @Version 1.0
 */
public class Array {
    public static void main(String args[]) {
        int data[] = new int[3]; /*开辟了一个长度为3的数组*/
        data[0] = 10; // 第一个元素
        data[1] = 20; // 第二个元素
        data[2] = 30; // 第三个元素
        for(int x = 0; x < data.length; x++) {
            System.out.println(data[x]); //通过循环控制索引
        }
    }
}

