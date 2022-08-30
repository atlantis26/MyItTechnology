package InterfaceAndClass;

/**
 * @Author: yangyong
 * @Date: 2020/4/8 15:16
 * @Version 1.0
 */

/**
 * 子类通过implements关键字实现接口，并对接口中定义的2个方法做具体的实现
 */
public class MemberServiceImplement implements MemberService{

    public void login(String userName, String passWord) {
        System.out.println("登录成功！用户名是："+userName+", 密码是："+passWord+", 常量是："+MemberService.NUM);
    }

    public void register(String userName, Integer age, String passWord){
        System.out.println("注册成功！用户名是："+userName+", 密码是："+passWord+"，年龄是："+age);
    }
}
