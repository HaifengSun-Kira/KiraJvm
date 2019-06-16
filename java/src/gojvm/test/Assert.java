package gojvm;

public class Assert {

    public static void assertTrue(boolean condition) {
        if (!condition) {
            fail();
        }
    }

    private static void fail() {
        throw new AssertionError("failed!");
    }

}
