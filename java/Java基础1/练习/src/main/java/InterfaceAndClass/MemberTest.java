package InterfaceAndClass;

/**
 * @Author: yangyong
 * @Date: 2020/4/8 15:21
 * @Version 1.0
 */
public class MemberTest {
    private MemberService memmmmm;

    public static void main(String[] args){

        //直接调用子类
        MemberServiceImplement member = new MemberServiceImplement();
        System.out.println("===========开始登录===============");
        member.login("yy", "123456");
        System.out.println("===========开始注册===============");
        member.register("yy",18,"123456");


        /**
         * 接口不能实例化，但可以实例化实现了该接口的类。MemberServiceImplement实现了接口MemberService
         * 则MemberServiceImplement实例化的对象可以给MemberService
         */
//        MemberService mem = new MemberServiceImplement();
//        System.out.println("===========开始登录===============");
//        mem.login("yy", "123456");
//        System.out.println("===========开始注册===============");
//        mem.register("yy",18,"123456");

    }
}
