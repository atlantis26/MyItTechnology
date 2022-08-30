package DataType;

import java.sql.Array;
import java.sql.SQLOutput;
import java.util.ArrayList;
import java.util.List;

/**
 * @Author: yangyong
 * @Date: 2020/4/9 18:47
 * @Version 1.0
 */
public class IntegerDemo {

    public static void main(String[] args) {

        /**
         * 阿里巴巴Java开发手册里的一句话：
         * 所有的包装类对象之间值的比较，全部使用equals方法，对于Integer var = ?
         * 在-128到127之间的赋值，Integer对象是在IntegerCache.cahche产生，会复用已有对象，这个区间内的Integer值可以直接使用==进行判断，
         * 但是这个区间之外的所有数据，都会在堆上产生，并不会复用已有对象，这是一个大坑，推荐使用equals方法进行判断
         */

        /**
         * 在Java 5中，引入了一个新功能来保存内存并提高Integer类型对象处理的性能。整数对象在内部缓存，并通过相同的引用对象重用。
         * 这适用于介于-127到+127（最大整数值）范围内的整数值。
         * 此整数缓存仅适用于自动装箱。使用构造函数构建整数对象时，它们不会被缓存。
         */

        /**
         * "=="的作用是判断两个对象的地址是不是相等。即判断两个对象是不是同一个对象。
         * 对于基本数据类型，==比较的是值。对于包装数据类型，==比较的是内存地址。
         */

        //值在-128到127之间，使用了自动装箱，所以会用到缓存，相同的引用对象重用都指向同一个缓存，所以相等是true
        Integer a1 = 127;
        Integer a2 = 127;
        System.out.println(a1 == a2);

        //值不在-128到127之间，使用了自动装箱, 但超出范围了，所以不会用到缓存，用==会是false
        Integer a3 = 129;
        Integer a4 = 129;
        System.out.println(a3 == a4);
//        System.out.println(a3.equals(a4));

        /**
         * 值在-128到127之间,但没有使用自动装箱
         * 但Integer a1 =xxx是自动装箱，自动装箱的本质是调用valueOf()，在valueOf()中使用了缓存
         * 而new Integer(xxx)是新创建对象,不是缓存，所以用==对比会是false
         */

        Integer a5 = new Integer(127);
        Integer a6 = new Integer(127);
        System.out.println(a5 == a6);
//        System.out.println(a5.equals(a6));

        //基础数据类型，==比较的是值，因此i7==i8会是true
        int i7 = 130;
        int i8 = 130;
        System.out.println(i7==i8);

    }
}