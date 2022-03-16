package oneseismic;

import java.nio.ByteBuffer;

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

    private long ptr;
    
    public DecoderJNI() {
        this.ptr = create_decoder();
    }

    public void process(byte[] data,int len) {
        buffer_and_process(ptr,data,len);
        status(ptr);
    }

    public String[] getAttr() {
        String[] result = header_attributes(this.ptr);
        return result;
    }

    public int[] getShape() {
        int[] result = header_shape(this.ptr);
        return result;
    }

    public void setWriter(String name,ByteBuffer buffer)
    {
        register_writer(this.ptr,name,buffer);
    }
    
    
    private native long create_decoder();

    private native void buffer_and_process(long ptr, byte[] data,int len);

    
    private native void header(long ptr);
    
    private native String[] header_attributes(long ptr);
    
    private native int[] header_shape(long ptr);

    private native void register_writer(long ptr,String name,ByteBuffer buf);
    
    private native void status(long ptr);

    private native void free_decoder(long ptr);
    
    private native void decode();
    
}
