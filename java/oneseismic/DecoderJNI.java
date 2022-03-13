package oneseismic;

public class DecoderJNI {

    static {
        System.loadLibrary("decoder");
    }

    public static void main(String[] args) {
        DecoderJNI decoder = new DecoderJNI();
        long ptr = decoder.create_decoder();
        decoder.status(ptr);
        
        // System.out.println("status:",status);
    }

    private native long create_decoder();
    
    private native void status(long ptr);

    private native void free_decoder(long ptr);
    
    private native void decode();
    
}
