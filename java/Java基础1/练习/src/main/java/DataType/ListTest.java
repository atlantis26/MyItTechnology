package DataType;

import java.util.ArrayList;
import java.util.List;

/**
 * @Author: yangyong
 * @Date: 2020/4/9 19:33
 * @Version 1.0
 */
public class ListTest {


    public static void main(String[] args) throws Exception {

        List list = new ArrayList();

        //List中添加基本数据类型，其实是会将其自动装箱成包装类，然后再添加进去
        list.add(111);//其实存储的是Integer
        list.add(222);//其实存储的是Integer
        list.add(333);//其实存储的是Integer

        /**
         * 删除的时候不会自动装箱,会把111当作索引，所以这样删除会报错
         * list.remove(111); 相当于删除索引位置111的数据
         */
        list.remove(111);

        //要删除111, 正确的写法是可以手动装箱
//        list.remove(Integer.valueOf(111));

        //或者111的index是0，list.remove(0)也可以
//        list.remove(0);

        System.out.println(list);


    }
}