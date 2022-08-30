package Data;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @Author: yangyong
 * @Date: 2020/4/10 13:11
 * @Version 1.0
 */

public class JsonAndMapAndList {
    public static void main(String []args) {
        String soap = "1111";
        List<String> list = new ArrayList<>();
        JSONObject json = new JSONObject();
        Map<String, String> map = new HashMap<>();

        //ArrayList以数组的方式存储数据，里面的元素是有顺序，可以重复的
        list.add(soap);
        list.add(soap);

        //json将数据以键值对的方式存储，键的key不可以相同，相同后面的值会将前面的值覆盖，json是无序的
        json.put("soap",soap);
        json.put("soap","2222");

        //HashMap将数据以键值对的方式存储，键的哈希码（hashCode）不可以相同，相同后面的值会将前面的值覆盖，HashMap是无序的,不是先进先出,要想有序，用LinkedHashMap
        map.put("soap",soap);
        map.put("soap","3333");
        map.put("aaa","5555");
        map.put("bbb","6666");
        map.put("ccc","7777");

        System.out.println("这是list的输出：" + list);
        System.out.println("这是json的输出：" + json);
        System.out.println("这是map的输出：" + map);

        JSONArray JsonArray = new JSONArray();
        JSONObject json1 = new JSONObject();
        JSONObject json2 = new JSONObject();

        json1.put("name","yy");
        json2.put("age", "18");
        JsonArray.add(json1);
        JsonArray.add(json2);
        System.out.println("这是JsonArray的输出：" + JsonArray);
    }
}
