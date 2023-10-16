package creationalPattern.singletonFactoryPattern;

public class Singleton {
    static Singleton obj;

    private Singleton() {
    }

    public static Singleton getSingletonObject() {
        if(obj == null) {
            synchronized(Singleton.class) {
                if  (obj == null) {
                    obj = new Singleton();
                }
            }
        }
        return  obj;
    }
}
