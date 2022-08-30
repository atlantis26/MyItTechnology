package InterfaceAndClass;

/**
 * @Author: yangyong
 * @Date: 2020/4/8 15:14
 * @Version 1.0
 */
public interface MemberService {
    /**
     * java默认接口的方法是public abstract的，所以这里写public会提示public 是多余的。同时，如果你使用private或者protected都会报错
     * 定义声明2个方法，没有方法体，不会写具体的实现
     */

    /**
     * 接口中定义的变量是public static final类型，且必须有初值，实现类中不能重新定义或改变
     * 因为接口的含义：就是提供一种统一的’协议’,而接口中的属性也属于’协议’中的成员.它们是公共的,静态的,最终的常量.相当于全局常量.
     * 所以接口中的属性必然是常量，只能读不能改
     */
    public static final Integer NUM = 10;
    public void login(String userName, String passWord);
    void register(String userName, Integer age, String passWord);
}
